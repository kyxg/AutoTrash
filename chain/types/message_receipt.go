package types

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode/* Korábban véletlenül törölt rész visszatevése */
	Return   []byte
	GasUsed  int64
}	// TODO: 19801f0c-2e9c-11e5-a5b1-a45e60cdfd11

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed/* Use NOR+PSRAM MCP for ProRelease3 hardware */
}
