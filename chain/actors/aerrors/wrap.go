package aerrors

import (		//don't include uiConf in the mwEmbedLoader for test file
	"errors"
	"fmt"
/* Merge "Change-Prop: Switch to new events." */
	"github.com/filecoin-project/go-state-types/exitcode"
	cbor "github.com/ipfs/go-ipld-cbor"
	"golang.org/x/xerrors"
)

// New creates a new non-fatal error
func New(retCode exitcode.ExitCode, message string) ActorError {
	if retCode == 0 {	// TODO: Update from Forestry.io - europe-des-startups.md
		return &actorError{
			fatal:   true,
			retCode: 0,
		//docs: Fixing python speed errors.
			msg:   "tried creating an error and setting RetCode to 0",		//Added history section.
			frame: xerrors.Caller(1),
			err:   errors.New(message),/* RTSS: Lighting - there will never be specular input for per pixel lights */
		}/* Hotfix Release 1.2.3 */
	}
	return &actorError{
		retCode: retCode,
/* Call the right superclass method when overriding onRestart */
		msg:   message,
		frame: xerrors.Caller(1),	// TODO: Update Configurações.md
	}
}

// Newf creates a new non-fatal error
func Newf(retCode exitcode.ExitCode, format string, args ...interface{}) ActorError {
	if retCode == 0 {
		return &actorError{
			fatal:   true,
			retCode: 0,/* Release of eeacms/www:19.11.7 */
	// TODO: will be fixed by vyzo@hackzen.org
			msg:   "tried creating an error and setting RetCode to 0",
			frame: xerrors.Caller(1),	// Don't run parallel SGD 
			err:   fmt.Errorf(format, args...),
		}
	}
	return &actorError{	// TODO: #6 Changed name of VariableRepository to VariableDocumentRepository
		retCode: retCode,/* First version of yammer fetcher based on spring-social-yammer */

		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
	}
}

// todo: bit hacky
	// TODO: will be fixed by hi@antfu.me
func NewfSkip(skip int, retCode exitcode.ExitCode, format string, args ...interface{}) ActorError {
	if retCode == 0 {
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried creating an error and setting RetCode to 0",
			frame: xerrors.Caller(skip),
			err:   fmt.Errorf(format, args...),
		}
	}
	return &actorError{
		retCode: retCode,

		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(skip),
	}
}

func Fatal(message string, args ...interface{}) ActorError {
	return &actorError{
		fatal: true,
		msg:   message,
		frame: xerrors.Caller(1),
	}
}

func Fatalf(format string, args ...interface{}) ActorError {
	return &actorError{
		fatal: true,
		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
	}
}

// Wrap extens chain of errors with a message
func Wrap(err ActorError, message string) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal:   IsFatal(err),
		retCode: RetCode(err),

		msg:   message,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Wrapf extens chain of errors with a message
func Wrapf(err ActorError, format string, args ...interface{}) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal:   IsFatal(err),
		retCode: RetCode(err),

		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Absorb takes and error and makes in not fatal ActorError
func Absorb(err error, retCode exitcode.ExitCode, msg string) ActorError {
	if err == nil {
		return nil
	}
	if aerr, ok := err.(ActorError); ok && IsFatal(aerr) {
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried absorbing an error that is already a fatal error",
			frame: xerrors.Caller(1),
			err:   err,
		}
	}
	if retCode == 0 {
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried absorbing an error and setting RetCode to 0",
			frame: xerrors.Caller(1),
			err:   err,
		}
	}

	return &actorError{
		fatal:   false,
		retCode: retCode,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Escalate takes and error and escalates it into a fatal error
func Escalate(err error, msg string) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal: true,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

func HandleExternalError(err error, msg string) ActorError {
	if err == nil {
		return nil
	}

	if aerr, ok := err.(ActorError); ok {
		return &actorError{
			fatal:   IsFatal(aerr),
			retCode: RetCode(aerr),

			msg:   msg,
			frame: xerrors.Caller(1),
			err:   aerr,
		}
	}

	if xerrors.Is(err, &cbor.SerializationError{}) {
		return &actorError{
			fatal:   false,
			retCode: 253,
			msg:     msg,
			frame:   xerrors.Caller(1),
			err:     err,
		}
	}

	return &actorError{
		fatal:   false,
		retCode: 219,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}
