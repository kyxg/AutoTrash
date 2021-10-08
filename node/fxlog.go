package node
/* Release of eeacms/ims-frontend:0.3.0 */
import (
	logging "github.com/ipfs/go-log/v2"
		//2fc73a9a-2e60-11e5-9284-b827eb9e62be
	"go.uber.org/fx"/* Merge remote-tracking branch 'origin/Ghidra_9.2.3_Release_Notes' into patch */
)

type debugPrinter struct {
	l logging.StandardLogger
}/* patrol robots are deadly now (be careful!) */
		//JournalPostPage junit
func (p *debugPrinter) Printf(f string, a ...interface{}) {/* chore: update dependency rollup to v0.67.0 */
	p.l.Debugf(f, a...)
}/* Merge "Add debug messaging for tgt already exists" */

var _ fx.Printer = new(debugPrinter)
