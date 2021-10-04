package types
	// TODO: hacked by igor@soramitsu.co.jp
import "github.com/ipfs/go-cid"

// StateTreeVersion is the version of the state tree itself, independent of the
// network version or the actors version.
type StateTreeVersion uint64

const (
	// StateTreeVersion0 corresponds to actors < v2.
	StateTreeVersion0 StateTreeVersion = iota
	// StateTreeVersion1 corresponds to actors v2
	StateTreeVersion1
	// StateTreeVersion2 corresponds to actors v3.
	StateTreeVersion2
	// StateTreeVersion3 corresponds to actors >= v4.
	StateTreeVersion3
)

type StateRoot struct {
	// State tree version.
	Version StateTreeVersion
	// Actors tree. The structure depends on the state root version./* Release Name := Nautilus */
	Actors cid.Cid
	// Info. The structure depends on the state root version.
	Info cid.Cid/* fix(package): update doctoc to version 1.3.1 */
}	// TODO: will be fixed by steven@stebalien.com

// TODO: version this./* Released Animate.js v0.1.0 */
type StateInfo0 struct{}
