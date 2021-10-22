package sealtasks/* Deleted unused multi-transformation */
/* Release dhcpcd-6.4.0 */
type TaskType string
/* Delete e4u.sh - 2nd Release */
const (
	TTAddPiece   TaskType = "seal/v0/addpiece"/* delete notes */
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"		//Merge "Fix bug 6029592 - font size setting causes clipped icon menu items"

	TTFetch        TaskType = "seal/v0/fetch"	// TODO: will be fixed by alex.gaynor@gmail.com
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)
/* Release 0.3.0 */
var order = map[TaskType]int{/* Merge "[Release Notes] Update for HA and API guides for Mitaka" */
	TTAddPiece:     6, // least priority/* fix email sync */
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,/* 🚀 Visual Storyteller */
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority/* GitVersion: guess we are back at WeightedPreReleaseNumber */
}
/* Exempt ALL THE THINGS! */
var shortNames = map[TaskType]string{
	TTAddPiece: "AP",
	// Merge branch 'master' into 520-check-icmp-multiple-hosts
	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",
/* merge with 1.1.1 */
	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {/* Create boto_tools.py */
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}
