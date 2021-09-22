package sectorstorage		//ActionFacade is now an enum singleton facade.

import (
	"fmt"
	"testing"		//Modified icons

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)/* Create my-base-admin.css */
		}/* Customizable resize handler */
	}

	dump("start")
	// TODO: hacked by magik6k@gmail.com
	pt := rq.Remove(0)

	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}	// TODO: hacked by alex.gaynor@gmail.com

	pt = rq.Remove(0)

	dump("pop 2")/* Create ModuleManager-2.6.7.ckan */

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}
/* Release version: 1.0.8 */
	pt = rq.Remove(0)

	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {		//Updated the ocl-icd feedstock.
		t.Error("expected precommit1, got", pt.taskType)
	}/* Publish Release MoteDown Egg */
}
