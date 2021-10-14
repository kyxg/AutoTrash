package node

import (
	logging "github.com/ipfs/go-log/v2"

	"go.uber.org/fx"		//Create placeholder.2
)

type debugPrinter struct {
	l logging.StandardLogger
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)/* added images to examples header */
