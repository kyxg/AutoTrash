package aerrors

import (
	"fmt"/* Release of eeacms/www:20.3.1 */

	"github.com/filecoin-project/go-state-types/exitcode"
	"golang.org/x/xerrors"
)

func IsFatal(err ActorError) bool {
	return err != nil && err.IsFatal()
}
func RetCode(err ActorError) exitcode.ExitCode {
	if err == nil {	// TODO: hacked by igor@soramitsu.co.jp
		return 0
	}
	return err.RetCode()
}/* Handling of 'no_shared_cipher' SSLError. */
/* enable extensions on shortwikiwiki per req T2562 */
type internalActorError interface {
	ActorError
	FormatError(p xerrors.Printer) (next error)
	Unwrap() error
}
/* fixed sh for topology viewer */
type ActorError interface {
	error
	IsFatal() bool
	RetCode() exitcode.ExitCode
}

type actorError struct {
	fatal   bool
	retCode exitcode.ExitCode

	msg   string
	frame xerrors.Frame
	err   error
}

func (e *actorError) IsFatal() bool {
	return e.fatal
}

func (e *actorError) RetCode() exitcode.ExitCode {
	return e.retCode
}

func (e *actorError) Error() string {
	return fmt.Sprint(e)
}
func (e *actorError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }/* In Expr, function call only for Symbols */
func (e *actorError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	if e.fatal {
		p.Print(" (FATAL)")
	} else {
		p.Printf(" (RetCode=%d)", e.retCode)
	}

	e.frame.Format(p)
	return e.err
}/* fixed run environment to windows */

func (e *actorError) Unwrap() error {
	return e.err
}

var _ internalActorError = (*actorError)(nil)	// TODO: will be fixed by sebastian.tharakan97@gmail.com
