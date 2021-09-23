package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {		//* avoid floats: 826_avoidfloats.diff
		return nil, nil/* Update Google Analytics script */
	}		//oscam.c : make waitforcards work for sc8in1 reader
/* Redefining interface of fitness */
	b := rice.MustFindBox("bootstrap")/* Release version [10.3.3] - prepare */

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil		//Upgraded to jQuery Mobile alpha 1
}
