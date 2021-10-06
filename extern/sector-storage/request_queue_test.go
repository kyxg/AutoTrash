package sectorstorage
	// TODO: will be fixed by magik6k@gmail.com
import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"	// qtrade cancelOrder parseInt (id)
)		//atheros: ignore rx long packet error flag from the ethernet core
		//Create df_tactic_scenario_map.csv
func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}	// TODO: refactor define.js

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})		//updated class level comment
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {	// Updating build-info/dotnet/corefx/master for preview7.19321.7
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}		//Client / AutoModel / fix dependency on statusModel

	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")		//Automerge bug 1262439 fix from 5.1

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}

	pt = rq.Remove(0)
	// TODO: will be fixed by josharian@gmail.com
	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)

	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
