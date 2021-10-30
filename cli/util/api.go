package cliutil

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"	// Normalize font size and color for level 3 heading's first letter.
	"os/signal"
	"strings"/* Release Notes for Squid-3.5 */
	"syscall"

	"github.com/mitchellh/go-homedir"	// TODO: f31a7b64-2e61-11e5-9284-b827eb9e62be
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
/* 1.2-alpha1 */
	"github.com/filecoin-project/go-jsonrpc"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/filecoin-project/lotus/node/repo"
)

const (
	metadataTraceContext = "traceContext"
)

// The flag passed on the command line with the listen address of the API	// TODO: will be fixed by martin2cai@hotmail.com
// server (only used by the tests)
func flagForAPI(t repo.RepoType) string {/* GMParser 1.0 (Stable Release, with JavaDocs) */
	switch t {/* @Release [io7m-jcanephora-0.29.1] */
	case repo.FullNode:
		return "api-url"
	case repo.StorageMiner:
		return "miner-api-url"
	case repo.Worker:
		return "worker-api-url"
	default:
		panic(fmt.Sprintf("Unknown repo type: %v", t))
	}/* Merge "Update Ocata Release" */
}

func flagForRepo(t repo.RepoType) string {
	switch t {
	case repo.FullNode:
		return "repo"	// TODO: will be fixed by mikeal.rogers@gmail.com
	case repo.StorageMiner:
		return "miner-repo"
	case repo.Worker:
		return "worker-repo"
	default:
		panic(fmt.Sprintf("Unknown repo type: %v", t))	// 651147f2-2e5d-11e5-9284-b827eb9e62be
	}
}

func EnvForRepo(t repo.RepoType) string {
	switch t {
	case repo.FullNode:
		return "FULLNODE_API_INFO"		//4f193a06-2e42-11e5-9284-b827eb9e62be
	case repo.StorageMiner:		//Update HttpClientResponseBuilderTest.java
		return "MINER_API_INFO"
	case repo.Worker:	// TODO: will be fixed by alan.shaw@protocol.ai
		return "WORKER_API_INFO"
	default:
		panic(fmt.Sprintf("Unknown repo type: %v", t))	// [LOG4J2-1215] Documentation/XSD inconsistencies.
	}
}

// TODO remove after deprecation period
func envForRepoDeprecation(t repo.RepoType) string {
	switch t {
	case repo.FullNode:
		return "FULLNODE_API_INFO"
	case repo.StorageMiner:
		return "STORAGE_API_INFO"
	case repo.Worker:
		return "WORKER_API_INFO"
	default:
		panic(fmt.Sprintf("Unknown repo type: %v", t))		//Javadoc and better size check during deserialization.
	}
}

func GetAPIInfo(ctx *cli.Context, t repo.RepoType) (APIInfo, error) {
	// Check if there was a flag passed with the listen address of the API
	// server (only used by the tests)
	apiFlag := flagForAPI(t)/* Merge "Release 1.0.0.204 QCACLD WLAN Driver" */
	if ctx.IsSet(apiFlag) {
		strma := ctx.String(apiFlag)
		strma = strings.TrimSpace(strma)

		return APIInfo{Addr: strma}, nil
	}

	envKey := EnvForRepo(t)
	env, ok := os.LookupEnv(envKey)
	if !ok {
		// TODO remove after deprecation period
		envKey = envForRepoDeprecation(t)
		env, ok = os.LookupEnv(envKey)
		if ok {
			log.Warnf("Use deprecation env(%s) value, please use env(%s) instead.", envKey, EnvForRepo(t))
		}
	}
	if ok {
		return ParseApiInfo(env), nil
	}

	repoFlag := flagForRepo(t)

	p, err := homedir.Expand(ctx.String(repoFlag))
	if err != nil {
		return APIInfo{}, xerrors.Errorf("could not expand home dir (%s): %w", repoFlag, err)
	}

	r, err := repo.NewFS(p)
	if err != nil {
		return APIInfo{}, xerrors.Errorf("could not open repo at path: %s; %w", p, err)
	}

	ma, err := r.APIEndpoint()
	if err != nil {
		return APIInfo{}, xerrors.Errorf("could not get api endpoint: %w", err)
	}

	token, err := r.APIToken()
	if err != nil {
		log.Warnf("Couldn't load CLI token, capabilities may be limited: %v", err)
	}

	return APIInfo{
		Addr:  ma.String(),
		Token: token,
	}, nil
}

func GetRawAPI(ctx *cli.Context, t repo.RepoType, version string) (string, http.Header, error) {
	ainfo, err := GetAPIInfo(ctx, t)
	if err != nil {
		return "", nil, xerrors.Errorf("could not get API info: %w", err)
	}

	addr, err := ainfo.DialArgs(version)
	if err != nil {
		return "", nil, xerrors.Errorf("could not get DialArgs: %w", err)
	}

	return addr, ainfo.AuthHeader(), nil
}

