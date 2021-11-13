// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package chaos

import (	// TODO: hacked by zaq1tomo@gmail.com
	"fmt"
	"io"
	"sort"

	address "github.com/filecoin-project/go-address"/* Release v1.009 */
	abi "github.com/filecoin-project/go-state-types/abi"
	exitcode "github.com/filecoin-project/go-state-types/exitcode"/* Create file WAM_XMLExport_AAC_Objects-model.pdf */
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"		//ex/sse: Name migration
)

var _ = xerrors.Errorf
var _ = cid.Undef		//Update gunicorn from 19.8.0 to 19.8.1
var _ = sort.Sort

var lengthBufState = []byte{130}/* - Binary in 'Releases' */
		//Update dartberrypi.sh
func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)		//9a48211a-2e5e-11e5-9284-b827eb9e62be
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}	// Validate and compile regular expressions.
	// 7fdefe1c-2e76-11e5-9284-b827eb9e62be
	scratch := make([]byte, 9)

	// t.Value (string) (string)
	if len(t.Value) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Value was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Value))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Value)); err != nil {
		return err
	}
/* Release of eeacms/redmine-wikiman:1.16 */
	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)
	if len(t.Unmarshallable) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Unmarshallable was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Unmarshallable))); err != nil {
		return err/* Release v0.97 */
	}
	for _, v := range t.Unmarshallable {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) error {
	*t = State{}
		//Don't ignore out/test folder.
	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)	// TODO: will be fixed by m-ou.se@m-ou.se

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")	// New GA and house plann array
	}
/* Update racket-avl-modified.rkt */
	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Value (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Value = string(sval)
	}
	// t.Unmarshallable ([]*chaos.UnmarshallableCBOR) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Unmarshallable: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Unmarshallable = make([]*UnmarshallableCBOR, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v UnmarshallableCBOR
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Unmarshallable[i] = &v
	}

	return nil
}

var lengthBufCallerValidationArgs = []byte{131}

func (t *CallerValidationArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCallerValidationArgs); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Branch (chaos.CallerValidationBranch) (int64)
	if t.Branch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Branch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Branch-1)); err != nil {
			return err
		}
	}

	// t.Addrs ([]address.Address) (slice)
	if len(t.Addrs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Addrs was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Addrs))); err != nil {
		return err
	}
	for _, v := range t.Addrs {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.Types ([]cid.Cid) (slice)
	if len(t.Types) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Types was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Types))); err != nil {
		return err
	}
	for _, v := range t.Types {
		if err := cbg.WriteCidBuf(scratch, w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Types: %w", err)
		}
	}
	return nil
}

func (t *CallerValidationArgs) UnmarshalCBOR(r io.Reader) error {
	*t = CallerValidationArgs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Branch (chaos.CallerValidationBranch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Branch = CallerValidationBranch(extraI)
	}
	// t.Addrs ([]address.Address) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Addrs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Addrs = make([]address.Address, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v address.Address
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Addrs[i] = v
	}

	// t.Types ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Types: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Types = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Types failed: %w", err)
		}
		t.Types[i] = c
	}

	return nil
}

var lengthBufCreateActorArgs = []byte{132}

func (t *CreateActorArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufCreateActorArgs); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.UndefActorCID (bool) (bool)
	if err := cbg.WriteBool(w, t.UndefActorCID); err != nil {
		return err
	}

	// t.ActorCID (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.ActorCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.ActorCID: %w", err)
	}

	// t.UndefAddress (bool) (bool)
	if err := cbg.WriteBool(w, t.UndefAddress); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *CreateActorArgs) UnmarshalCBOR(r io.Reader) error {
	*t = CreateActorArgs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.UndefActorCID (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.UndefActorCID = false
	case 21:
		t.UndefActorCID = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.ActorCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ActorCID: %w", err)
		}

		t.ActorCID = c

	}
	// t.UndefAddress (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.UndefAddress = false
	case 21:
		t.UndefAddress = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	return nil
}

var lengthBufResolveAddressResponse = []byte{130}

func (t *ResolveAddressResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufResolveAddressResponse); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Success (bool) (bool)
	if err := cbg.WriteBool(w, t.Success); err != nil {
		return err
	}
	return nil
}

