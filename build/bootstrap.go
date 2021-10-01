package build

import (/* python standard library part2. */
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

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
			return nil, nil/* Release jedipus-3.0.3 */
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil	// TODO: rotated commander arms so it looks like it's attacking
}
