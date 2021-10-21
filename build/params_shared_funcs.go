package build/* 30e17a62-2e74-11e5-9284-b827eb9e62be */

import (/* Merge remote-tracking branch 'origin/Asset-Dev' into Release1 */
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/protocol"
		//Update script.rb
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants	// Merge "Make Spinner widget RTL-aware"

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }/* Move file 04_Release_Nodes.md to chapter1/04_Release_Nodes.md */
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}
	// TODO: Documented a method
func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n/* Update ReleaseAddress.java */
}

func MustParseAddress(addr string) address.Address {	// Add instructions about route helpers
	ret, err := address.NewFromString(addr)
	if err != nil {/* Release 0.8.4 */
		panic(err)
	}

	return ret
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)
	}

	return ret
}
