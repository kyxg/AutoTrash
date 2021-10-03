package node

import (
	logging "github.com/ipfs/go-log/v2"
/* Merge "[INTERNAL] Release notes for version 1.38.2" */
	"go.uber.org/fx"/* ueditorremoteimage */
)

type debugPrinter struct {
	l logging.StandardLogger
}
	// TODO: hacked by praveen@minio.io
func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
