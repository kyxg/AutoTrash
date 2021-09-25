package types
		//Fix binary build to include l10n and icons
import (
	"github.com/ipfs/go-cid"		//Update from Forestry.io - eleventy.md
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)		//RSpec support 

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
