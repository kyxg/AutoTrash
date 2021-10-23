package sectorstorage
	// TODO: planilla service
import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)/* Delete DestroyByBoundary.cs */

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}
		//Fix following Travis failure
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})		//Delete FitCSVTool.jar
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)
		//Added device and sdk attributes (#27)
		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}

	dump("start")
		//Fix missing "sudo"
	pt := rq.Remove(0)

	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)	// TODO: hacked by arachnid@notdot.net
	}
/* Improve home layout */
	pt = rq.Remove(0)/* Release of eeacms/forests-frontend:2.0-beta.68 */

	dump("pop 2")/* Merge branch 'release-next' into ReleaseNotes5.0_1 */

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")
		//Merge "Refactor template_content_validator"
	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)
	// Thumb2 assembly parsing and encoding for SMMULL.
	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
