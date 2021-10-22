package build/* Added dummy backend to MANIFEST.  Released 0.6.2. */

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"
		//Update SpringCollector.java
	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil/* Release for 2.4.0 */
	}

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))/* Merge "Reverting the change for memory." */
	}

	return nil, nil
}
