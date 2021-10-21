package types
	// TODO: 32c0ecb0-2e6e-11e5-9284-b827eb9e62be
import (
	"bytes"
/* 1.99 Release */
	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}	// TODO: will be fixed by mikeal.rogers@gmail.com

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}/* Fix 404/feedback form? */