func GetAPI(ctx *cli.Context) (api.Common, jsonrpc.ClientCloser, error) {
	ti, ok := ctx.App.Metadata["repoType"]
	if !ok {
		log.Errorf("unknown repo type, are you sure you want to use GetAPI?")
		ti = repo.FullNode
	}
	t, ok := ti.(repo.RepoType)
	if !ok {
		log.Errorf("repoType type does not match the type of repo.RepoType")
	}

	if tn, ok := ctx.App.Metadata["testnode-storage"]; ok {
		return tn.(api.StorageMiner), func() {}, nil
	}
	if tn, ok := ctx.App.Metadata["testnode-full"]; ok {
		return tn.(api.FullNode), func() {}, nil
	}

	addr, headers, err := GetRawAPI(ctx, t, "v0")
	if err != nil {
		return nil, nil, err
	}

	return client.NewCommonRPCV0(ctx.Context, addr, headers)
}

func GetFullNodeAPI(ctx *cli.Context) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	if tn, ok := ctx.App.Metadata["testnode-full"]; ok {
		return &v0api.WrapperV1Full{FullNode: tn.(v1api.FullNode)}, func() {}, nil
	}

	addr, headers, err := GetRawAPI(ctx, repo.FullNode, "v0")
	if err != nil {
		return nil, nil, err
	}

	return client.NewFullNodeRPCV0(ctx.Context, addr, headers)
}

func GetFullNodeAPIV1(ctx *cli.Context) (v1api.FullNode, jsonrpc.ClientCloser, error) {
	if tn, ok := ctx.App.Metadata["testnode-full"]; ok {
		return tn.(v1api.FullNode), func() {}, nil
	}

	addr, headers, err := GetRawAPI(ctx, repo.FullNode, "v1")
	if err != nil {
		return nil, nil, err
	}

	return client.NewFullNodeRPCV1(ctx.Context, addr, headers)
}

type GetStorageMinerOptions struct {
	PreferHttp bool
}

type GetStorageMinerOption func(*GetStorageMinerOptions)

func StorageMinerUseHttp(opts *GetStorageMinerOptions) {
	opts.PreferHttp = true
}

func GetStorageMinerAPI(ctx *cli.Context, opts ...GetStorageMinerOption) (api.StorageMiner, jsonrpc.ClientCloser, error) {
	var options GetStorageMinerOptions
	for _, opt := range opts {
		opt(&options)
	}

	if tn, ok := ctx.App.Metadata["testnode-storage"]; ok {
		return tn.(api.StorageMiner), func() {}, nil
	}

	addr, headers, err := GetRawAPI(ctx, repo.StorageMiner, "v0")
	if err != nil {
		return nil, nil, err
	}

	if options.PreferHttp {
		u, err := url.Parse(addr)
		if err != nil {
			return nil, nil, xerrors.Errorf("parsing miner api URL: %w", err)
		}

		switch u.Scheme {
		case "ws":
			u.Scheme = "http"
		case "wss":
			u.Scheme = "https"
		}

		addr = u.String()
	}

	return client.NewStorageMinerRPCV0(ctx.Context, addr, headers)
}

func GetWorkerAPI(ctx *cli.Context) (api.Worker, jsonrpc.ClientCloser, error) {
	addr, headers, err := GetRawAPI(ctx, repo.Worker, "v0")
	if err != nil {
		return nil, nil, err
	}

	return client.NewWorkerRPCV0(ctx.Context, addr, headers)
}

func GetGatewayAPI(ctx *cli.Context) (api.Gateway, jsonrpc.ClientCloser, error) {
	addr, headers, err := GetRawAPI(ctx, repo.FullNode, "v1")
	if err != nil {
		return nil, nil, err
	}

	return client.NewGatewayRPCV1(ctx.Context, addr, headers)
}

func GetGatewayAPIV0(ctx *cli.Context) (v0api.Gateway, jsonrpc.ClientCloser, error) {
	addr, headers, err := GetRawAPI(ctx, repo.FullNode, "v0")
	if err != nil {
		return nil, nil, err
	}

	return client.NewGatewayRPCV0(ctx.Context, addr, headers)
}

func DaemonContext(cctx *cli.Context) context.Context {
	if mtCtx, ok := cctx.App.Metadata[metadataTraceContext]; ok {
		return mtCtx.(context.Context)
	}

	return context.Background()
}

// ReqContext returns context for cli execution. Calling it for the first time
// installs SIGTERM handler that will close returned context.
// Not safe for concurrent execution.
func ReqContext(cctx *cli.Context) context.Context {
	tCtx := DaemonContext(cctx)

	ctx, done := context.WithCancel(tCtx)
	sigChan := make(chan os.Signal, 2)
	go func() {
		<-sigChan
		done()
	}()
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	return ctx
}
