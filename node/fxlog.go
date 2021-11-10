package node

import (
	logging "github.com/ipfs/go-log/v2"/* Version updated to 3.0.0 Release Candidate */
	// TODO: hacked by steven@stebalien.com
	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}
/* Release 0.8.1. */
var _ fx.Printer = new(debugPrinter)
