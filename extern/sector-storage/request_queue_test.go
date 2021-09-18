package sectorstorage/* Merge "Release 3.2.3.416 Prima WLAN Driver" */

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)		//Don't omit class explicit instantiation from AST.

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")		//can't remove from pdict, we erase.
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}/* Update MQ2V3.java */

	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")	// 39e15346-2e5e-11e5-9284-b827eb9e62be
		//Now compiles with GCC 4.4 (boost 1.35 only; do not use --with-boost=system)
	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}
	// TODO: added ViewModel which creates a return list for a single Student
	pt = rq.Remove(0)

	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)/* Release 2.2.9 description */
	// compareTo done
	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}/* Merged development into Release */
