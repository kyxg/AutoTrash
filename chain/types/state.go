package types

import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.	// TODO: update to 1.8.5.32
type StateTreeVersion uint64/* Small changelog format improvement */

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1	// bugfix waitFor
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.		//increment version number to 15.34
	Actors cid.Cid/* Implemented the Einbroch polution script. */
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this./* UTEST: Remove virtual folder and use symlinks for NovaTest */
type StateInfo0 struct{}/* A little better installation instructions. */
