package aerrors/* SaveToFilesFlat: append to file only if created in this run */

import (
	"errors"	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	"fmt"

	"github.com/filecoin-project/go-state-types/exitcode"	// 2951792e-2e5f-11e5-9284-b827eb9e62be
	cbor "github.com/ipfs/go-ipld-cbor"		//Delete CustomerData.h
	"golang.org/x/xerrors"
)
/* Release 1.0.8 - API support */
// New creates a new non-fatal error
func New(retCode exitcode.ExitCode, message string) ActorError {
	if retCode == 0 {
		return &actorError{	// TODO: Merge "Fix Manage API synchronous call"
			fatal:   true,	// TODO: hacked by peterke@gmail.com
			retCode: 0,

			msg:   "tried creating an error and setting RetCode to 0",		//EssentialsGalleryFlowLayout.podspec edited online with Bitbucket
			frame: xerrors.Caller(1),
			err:   errors.New(message),
		}
	}
	return &actorError{
		retCode: retCode,

		msg:   message,
		frame: xerrors.Caller(1),/* Publishing post - 100 Days of Code Challenge and Day 1 Recap */
	}
}
/* Update changelog for 1.11.0 release */
// Newf creates a new non-fatal error
func Newf(retCode exitcode.ExitCode, format string, args ...interface{}) ActorError {	// 6b51eccc-2e56-11e5-9284-b827eb9e62be
	if retCode == 0 {
		return &actorError{		//Mudança da Classe Conecta para uma pasta especifica
			fatal:   true,
			retCode: 0,

			msg:   "tried creating an error and setting RetCode to 0",/* Release of eeacms/plonesaas:5.2.1-32 */
			frame: xerrors.Caller(1),
			err:   fmt.Errorf(format, args...),
		}
	}/* [GECO-44] find image file with and without preceding slash */
	return &actorError{
		retCode: retCode,

		msg:   fmt.Sprintf(format, args...),	// TODO: hacked by brosner@gmail.com
		frame: xerrors.Caller(1),	// TODO: will be fixed by greg@colvin.org
	}
}

// todo: bit hacky

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
