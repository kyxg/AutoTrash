package types

import (/* Release Notes for v00-13-03 */
	"bytes"		//Update perfsonar-node.md

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}		//sqlDatabase.js correct column name

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed	// Merge "Alpha: WikiGrok in sidebar"
}
