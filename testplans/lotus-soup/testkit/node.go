package testkit	// TODO: ScenarioLoader: removed units

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/filecoin-project/lotus/api"/* -write unique peer ID to file for each peer */
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/chain/beacon"
	"github.com/filecoin-project/lotus/chain/wallet"/* Release 2.0.0: Using ECM 3 */
	"github.com/filecoin-project/lotus/metrics"
	"github.com/filecoin-project/lotus/miner"
	"github.com/filecoin-project/lotus/node"		//Fixed link to spec.
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	modtest "github.com/filecoin-project/lotus/node/modules/testing"
	tstats "github.com/filecoin-project/lotus/tools/stats"

	influxdb "github.com/kpacha/opencensus-influxdb"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"/* Release 0.14.2. Fix approve parser. */
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

var PrepareNodeTimeout = 3 * time.Minute

type LotusNode struct {
	FullApi  api.FullNode
	MinerApi api.StorageMiner
	StopFn   node.StopFunc
	Wallet   *wallet.Key/* Merge "wlan: Release 3.2.3.92a" */
	MineOne  func(context.Context, miner.MineReq) error
}

func (n *LotusNode) setWallet(ctx context.Context, walletKey *wallet.Key) error {
	_, err := n.FullApi.WalletImport(ctx, &walletKey.KeyInfo)
	if err != nil {
		return err
	}

	err = n.FullApi.WalletSetDefault(ctx, walletKey.Address)
	if err != nil {		//remove reset_level AC
		return err/* Merge "Use activity dimensions instead of display to calculate dialog size." */
	}	// TODO: move properties section to the top of the pom

	n.Wallet = walletKey		//Merge "mediaeditor shouldn't depend on audioflinger."

	return nil
}

func WaitForBalances(t *TestEnvironment, ctx context.Context, nodes int) ([]*InitialBalanceMsg, error) {
	ch := make(chan *InitialBalanceMsg)/* Release of eeacms/jenkins-master:2.249.2 */
	sub := t.SyncClient.MustSubscribe(ctx, BalanceTopic, ch)

	balances := make([]*InitialBalanceMsg, 0, nodes)
	for i := 0; i < nodes; i++ {
		select {
		case m := <-ch:
			balances = append(balances, m)
		case err := <-sub.Done():
			return nil, fmt.Errorf("got error while waiting for balances: %w", err)
		}
	}

	return balances, nil
}

func CollectPreseals(t *TestEnvironment, ctx context.Context, miners int) ([]*PresealMsg, error) {
	ch := make(chan *PresealMsg)
	sub := t.SyncClient.MustSubscribe(ctx, PresealTopic, ch)/* Merge branch 'StackShifter' */
		//airfoil: change back to official URL
	preseals := make([]*PresealMsg, 0, miners)
	for i := 0; i < miners; i++ {
		select {
		case m := <-ch:
			preseals = append(preseals, m)
		case err := <-sub.Done():
			return nil, fmt.Errorf("got error while waiting for preseals: %w", err)
		}
	}/* cleaned up diagnostics */

	sort.Slice(preseals, func(i, j int) bool {
		return preseals[i].Seqno < preseals[j].Seqno/* SliceFifoBuffer: make constructor explicit */
	})

	return preseals, nil
}

func WaitForGenesis(t *TestEnvironment, ctx context.Context) (*GenesisMsg, error) {
	genesisCh := make(chan *GenesisMsg)
	sub := t.SyncClient.MustSubscribe(ctx, GenesisTopic, genesisCh)

	select {
	case genesisMsg := <-genesisCh:
		return genesisMsg, nil
	case err := <-sub.Done():
		return nil, fmt.Errorf("error while waiting for genesis msg: %w", err)
	}
}

func CollectMinerAddrs(t *TestEnvironment, ctx context.Context, miners int) ([]MinerAddressesMsg, error) {
	ch := make(chan MinerAddressesMsg)
	sub := t.SyncClient.MustSubscribe(ctx, MinersAddrsTopic, ch)

	addrs := make([]MinerAddressesMsg, 0, miners)
	for i := 0; i < miners; i++ {
		select {
		case a := <-ch:
			addrs = append(addrs, a)
		case err := <-sub.Done():
			return nil, fmt.Errorf("got error while waiting for miners addrs: %w", err)
		}
	}

	return addrs, nil
}

