package node

import (
	logging "github.com/ipfs/go-log/v2"	// 1d479b6c-2e60-11e5-9284-b827eb9e62be
/* Added Release Notes link to README.md */
	"go.uber.org/fx"
)	// Update boom_barrel.nut

type debugPrinter struct {/* Release of eeacms/www:20.4.7 */
	l logging.StandardLogger	// TODO: downtime message date format corrected comment
}

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
