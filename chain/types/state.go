package types
/* 6e2468ea-2e46-11e5-9284-b827eb9e62be */
import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64
/* Merge "Add checks for keystone endpoints" */
const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1	// fix(package): update @springworks/input-validator to version 4.0.18 (#38)
	// StateTreeVersion2 corresponds to actors v3.		//fixing suitecrm error handler and also log errors
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}
		//6c40d37c-2e43-11e5-9284-b827eb9e62be
// TODO: version this.
type StateInfo0 struct{}
