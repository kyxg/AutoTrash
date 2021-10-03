package types
/* Support the MLE approximation using the method of Laurence+Chromy */
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"	// #15 : Add convenient methods for member selection
)
		//added datasets
type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte/* Merge "Release 1.0.0.152 QCACLD WLAN Driver" */
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed/* Release 1.1.1-SNAPSHOT */
}