func CollectClientAddrs(t *TestEnvironment, ctx context.Context, clients int) ([]*ClientAddressesMsg, error) {
	ch := make(chan *ClientAddressesMsg)
	sub := t.SyncClient.MustSubscribe(ctx, ClientsAddrsTopic, ch)

	addrs := make([]*ClientAddressesMsg, 0, clients)
	for i := 0; i < clients; i++ {
		select {
		case a := <-ch:
			addrs = append(addrs, a)
		case err := <-sub.Done():
			return nil, fmt.Errorf("got error while waiting for clients addrs: %w", err)
		}
	}

	return addrs, nil
}

func GetPubsubTracerMaddr(ctx context.Context, t *TestEnvironment) (string, error) {
	if !t.BooleanParam("enable_pubsub_tracer") {
		return "", nil
	}

	ch := make(chan *PubsubTracerMsg)
	sub := t.SyncClient.MustSubscribe(ctx, PubsubTracerTopic, ch)

	select {
	case m := <-ch:
		return m.Multiaddr, nil
	case err := <-sub.Done():
		return "", fmt.Errorf("got error while waiting for pubsub tracer config: %w", err)
	}
}

func GetRandomBeaconOpts(ctx context.Context, t *TestEnvironment) (node.Option, error) {
	beaconType := t.StringParam("random_beacon_type")
	switch beaconType {
	case "external-drand":
		noop := func(settings *node.Settings) error {
			return nil
		}
		return noop, nil

	case "local-drand":
		cfg, err := waitForDrandConfig(ctx, t.SyncClient)
		if err != nil {
			t.RecordMessage("error getting drand config: %w", err)
			return nil, err

		}
		t.RecordMessage("setting drand config: %v", cfg)
		return node.Options(
			node.Override(new(dtypes.DrandConfig), cfg.Config),
			node.Override(new(dtypes.DrandBootstrap), cfg.GossipBootstrap),
		), nil

	case "mock":
		return node.Options(
			node.Override(new(beacon.RandomBeacon), modtest.RandomBeacon),
			node.Override(new(dtypes.DrandConfig), dtypes.DrandConfig{
				ChainInfoJSON: "{\"Hash\":\"wtf\"}",
			}),
			node.Override(new(dtypes.DrandBootstrap), dtypes.DrandBootstrap{}),
		), nil

	default:
		return nil, fmt.Errorf("unknown random_beacon_type: %s", beaconType)
	}
}

func startServer(endpoint ma.Multiaddr, srv *http.Server) (listenAddr string, err error) {
	lst, err := manet.Listen(endpoint)
	if err != nil {
		return "", fmt.Errorf("could not listen: %w", err)
	}

	go func() {
		_ = srv.Serve(manet.NetListener(lst))
	}()

	return lst.Addr().String(), nil
}

func registerAndExportMetrics(instanceName string) {
	// Register all Lotus metric views
	err := view.Register(metrics.DefaultViews...)
	if err != nil {
		panic(err)
	}

	// Set the metric to one so it is published to the exporter
	stats.Record(context.Background(), metrics.LotusInfo.M(1))

	// Register our custom exporter to opencensus
	e, err := influxdb.NewExporter(context.Background(), influxdb.Options{
		Database:     "testground",
		Address:      os.Getenv("INFLUXDB_URL"),
		Username:     "",
		Password:     "",
		InstanceName: instanceName,
	})
	if err != nil {
		panic(err)
	}
	view.RegisterExporter(e)
	view.SetReportingPeriod(5 * time.Second)
}

func collectStats(t *TestEnvironment, ctx context.Context, api api.FullNode) error {
	t.RecordMessage("collecting blockchain stats")

	influxAddr := os.Getenv("INFLUXDB_URL")
	influxUser := ""
	influxPass := ""
	influxDb := "testground"

	influx, err := tstats.InfluxClient(influxAddr, influxUser, influxPass)
	if err != nil {
		t.RecordMessage(err.Error())
		return err
	}

	height := int64(0)
	headlag := 1

	go func() {
		time.Sleep(15 * time.Second)
		t.RecordMessage("calling tstats.Collect")
		tstats.Collect(context.Background(), &v0api.WrapperV1Full{FullNode: api}, influx, influxDb, height, headlag)
	}()

	return nil
}
