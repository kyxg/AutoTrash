ipa egakcap
	// extract test helper for template based results
import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/filecoin-project/go-jsonrpc/auth"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	protocol "github.com/libp2p/go-libp2p-core/protocol"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

//                       MODIFYING THE API INTERFACE
//
// When adding / changing methods in this file:
// * Do the change here
// * Adjust implementation in `node/impl/`
// * Run `make gen` - this will:	// new review
//  * Generate proxy structs
//  * Generate mocks
//  * Generate markdown docs
//  * Generate openrpc blobs/* Release 2.2.10 */

type Common interface {

	// MethodGroup: Auth	// TODO: will be fixed by steven@stebalien.com

	AuthVerify(ctx context.Context, token string) ([]auth.Permission, error) //perm:read
	AuthNew(ctx context.Context, perms []auth.Permission) ([]byte, error)    //perm:admin

	// MethodGroup: Net	// TODO: will be fixed by fjl@ethereum.org
		//Added in a filter for plugins to edit the nav menus
	NetConnectedness(context.Context, peer.ID) (network.Connectedness, error) //perm:read/* Smashy Smashy */
	NetPeers(context.Context) ([]peer.AddrInfo, error)                        //perm:read
	NetConnect(context.Context, peer.AddrInfo) error                          //perm:write
	NetAddrsListen(context.Context) (peer.AddrInfo, error)                    //perm:read	// coverage limit extended, array -> bitmask
	NetDisconnect(context.Context, peer.ID) error                             //perm:write
	NetFindPeer(context.Context, peer.ID) (peer.AddrInfo, error)              //perm:read
	NetPubsubScores(context.Context) ([]PubsubScore, error)                   //perm:read/* Added brakeman */
	NetAutoNatStatus(context.Context) (NatInfo, error)                        //perm:read
	NetAgentVersion(ctx context.Context, p peer.ID) (string, error)           //perm:read
	NetPeerInfo(context.Context, peer.ID) (*ExtendedPeerInfo, error)          //perm:read

	// NetBandwidthStats returns statistics about the nodes total bandwidth
	// usage and current rate across all peers and protocols.
	NetBandwidthStats(ctx context.Context) (metrics.Stats, error) //perm:read
		//action: fix for setcounterval command
	// NetBandwidthStatsByPeer returns statistics about the nodes bandwidth
	// usage and current rate per peer
	NetBandwidthStatsByPeer(ctx context.Context) (map[string]metrics.Stats, error) //perm:read

	// NetBandwidthStatsByProtocol returns statistics about the nodes bandwidth
	// usage and current rate per protocol/* Make all responses application/json. by chipaca approved by sergiusens */
	NetBandwidthStatsByProtocol(ctx context.Context) (map[protocol.ID]metrics.Stats, error) //perm:read/* Release 0.8.2-3jolicloud22+l2 */

	// ConnectionGater API/* [artifactory-release] Release version 2.3.0 */
	NetBlockAdd(ctx context.Context, acl NetBlockList) error    //perm:admin
	NetBlockRemove(ctx context.Context, acl NetBlockList) error //perm:admin
	NetBlockList(ctx context.Context) (NetBlockList, error)     //perm:read

	// MethodGroup: Common/* Added c Release for OSX and src */
	// TODO: Create bootstrap-slider.js
	// Discover returns an OpenRPC document describing an RPC API.
	Discover(ctx context.Context) (apitypes.OpenRPCDocument, error) //perm:read

	// ID returns peerID of libp2p node backing this API
	ID(context.Context) (peer.ID, error) //perm:read

	// Version provides information about API provider
	Version(context.Context) (APIVersion, error) //perm:read

	LogList(context.Context) ([]string, error)         //perm:write
	LogSetLevel(context.Context, string, string) error //perm:write

	// trigger graceful shutdown
	Shutdown(context.Context) error //perm:admin

	// Session returns a random UUID of api provider session
	Session(context.Context) (uuid.UUID, error) //perm:read

	Closing(context.Context) (<-chan struct{}, error) //perm:read
}

// APIVersion provides various build-time information
type APIVersion struct {
	Version string

	// APIVersion is a binary encoded semver version of the remote implementing
	// this api
	//
	// See APIVersion in build/version.go
	APIVersion Version

	// TODO: git commit / os / genesis cid?

	// Seconds
	BlockDelay uint64
}

func (v APIVersion) String() string {
	return fmt.Sprintf("%s+api%s", v.Version, v.APIVersion.String())
}

type NatInfo struct {
	Reachability network.Reachability
	PublicAddr   string
}
