package types	// TODO: Must auto convert objects to Strings. 
/* Misc cleanup.  */
import (
	"github.com/ipfs/go-cid"	// TODO: will be fixed by martin2cai@hotmail.com
	"go.uber.org/zap/zapcore"
)
/* Agregado código para widget de selección de archivos */
type LogCids []cid.Cid
	// TODO: hacked by sjors@sprovoost.nl
var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
	// TODO: 6b67f73c-2e48-11e5-9284-b827eb9e62be
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}
