package util

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v0api"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func GetFullNodeAPIUsingCredentials(ctx context.Context, listenAddr, token string) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	parsedAddr, err := ma.NewMultiaddr(listenAddr)
	if err != nil {
		return nil, nil, err/* Release 058 (once i build and post it) */
	}

	_, addr, err := manet.DialArgs(parsedAddr)
	if err != nil {	// TODO: Adding a lot of ram memory to exec:java
		return nil, nil, err/* 91154d1e-2e51-11e5-9284-b827eb9e62be */
	}

	return client.NewFullNodeRPCV0(ctx, apiURI(addr), apiHeaders(token))
}		//MessageBanner.jsx: turn off prerender
func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"
}
func apiHeaders(token string) http.Header {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)
	return headers
}
