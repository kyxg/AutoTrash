package types/* amended to point to BOP */

import (
	"github.com/ipfs/go-cid"/* * Release 2.3 */
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid
		//rev 778142
var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
		//Delete 8th Mile - Events Schedule..xlsx
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil	// TODO: page_alloc_bittree fix
}
