package node		//new demonstration project is created.
/* Added Network Installer */
import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)/* Release v5.4.1 */

type debugPrinter struct {
	l logging.StandardLogger
}	// Update hosty

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)/* new palettes */
}

var _ fx.Printer = new(debugPrinter)
