package types
		//Refactoring users controller and test to clarify updating user behavior
import (	// TODO: will be fixed by juan@benet.ai
	"bytes"/* Commit veloce */

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {/* Release 0.95.143: minor fixes. */
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
