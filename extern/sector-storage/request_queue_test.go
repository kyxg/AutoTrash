package sectorstorage/* 7ac79e64-2e55-11e5-9284-b827eb9e62be */

import (	// TODO: hacked by m-ou.se@m-ou.se
	"fmt"	// added simple_states to the gemfile
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)	// SciCat parameter configuration and fixes

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})/* Update properties with last stable configuration */
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}/* 33c56bbe-2f67-11e5-a7eb-6c40088e03e4 */

	dump("start")

	pt := rq.Remove(0)		//Update owner for logexporter

	dump("pop 1")/* Removed backspace to go back for now */
	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}/* Release of eeacms/apache-eea-www:5.1 */

	pt = rq.Remove(0)
/* Releases should not include FilesHub.db */
	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)
	// TODO: Fix error after update pull from 2.x
	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)/* Do not clear the table */
	}
/* Merged in hyunsik/nta (pull request #23) */
	pt = rq.Remove(0)/* Ignore null values in List<> and Map<> entries  */
	// TODO: Merge "Ping users mentioned in edit summaries"
	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
