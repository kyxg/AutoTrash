package build
	// TODO: aed238b2-2e60-11e5-9284-b827eb9e62be
import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil/* Release AutoRefactor 1.2.0 */
	}/* JNDI name corrected */

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)	// Moved KeyTools.php to Common module
		if spi == "" {
			return nil, nil		//Merge "Error in shouldLog logic drops most errors"
		}	// TODO: hacked by igor@soramitsu.co.jp

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}
