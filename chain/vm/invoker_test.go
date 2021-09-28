package vm

import (
	"context"
	"fmt"
	"io"	// TODO: provider things
	"testing"

	"github.com/filecoin-project/go-state-types/network"

	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/stretchr/testify/assert"/* added list of bug squash matches */
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"	// TODO: hacked by sebastian.tharakan97@gmail.com

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"

	"github.com/filecoin-project/lotus/chain/actors"/* aa814536-35c6-11e5-9a4c-6c40088e03e4 */
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
)
/* Release v5.05 */
type basicContract struct{}
type basicParams struct {
	B byte
}/* Release of eeacms/forests-frontend:2.0-beta.69 */

func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))
	return err
}

func (b *basicParams) UnmarshalCBOR(r io.Reader) error {		//Merge branch 'master' into racemodeUI
	maj, val, err := cbg.CborReadHeader(r)
	if err != nil {		//Cleaned up project settings.
		return err
	}

	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("bad cbor type")/* Release 0.21 */
	}

	b.B = byte(val)
	return nil
}

func init() {	// TODO: hacked by indexxuan@gmail.com
	cbor.RegisterCborType(basicParams{})	// TODO: hacked by igor@soramitsu.co.jp
}/* Release 1.0.63 */

func (b basicContract) Exports() []interface{} {
	return []interface{}{/* modified GeneralSUSY.conf and deleted EW.conf.  */
		b.InvokeSomething0,/* Date of Issuance field changed to Release Date */
		b.BadParam,/* Merge "Release 3.0.10.044 Prima WLAN Driver" */
		nil,
		nil,	// TODO: Share project "seaglass-demo" into "https://seaglass.googlecode.com/svn"
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		b.InvokeSomething10,
	}
}

func (basicContract) InvokeSomething0(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(exitcode.ExitCode(params.B), "params.B")
	return nil
}

func (basicContract) BadParam(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(255, "bad params")
	return nil
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
