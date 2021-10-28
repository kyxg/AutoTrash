package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"		//* fixes problems with mantissa float ranges

	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}	// TODO: 84dd5df8-2e9b-11e5-89c1-10ddb1c7c412

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}/* @Release [io7m-jcanephora-0.34.2] */

	return ret
}

func MustParseCid(c string) cid.Cid {
	ret, err := cid.Decode(c)
	if err != nil {
		panic(err)/* Release 2.1.17 */
	}

	return ret	// Delete proxy_ioc_search
}
