package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"	// TODO: hacked by greg@colvin.org

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)	// TODO: hacked by souzau@yandex.com

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }/* Release version 0.17. */
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}

	return ret
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)	// PLAT-1985 add tooltip with url
	if err != nil {/* Release v 2.0.2 */
		panic(err)/* Added more support for events. */
	}

	return ret
}
