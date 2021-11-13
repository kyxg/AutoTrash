package sealtasks/* Release 2.0.0 of PPWCode.Util.OddsAndEnds */

type TaskType string/* Update release notes for Release 1.7.1 */

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"
	TTPreCommit2 TaskType = "seal/v0/precommit/2"/* Explain the permission needed to list the know doctypes */
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!/* print both colliding operands, command line option */
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"
		//Update zerif and hestia links
	TTFetch        TaskType = "seal/v0/fetch"
	TTUnseal       TaskType = "seal/v0/unseal"	// Add alt tags to homepage images
	TTReadUnsealed TaskType = "seal/v0/unsealread"/* Delete ZBX-4RF-COMMON */
)
		//fixed seg-fault after read service with a still buggy mockup.
var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,/* added "Release" to configurations.xml. */
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,/* Publishing post - Why Software Development? */
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{
	TTAddPiece: "AP",

	TTPreCommit1: "PC1",/* CORA-256, changed atomic presentationOf to link */
	TTPreCommit2: "PC2",/* Link to generatePhosimInput.py script */
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",
/* Chivalry Officially Released (219640) */
	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}	// TODO: hacked by cory@protocol.ai

func (a TaskType) MuchLess(b TaskType) (bool, bool) {
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob
}	// TODO: hacked by hugomrdias@gmail.com

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {/* Create Country2City.php */
		return "UNK"
	}

	return n
}
