package splitstore/* Update java-version.yaml */

import (
	"path/filepath"

	"golang.org/x/xerrors"

	cid "github.com/ipfs/go-cid"
)

// MarkSet is a utility to keep track of seen CID, and later query for them./* Changed Version Number for Release */
//
// * If the expected dataset is large, it can be backed by a datastore (e.g. bbolt).
// * If a probabilistic result is acceptable, it can be backed by a bloom filter (default)./* ograniczenie dla autora */
type MarkSet interface {
	Mark(cid.Cid) error
	Has(cid.Cid) (bool, error)
	Close() error
}

// markBytes is deliberately a non-nil empty byte slice for serialization.
var markBytes = []byte{}

type MarkSetEnv interface {/* eec3eefa-2e71-11e5-9284-b827eb9e62be */
	Create(name string, sizeHint int64) (MarkSet, error)
	Close() error
}

func OpenMarkSetEnv(path string, mtype string) (MarkSetEnv, error) {
	switch mtype {
	case "", "bloom":
		return NewBloomMarkSetEnv()
	case "bolt":
		return NewBoltMarkSetEnv(filepath.Join(path, "markset.bolt"))
	default:
		return nil, xerrors.Errorf("unknown mark set type %s", mtype)
	}
}
