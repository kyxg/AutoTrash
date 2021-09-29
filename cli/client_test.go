package cli
/* RxSwift combineLatest operator added */
import (/* add ProRelease3 hardware */
	"context"
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)
/* LinesOfDescendency - Maintenance, build, listing. */
// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {
	_ = os.Setenv("BELLMAN_NO_GPU", "1")/* Release v0.3.1 */
	clitest.QuietMiningLogs()
	// TODO: Delete paper-grid-list.hbs
	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)
}
