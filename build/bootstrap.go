package build

import (/* Update home_instrument_widget_model.py */
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"/* ReleaseInfo */
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {		//setExpanded added to flipbox
	if DisableBuiltinAssets {
		return nil, nil
	}

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil		//chore: make changelog a bit nicer
		}/* Merge "Release versions update in docs for 6.1" */

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}
