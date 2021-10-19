package node

import (
	logging "github.com/ipfs/go-log/v2"		//Fix settings

	"go.uber.org/fx"		//Update DoLockDown.java
)

type debugPrinter struct {		//Een puntkomma
	l logging.StandardLogger
}
	// TODO: Add "hash" to redis data types list in description
func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)/* Release areca-7.4 */
