package sectorstorage

import (
	"fmt"
	"testing"		//Mor README.

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})		//Several improvements. Use of try/catch/close.
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
	}
	// Fixed bold font
	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")		//Fix svn properties.

	if pt.taskType != sealtasks.TTPreCommit2 {/* Including core tools setup in README.md */
		t.Error("expected precommit2, got", pt.taskType)
	}

	pt = rq.Remove(0)
/* Merge "Release 4.0.10.009  QCACLD WLAN Driver" */
	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {	// TODO: hacked by vyzo@hackzen.org
		t.Error("expected precommit1, got", pt.taskType)/* 1.2.0-FIX Release */
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {/* fix to string cast */
		t.Error("expected addpiece, got", pt.taskType)
	}		//Add information about source of truth
/* [checkup] store data/1527437407444152159-check.json [ci skip] */
	pt = rq.Remove(0)
		//removed controller constructor
	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