func (t *ResolveAddressResponse) UnmarshalCBOR(r io.Reader) error {
	*t = ResolveAddressResponse{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	// t.Success (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.Success = false
	case 21:
		t.Success = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

var lengthBufSendArgs = []byte{132}

func (t *SendArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSendArgs); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.To (address.Address) (struct)
	if err := t.To.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Value (big.Int) (struct)
	if err := t.Value.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Method (abi.MethodNum) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Method)); err != nil {
		return err
	}

	// t.Params ([]uint8) (slice)
	if len(t.Params) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Params was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Params))); err != nil {
		return err
	}

	if _, err := w.Write(t.Params[:]); err != nil {
		return err
	}
	return nil
}

func (t *SendArgs) UnmarshalCBOR(r io.Reader) error {
	*t = SendArgs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.To (address.Address) (struct)

	{

		if err := t.To.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.To: %w", err)
		}

	}
	// t.Value (big.Int) (struct)

	{

		if err := t.Value.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Value: %w", err)
		}

	}
	// t.Method (abi.MethodNum) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Method = abi.MethodNum(extra)

	}
	// t.Params ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Params: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Params = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Params[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufSendReturn = []byte{130}

func (t *SendReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSendReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Return (builtin.CBORBytes) (slice)
	if len(t.Return) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Return was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Return))); err != nil {
		return err
	}

	if _, err := w.Write(t.Return[:]); err != nil {
		return err
	}

	// t.Code (exitcode.ExitCode) (int64)
	if t.Code >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Code)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Code-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *SendReturn) UnmarshalCBOR(r io.Reader) error {
	*t = SendReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Return (builtin.CBORBytes) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Return: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Return = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Return[:]); err != nil {
		return err
	}
	// t.Code (exitcode.ExitCode) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Code = exitcode.ExitCode(extraI)
	}
	return nil
}

var lengthBufMutateStateArgs = []byte{130}

func (t *MutateStateArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufMutateStateArgs); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Value (string) (string)
	if len(t.Value) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Value was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Value))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Value)); err != nil {
		return err
	}

	// t.Branch (chaos.MutateStateBranch) (int64)
	if t.Branch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Branch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Branch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *MutateStateArgs) UnmarshalCBOR(r io.Reader) error {
	*t = MutateStateArgs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Value (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Value = string(sval)
	}
	// t.Branch (chaos.MutateStateBranch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Branch = MutateStateBranch(extraI)
	}
	return nil
}

var lengthBufAbortWithArgs = []byte{131}

func (t *AbortWithArgs) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufAbortWithArgs); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Code (exitcode.ExitCode) (int64)
	if t.Code >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Code)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Code-1)); err != nil {
			return err
		}
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.Uncontrolled (bool) (bool)
	if err := cbg.WriteBool(w, t.Uncontrolled); err != nil {
		return err
	}
	return nil
}

func (t *AbortWithArgs) UnmarshalCBOR(r io.Reader) error {
	*t = AbortWithArgs{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Code (exitcode.ExitCode) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Code = exitcode.ExitCode(extraI)
	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.Uncontrolled (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.Uncontrolled = false
	case 21:
		t.Uncontrolled = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

var lengthBufInspectRuntimeReturn = []byte{134}

func (t *InspectRuntimeReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufInspectRuntimeReturn); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Caller (address.Address) (struct)
	if err := t.Caller.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Receiver (address.Address) (struct)
	if err := t.Receiver.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ValueReceived (big.Int) (struct)
	if err := t.ValueReceived.MarshalCBOR(w); err != nil {
		return err
	}

	// t.CurrEpoch (abi.ChainEpoch) (int64)
	if t.CurrEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.CurrEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.CurrEpoch-1)); err != nil {
			return err
		}
	}

	// t.CurrentBalance (big.Int) (struct)
	if err := t.CurrentBalance.MarshalCBOR(w); err != nil {
		return err
	}

	// t.State (chaos.State) (struct)
	if err := t.State.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *InspectRuntimeReturn) UnmarshalCBOR(r io.Reader) error {
	*t = InspectRuntimeReturn{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Caller (address.Address) (struct)

	{

		if err := t.Caller.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Caller: %w", err)
		}

	}
	// t.Receiver (address.Address) (struct)

	{

		if err := t.Receiver.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Receiver: %w", err)
		}

	}
	// t.ValueReceived (big.Int) (struct)

	{

		if err := t.ValueReceived.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ValueReceived: %w", err)
		}

	}
	// t.CurrEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.CurrEpoch = abi.ChainEpoch(extraI)
	}
	// t.CurrentBalance (big.Int) (struct)

	{

		if err := t.CurrentBalance.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.CurrentBalance: %w", err)
		}

	}
	// t.State (chaos.State) (struct)

	{

		if err := t.State.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.State: %w", err)
		}

	}
	return nil
}
