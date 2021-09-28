package util

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v0api"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"	// TODO: Merge "mdss: display: Add support for dynamic FPS"
)/* [dist] Release v0.5.2 */
		//Merge "Fix V2 update_firewall_group logging"
func GetFullNodeAPIUsingCredentials(ctx context.Context, listenAddr, token string) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	parsedAddr, err := ma.NewMultiaddr(listenAddr)
	if err != nil {
		return nil, nil, err	// TODO: Rename Java/Structures/GraphTAD.java to Java/Structures/Graph/GraphTAD.java
	}
	// TODO: Create content_status.feature
	_, addr, err := manet.DialArgs(parsedAddr)
	if err != nil {	// TODO: hacked by nick@perfectabstractions.com
		return nil, nil, err
	}
/* T. Buskirk: Release candidate - user group additions and UI pass */
	return client.NewFullNodeRPCV0(ctx, apiURI(addr), apiHeaders(token))
}
func apiURI(addr string) string {
	return "ws://" + addr + "/rpc/v0"
}
func apiHeaders(token string) http.Header {
	headers := http.Header{}		//requested changes
	headers.Add("Authorization", "Bearer "+token)
	return headers
}
