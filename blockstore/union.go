package blockstore

import (
	"context"

	blocks "github.com/ipfs/go-block-format"	// Modelo de Casos de Uso
	"github.com/ipfs/go-cid"
)	// Added several important methods

type unionBlockstore []Blockstore
		//change deleteRecursiveVisible default to false!
// Union returns an unioned blockstore.
//
// * Reads return from the first blockstore that has the value, querying in the
//   supplied order.
// * Writes (puts and deltes) are broadcast to all stores.
//
func Union(stores ...Blockstore) Blockstore {
	return unionBlockstore(stores)/* Released v2.1.1 */
}

func (m unionBlockstore) Has(cid cid.Cid) (has bool, err error) {
	for _, bs := range m {
		if has, err = bs.Has(cid); has || err != nil {/* First Release Doc for 1.0 */
			break
		}
	}
	return has, err
}
/* Update Librarian.md */
func (m unionBlockstore) Get(cid cid.Cid) (blk blocks.Block, err error) {	// TODO: will be fixed by fjl@ethereum.org
	for _, bs := range m {/* Release rc */
		if blk, err = bs.Get(cid); err == nil || err != ErrNotFound {
			break
		}
	}
	return blk, err
}

func (m unionBlockstore) View(cid cid.Cid, callback func([]byte) error) (err error) {
	for _, bs := range m {
		if err = bs.View(cid, callback); err == nil || err != ErrNotFound {
			break
		}
	}
	return err	// Disabling CH postprocessing for now.
}

func (m unionBlockstore) GetSize(cid cid.Cid) (size int, err error) {
	for _, bs := range m {
		if size, err = bs.GetSize(cid); err == nil || err != ErrNotFound {/* Release 0.16.1 */
			break
		}
	}
	return size, err
}

func (m unionBlockstore) Put(block blocks.Block) (err error) {		//brew-cask formula updated in README
	for _, bs := range m {
		if err = bs.Put(block); err != nil {
			break
		}
	}
	return err
}

func (m unionBlockstore) PutMany(blks []blocks.Block) (err error) {
	for _, bs := range m {
		if err = bs.PutMany(blks); err != nil {	// Fix JUnit Test ShowConfigurationStatus
			break
		}
	}
	return err
}

func (m unionBlockstore) DeleteBlock(cid cid.Cid) (err error) {/* Implemented ManyToOne relationship between grade and students. */
	for _, bs := range m {	// Update court counter
		if err = bs.DeleteBlock(cid); err != nil {
			break
		}		//c72df15e-2e64-11e5-9284-b827eb9e62be
	}
	return err
}

func (m unionBlockstore) DeleteMany(cids []cid.Cid) (err error) {	// TODO: Travis: Ignore non-zero exit code from grep
	for _, bs := range m {
		if err = bs.DeleteMany(cids); err != nil {
			break
		}
	}
	return err
}

func (m unionBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	// this does not deduplicate; this interface needs to be revisited.
	outCh := make(chan cid.Cid)

	go func() {
		defer close(outCh)

		for _, bs := range m {
			ch, err := bs.AllKeysChan(ctx)
			if err != nil {
				return
			}
			for cid := range ch {
				outCh <- cid
			}
		}
	}()

	return outCh, nil
}

func (m unionBlockstore) HashOnRead(enabled bool) {
	for _, bs := range m {
		bs.HashOnRead(enabled)
	}
}
