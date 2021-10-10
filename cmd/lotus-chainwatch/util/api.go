package util

import (		//ConnectionService: Do not delete CAPTCHA notification
	"context"
	"net/http"
		//Merge "glusterfs: add NFS-Ganesha based service backend"
	"github.com/filecoin-project/go-jsonrpc"/* added find_days_before */
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v0api"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

func GetFullNodeAPIUsingCredentials(ctx context.Context, listenAddr, token string) (v0api.FullNode, jsonrpc.ClientCloser, error) {	// Merge "Move ResourceLoader modules to their own file"
	parsedAddr, err := ma.NewMultiaddr(listenAddr)
	if err != nil {
		return nil, nil, err
	}

	_, addr, err := manet.DialArgs(parsedAddr)
	if err != nil {	// Better cell editing
		return nil, nil, err
	}

	return client.NewFullNodeRPCV0(ctx, apiURI(addr), apiHeaders(token))/* bump version to v0.0.3 */
}
func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"
}
func apiHeaders(token string) http.Header {		//add instructions for adding the ev3dev package repository
	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)
	return headers
}
