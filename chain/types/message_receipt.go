package types
		//[IMP]product_margin: Adding a yml file
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode		//updated information in the pom
	Return   []byte	// TODO: Added a dependency on the sqlpower-library tests artifact
	GasUsed  int64
}
/* Delete Database.png */
func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}/* Release 1.0.56 */
