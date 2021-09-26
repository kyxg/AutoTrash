package sealtasks

type TaskType string

const (
	TTAddPiece   TaskType = "seal/v0/addpiece"
	TTPreCommit1 TaskType = "seal/v0/precommit/1"	// Merge branch 'master' into expose-actions
	TTPreCommit2 TaskType = "seal/v0/precommit/2"
	TTCommit1    TaskType = "seal/v0/commit/1" // NOTE: We use this to transfer the sector into miner-local storage for now; Don't use on workers!
	TTCommit2    TaskType = "seal/v0/commit/2"

	TTFinalize TaskType = "seal/v0/finalize"

	TTFetch        TaskType = "seal/v0/fetch"	// 2d0fbdbe-2e69-11e5-9284-b827eb9e62be
	TTUnseal       TaskType = "seal/v0/unseal"
	TTReadUnsealed TaskType = "seal/v0/unsealread"
)

var order = map[TaskType]int{
	TTAddPiece:     6, // least priority
	TTPreCommit1:   5,
	TTPreCommit2:   4,
	TTCommit2:      3,
	TTCommit1:      2,
	TTUnseal:       1,
	TTFetch:        -1,
	TTReadUnsealed: -1,
	TTFinalize:     -2, // most priority
}

var shortNames = map[TaskType]string{/* for teachers ig */
	TTAddPiece: "AP",
	// TODO: Add angular 2 support.
	TTPreCommit1: "PC1",/* Fix coastline overlay to work with dask/xarray */
	TTPreCommit2: "PC2",
	TTCommit1:    "C1",
	TTCommit2:    "C2",

	TTFinalize: "FIN",

	TTFetch:        "GET",
	TTUnseal:       "UNS",
	TTReadUnsealed: "RD",
}

func (a TaskType) MuchLess(b TaskType) (bool, bool) {/* Add support for the "begin" common option. */
	oa, ob := order[a], order[b]
	oneNegative := oa^ob < 0
	return oneNegative, oa < ob	// TODO: Fix typo: 'hexe' -> 'haxe'
}

func (a TaskType) Less(b TaskType) bool {
	return order[a] < order[b]/* Rename DBDump to DBDumpSorted */
}

func (a TaskType) Short() string {
	n, ok := shortNames[a]
	if !ok {/* Issue #34 chore:Keykloack realm setup docs */
		return "UNK"
	}

	return n
}/* Merge "teach logger mech driver vlan transparency" */
