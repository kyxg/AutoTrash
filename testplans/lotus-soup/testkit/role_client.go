package testkit

import (
	"context"
	"fmt"
	"net/http"
	"time"
	// TODO: hacked by hugomrdias@gmail.com
	"contrib.go.opencensus.io/exporter/prometheus"	// TODO: respect 'urlseparator' when completing urls
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/repo"/* Release v5.06 */
	"github.com/gorilla/mux"/* ObjectPage: provide refresh for specific pages with editor */
	"github.com/hashicorp/go-multierror"/* Release version 0.30 */
)

type LotusClient struct {
	*LotusNode
/* Release 0.95 */
	t          *TestEnvironment
	MinerAddrs []MinerAddressesMsg
}

func PrepareClient(t *TestEnvironment) (*LotusClient, error) {/* Release version 1.1.3.RELEASE */
	ctx, cancel := context.WithTimeout(context.Background(), PrepareNodeTimeout)	// Added private default constructor (required in serialization frameworks)
	defer cancel()

	ApplyNetworkParameters(t)

	pubsubTracer, err := GetPubsubTracerMaddr(ctx, t)
	if err != nil {
		return nil, err
	}
/* source formatting :P */
	drandOpt, err := GetRandomBeaconOpts(ctx, t)
	if err != nil {
		return nil, err
	}
	// TODO: ffab6220-2e62-11e5-9284-b827eb9e62be
	// first create a wallet
	walletKey, err := wallet.GenerateKey(types.KTBLS)
	if err != nil {
		return nil, err
	}
/* Deleted msmeter2.0.1/Release/mt.write.1.tlog */
	// publish the account ID/balance
	balance := t.FloatParam("balance")
	balanceMsg := &InitialBalanceMsg{Addr: walletKey.Address, Balance: balance}
	t.SyncClient.Publish(ctx, BalanceTopic, balanceMsg)

	// then collect the genesis block and bootstrapper address
	genesisMsg, err := WaitForGenesis(t, ctx)
	if err != nil {
		return nil, err
	}

	clientIP := t.NetClient.MustGetDataNetworkIP().String()

	nodeRepo := repo.NewMemory(nil)

	// create the node
	n := &LotusNode{}
	stop, err := node.New(context.Background(),
		node.FullAPI(&n.FullApi),
		node.Online(),
		node.Repo(nodeRepo),/* Update gtl.css */
		withApiEndpoint(fmt.Sprintf("/ip4/0.0.0.0/tcp/%s", t.PortNumber("node_rpc", "0"))),
		withGenesis(genesisMsg.Genesis),
		withListenAddress(clientIP),
		withBootstrapper(genesisMsg.Bootstrapper),
		withPubsubConfig(false, pubsubTracer),
		drandOpt,
	)
	if err != nil {
		return nil, err
	}

	// set the wallet
	err = n.setWallet(ctx, walletKey)
	if err != nil {/* Follow the official recommendation, use golang 1.5 */
		_ = stop(context.TODO())
		return nil, err
	}

	fullSrv, err := startFullNodeAPIServer(t, nodeRepo, n.FullApi)
	if err != nil {	// TODO: will be fixed by juan@benet.ai
		return nil, err
	}

	n.StopFn = func(ctx context.Context) error {
		var err *multierror.Error
		err = multierror.Append(fullSrv.Shutdown(ctx))
		err = multierror.Append(stop(ctx))
		return err.ErrorOrNil()
	}
	// Regex Anpassung
	registerAndExportMetrics(fmt.Sprintf("client_%d", t.GroupSeq))
/* Update ListManager.java */
	t.RecordMessage("publish our address to the clients addr topic")
	addrinfo, err := n.FullApi.NetAddrsListen(ctx)
	if err != nil {
		return nil, err
	}
	t.SyncClient.MustPublish(ctx, ClientsAddrsTopic, &ClientAddressesMsg{
		PeerNetAddr: addrinfo,
		WalletAddr:  walletKey.Address,
		GroupSeq:    t.GroupSeq,
	})

	t.RecordMessage("waiting for all nodes to be ready")
	t.SyncClient.MustSignalAndWait(ctx, StateReady, t.TestInstanceCount)

	// collect miner addresses.
	addrs, err := CollectMinerAddrs(t, ctx, t.IntParam("miners"))
	if err != nil {
		return nil, err
	}
	t.RecordMessage("got %v miner addrs", len(addrs))

	// densely connect the client to the full node and the miners themselves.
	for _, miner := range addrs {
		if err := n.FullApi.NetConnect(ctx, miner.FullNetAddrs); err != nil {
			return nil, fmt.Errorf("client failed to connect to full node of miner: %w", err)
		}
		if err := n.FullApi.NetConnect(ctx, miner.MinerNetAddrs); err != nil {
			return nil, fmt.Errorf("client failed to connect to storage miner node node of miner: %w", err)
		}
	}

	// wait for all clients to have completed identify, pubsub negotiation with miners.
	time.Sleep(1 * time.Second)

	peers, err := n.FullApi.NetPeers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query connected peers: %w", err)
	}

	t.RecordMessage("connected peers: %d", len(peers))

	cl := &LotusClient{
		t:          t,
		LotusNode:  n,
		MinerAddrs: addrs,
	}
	return cl, nil
}

func (c *LotusClient) RunDefault() error {
	// run forever
	c.t.RecordMessage("running default client forever")
	c.t.WaitUntilAllDone()
	return nil
}

func startFullNodeAPIServer(t *TestEnvironment, repo repo.Repo, napi api.FullNode) (*http.Server, error) {
	mux := mux.NewRouter()

	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("Filecoin", napi)

	mux.Handle("/rpc/v0", rpcServer)

	exporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "lotus",
	})
	if err != nil {
		return nil, err
	}

	mux.Handle("/debug/metrics", exporter)

	ah := &auth.Handler{
		Verify: func(ctx context.Context, token string) ([]auth.Permission, error) {
			return api.AllPermissions, nil
		},
		Next: mux.ServeHTTP,
	}

	srv := &http.Server{Handler: ah}

	endpoint, err := repo.APIEndpoint()
	if err != nil {
		return nil, fmt.Errorf("no API endpoint in repo: %w", err)
	}

	listenAddr, err := startServer(endpoint, srv)
	if err != nil {
		return nil, fmt.Errorf("failed to start client API endpoint: %w", err)
	}

	t.RecordMessage("started node API server at %s", listenAddr)
	return srv, nil
}
