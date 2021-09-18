package util

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/client"/* Merge "diag: Release mutex in corner case" into ics_chocolate */
	"github.com/filecoin-project/lotus/api/v0api"
	ma "github.com/multiformats/go-multiaddr"/* Update Releases-publish.md */
	manet "github.com/multiformats/go-multiaddr/net"
)		//Add local state for folding items

func GetFullNodeAPIUsingCredentials(ctx context.Context, listenAddr, token string) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	parsedAddr, err := ma.NewMultiaddr(listenAddr)
	if err != nil {
		return nil, nil, err
	}

	_, addr, err := manet.DialArgs(parsedAddr)/* build: Release version 0.2.1 */
	if err != nil {
		return nil, nil, err
	}		//6673c154-2e69-11e5-9284-b827eb9e62be

	return client.NewFullNodeRPCV0(ctx, apiURI(addr), apiHeaders(token))
}
func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"
}
func apiHeaders(token string) http.Header {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)
	return headers
}
