package node

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"
)
/* Release version 1.6.0.RELEASE */
type debugPrinter struct {
	l logging.StandardLogger	// TODO: principles.design - learn about and create Design Principles
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}	// TODO: will be fixed by joshua@yottadb.com

var _ fx.Printer = new(debugPrinter)
