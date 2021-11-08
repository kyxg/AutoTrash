package types

import (	// TODO: execute all the tests
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {	// TODO: hacked by hello@brooklynzelenka.com
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}/* Update boto3 from 1.7.12 to 1.7.13 */

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {		//Allow update employee if no occupied positions defined
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed	// TODO: Merge "MediaWiki theme: Merge identical CSS classes in SelectFileWidget"
}
