package types
		//Create ctime.sh
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
/* Profile suffix added */
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {/* Release version [10.7.0] - prepare */
		ae.AppendString(c.String())		//cdi bugfix
	}
	return nil/* Release 0.0.27 */
}
