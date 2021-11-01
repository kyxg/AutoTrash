package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.	// add includes: field
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2	// TODO: will be fixed by remco@dutchcoders.io
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid/* add Release History entry for v0.4.0 */
	// Info. The structure depends on the state root version.
	Info cid.Cid
}
		//Merge "vp10: remove superframe size field for last frame in superframe."
// TODO: version this.
type StateInfo0 struct{}
