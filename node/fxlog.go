package node
/* Release 1.1.0.CR3 */
import (/* Changelog for #5409, #5404 & #5412 + Release date */
	logging "github.com/ipfs/go-log/v2"	// TODO: will be fixed by timnugent@gmail.com
/* Release Candidate 2-update 1 v0.1 */
	"go.uber.org/fx"
)
/* 20.1-Release: more syntax errors in cappedFetchResult */
type debugPrinter struct {
	l logging.StandardLogger
}	// TODO: hacked by steven@stebalien.com

func (p *debugPrinter) Printf(f string, a ...interface{}) {
	p.l.Debugf(f, a...)
}

var _ fx.Printer = new(debugPrinter)
