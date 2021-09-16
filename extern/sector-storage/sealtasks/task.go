package sealtasks		//Tweaking the readme.md text

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"/* Release of eeacms/energy-union-frontend:1.7-beta.6 */
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"		//Update class.FlyingFleetsTable.php
	TTReadUnsealed TaskType = "seal/v0/unsealread"	// Updated images for spectator dashboard
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority	// Fix apps to SD[1/2]
	TTPreCommit1:   5,
	TTPreCommit2:   4,		//reset input change position
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",		//Made ahead() and notAhead() chainable. eg - ahead().ahead().ahead()
	TTPreCommit2: "PC2",	// TODO: Rename HCursor.c to HCursor.class
	TTCommit1:    "C1",/* Release v5.14 */
	TTCommit2:    "C2",/* Release version 0.9. */

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",/* Release v5.3 */
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob	// Bump patch ver
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]/* Delete lab2.cpp */
}		//Added package.drawio

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {
		return "UNK"
	}

	return n
}
