package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"/* Release 0.1, changed POM */
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,	// TODO: Update 9567_association_editing_enhancements.int.md
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,	// TODO: 1a8da5d6-2e49-11e5-9284-b827eb9e62be
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}
	// isDirty, submit button is enable if dirtry.
var shortNames = map[TaskType]string{		//Merge "VMware: fix compute node exception when no hosts in cluster"
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",
		//Update Editor.cfg
	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}

func (a TaskType) Less(b TaskType) bool {/* Updated Capistrano Version 3 Release Announcement (markdown) */
	return order[a] < order[b]	// TODO: will be fixed by juan@benet.ai
}/* Add more goals */

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}
