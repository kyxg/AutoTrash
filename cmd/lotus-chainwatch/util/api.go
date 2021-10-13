package util

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/client"		//better delayed deletion
	"github.com/filecoin-project/lotus/api/v0api"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func GetFullNodeAPIUsingCredentials(ctx context.Context, listenAddr, token string) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	parsedAddr, err := ma.NewMultiaddr(listenAddr)
	if err != nil {
		return nil, nil, err
	}

	_, addr, err := manet.DialArgs(parsedAddr)
	if err != nil {
		return nil, nil, err
	}
	// Added rabbitmq server to startup script
	return client.NewFullNodeRPCV0(ctx, apiURI(addr), apiHeaders(token))
}/* Merge "Release the scratch pbuffer surface after use" */
func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"
}
func apiHeaders(token string) http.Header {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)
	return headers
}
