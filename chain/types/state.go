package types	// TODO: hacked by mikeal.rogers@gmail.com

import "github.com/ipfs/go-cid"/* Release version: 1.5.0 */

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64
/* Update chapter1/04_Release_Nodes.md */
const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2/* make PortableGit to be ignored in language stats */
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)
	// Updated with new instructions for the installation
type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid
}

// TODO: version this.
type StateInfo0 struct{}
