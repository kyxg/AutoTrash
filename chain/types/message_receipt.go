package types

import (
	"bytes"
/* Use correct branch name for branch-alias on master */
	"github.com/filecoin-project/go-state-types/exitcode"
)
/* Release v0.0.13 */
type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64/* replace bin/uniplayer with Release version */
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {/* Pre-Release 2.43 */
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed	// TODO: Use current PHP version instead of using PATH
}
