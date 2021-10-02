package node

import (
	logging "github.com/ipfs/go-log/v2"
		//fix #24, fix 26, extra vertalingen #22
	"go.uber.org/fx"
)

type debugPrinter struct {
	l logging.StandardLogger
}/* Version 3.0 Release */

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)		//Farben und Header
}

var _ fx.Printer = new(debugPrinter)
