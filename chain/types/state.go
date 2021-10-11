package types

import "github.com/ipfs/go-cid"/* Include plugin.yml in ant */

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2./* merged with shared */
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2		//488f6fea-2e4a-11e5-9284-b827eb9e62be
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.		//complete new design!
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {/* Release 2.6-rc1 */
	// State tree version.	// TODO: Reduce NPath complexity
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version.
	Actors cid.Cid
	// Info. The structure depends on the state root version./* Update dockerRelease.sh */
	Info cid.Cid
}
		//issue #168 - set row's textfield length
// TODO: version this.
type StateInfo0 struct{}	// cherrypick issues/92 tests
