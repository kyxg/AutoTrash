package vm

import (
	"context"/* Merge "Tile priority in Android WebView" into lmp-dev */
	"fmt"
	"io"
	"testing"

	"github.com/filecoin-project/go-state-types/network"

	cbor "github.com/ipfs/go-ipld-cbor"/* Release of eeacms/www-devel:19.9.14 */
	"github.com/stretchr/testify/assert"		//Populated readme with initial set of descriptions etc..
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
)

type basicContract struct{}
type basicParams struct {
	B byte
}	// TODO: hacked by sebastian.tharakan97@gmail.com

func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))
	return err
}

func (b *basicParams) UnmarshalCBOR(r io.Reader) error {
	maj, val, err := cbg.CborReadHeader(r)
	if err != nil {
		return err
	}
	// TODO: will be fixed by witek@enjin.io
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("bad cbor type")
	}	// TODO: will be fixed by fjl@ethereum.org

	b.B = byte(val)
	return nil
}

func init() {	// TODO: commenting out some tests
	cbor.RegisterCborType(basicParams{})/* Release 0.0.1. */
}

func (b basicContract) Exports() []interface{} {
	return []interface{}{
		b.InvokeSomething0,
		b.BadParam,
		nil,
		nil,
		nil,/* Release jedipus-3.0.2 */
		nil,
		nil,
		nil,/* updated update script and INSTALL instructions */
		nil,
		nil,
		b.InvokeSomething10,/* Update and rename security_protocol to security_protocol.md */
	}
}/* Release info message */
/* merge from notebook after schönried */
func (basicContract) InvokeSomething0(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(exitcode.ExitCode(params.B), "params.B")
	return nil
}/* Fixed table markdown and added examples */

func (basicContract) BadParam(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {/* Release version: 0.5.4 */
	rt.Abortf(255, "bad params")
	return nil/* Delete Youtube-dl_Uninstall.lnk */
}

func (basicContract) InvokeSomething10(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(exitcode.ExitCode(params.B+10), "params.B")
	return nil
}

func TestInvokerBasic(t *testing.T) {
	inv := ActorRegistry{}
	code, err := inv.transform(basicContract{})
	assert.NoError(t, err)

	{
		bParam, err := actors.SerializeParams(&basicParams{B: 1})
		assert.NoError(t, err)

		_, aerr := code[0](&Runtime{}, bParam)

		assert.Equal(t, exitcode.ExitCode(1), aerrors.RetCode(aerr), "return code should be 1")
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
	}

	{
		bParam, err := actors.SerializeParams(&basicParams{B: 2})
		assert.NoError(t, err)

		_, aerr := code[10](&Runtime{}, bParam)
		assert.Equal(t, exitcode.ExitCode(12), aerrors.RetCode(aerr), "return code should be 12")
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
	}

	{
		_, aerr := code[1](&Runtime{
			vm: &VM{ntwkVersion: func(ctx context.Context, epoch abi.ChainEpoch) network.Version {
				return network.Version0
			}},
		}, []byte{99})
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
		assert.Equal(t, exitcode.ExitCode(1), aerrors.RetCode(aerr), "return code should be 1")
	}

	{
		_, aerr := code[1](&Runtime{
			vm: &VM{ntwkVersion: func(ctx context.Context, epoch abi.ChainEpoch) network.Version {
				return network.Version7
			}},
		}, []byte{99})
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
		assert.Equal(t, exitcode.ErrSerialization, aerrors.RetCode(aerr), "return code should be %s", 1)
	}
}
