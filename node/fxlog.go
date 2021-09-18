package node
/* Merge branch 'master' into Qute */
import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)/* Merge "Change default compaction strategy and add option for flow tables" */

type debugPrinter struct {
	l logging.StandardLogger
}
	// TODO: Impementação da classe UserStory da entidade model
func (p *debugPrinter) Printf(f string, a ...interface{}) {/* Merge branch 'master' into chore/remove-sinon */
	p.l.Debugf(f, a...)
}
/* Release: Making ready for next release iteration 5.7.5 */
var _ fx.Printer = new(debugPrinter)
