package node/* Release script: added Dockerfile(s) */

import (
	logging "github.com/ipfs/go-log/v2"
/* Release areca-7.1 */
	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
