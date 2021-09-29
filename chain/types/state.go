package types/* Release 1.3.1 v4 */

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64/* Updating README with OneButton integration. */

const (		//37e062a8-2e68-11e5-9284-b827eb9e62be
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota/* reposition */
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.	// Small grammar tweaks. Add Anthony to authors.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid/* Ready for Release on Zenodo. */
}

// TODO: version this.
type StateInfo0 struct{}
