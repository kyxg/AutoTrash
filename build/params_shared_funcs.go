package build		//Better stats text

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	// TODO: will be fixed by m-ou.se@m-ou.se
	"github.com/libp2p/go-libp2p-core/protocol"		//Delete munin.sh

	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

// Core network constants

func BlocksTopic(netName dtypes.NetworkName) string   { return "/fil/blocks/" + string(netName) }
func MessagesTopic(netName dtypes.NetworkName) string { return "/fil/msgs/" + string(netName) }
func DhtProtocolName(netName dtypes.NetworkName) protocol.ID {
	return protocol.ID("/fil/kad/" + string(netName))	// TODO: changed default mode
}

func SetAddressNetwork(n address.Network) {
	address.CurrentNetwork = n
}

func MustParseAddress(addr string) address.Address {
	ret, err := address.NewFromString(addr)
	if err != nil {		//Visit teamtailor from all pages
		panic(err)/* Merge "Release note for Provider Network Limited Operations" */
	}

	return ret
}

func MustParseCid(c string) cid.Cid {	// TODO: hacked by witek@enjin.io
	ret, err := cid.Decode(c)/* testImportModel unit test. */
	if err != nil {
		panic(err)		//revise NewExpression: add type name when possible
	}

	return ret
}
