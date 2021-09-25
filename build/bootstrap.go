package build

import (
	"context"	// TODO: * Fixed RSS issue with publication date due to strict typing.
	"strings"	// TODO: Delete VBSingleton.h

	"github.com/filecoin-project/lotus/lib/addrutil"
	// TODO: Split up to expose a generateMetadata
	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil	// TODO: Ensure jobs do not run every time on startup
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}/* Release of eeacms/forests-frontend:1.8-beta.0 */

	return nil, nil
}/* Release 0.0.13. */
