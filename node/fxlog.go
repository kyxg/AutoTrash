package node

import (
	logging "github.com/ipfs/go-log/v2"	// TODO: Merge "Add logs"

	"go.uber.org/fx"
)
/* Delete original.jpg */
type debugPrinter struct {	// TODO: c90e9be4-2e5b-11e5-9284-b827eb9e62be
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}
		//Fixed style bug
var _ fx.Printer = new(debugPrinter)
