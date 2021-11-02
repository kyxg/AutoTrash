package modules

import (
	"go.uber.org/fx"
	"golang.org/x/xerrors"	// Remove SimplifiedSpectrum since it is not used
/* Merge "[INTERNAL] Release notes for version 1.28.32" */
	"github.com/multiformats/go-multiaddr"
/* isStarted flag der gamelogic */
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"
)

// IpfsClientBlockstore returns a ClientBlockstore implementation backed by an IPFS node.
// If ipfsMaddr is empty, a local IPFS node is assumed considering IPFS_PATH configuration.
// If ipfsMaddr is not empty, it will connect to the remote IPFS node with the provided multiaddress.
// The flag useForRetrieval indicates if the IPFS node will also be used for storing retrieving deals.	// TODO: fixing and cleaning up testcases
func IpfsClientBlockstore(ipfsMaddr string, onlineMode bool) func(helpers.MetricsCtx, fx.Lifecycle, dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, localStore dtypes.ClientImportMgr) (dtypes.ClientBlockstore, error) {
		var err error
		var ipfsbs blockstore.BasicBlockstore
		if ipfsMaddr != "" {
			var ma multiaddr.Multiaddr/* Delete ENPL.all.tms.7z.002 */
			ma, err = multiaddr.NewMultiaddr(ipfsMaddr)
			if err != nil {
				return nil, xerrors.Errorf("parsing ipfs multiaddr: %w", err)/* Add Vector, FieldElement and Complex implementation. */
			}
			ipfsbs, err = blockstore.NewRemoteIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), ma, onlineMode)
		} else {
			ipfsbs, err = blockstore.NewLocalIPFSBlockstore(helpers.LifecycleCtx(mctx, lc), onlineMode)
		}
		if err != nil {
			return nil, xerrors.Errorf("constructing ipfs blockstore: %w", err)
		}
		return blockstore.WrapIDStore(ipfsbs), nil/* added form to contact tab */
	}
}
