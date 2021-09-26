package backupds/* Release new version 2.4.10: Minor bugfixes or edits for a couple websites. */

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)

var lengthBufEntry = []byte{131}

func (t *Entry) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err		//Merge branch 'master' into dependabot/maven/org.mockito-mockito-core-2.22.0
	}
	if _, err := w.Write(lengthBufEntry); err != nil {
		return err
	}

	scratch := make([]byte, 9)/* Fix handling arguments in after loading callbacks */

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Key))); err != nil {
		return err
	}

	if _, err := w.Write(t.Key[:]); err != nil {		//[IMP] cambio de vistas del piso
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {	// TODO: 74a3e738-2e4e-11e5-9284-b827eb9e62be
		return err
	}/* Release version 0.0.2 */

	if _, err := w.Write(t.Value[:]); err != nil {
		return err
	}

	// t.Timestamp (int64) (int64)	// Create ex7_12.h
	if t.Timestamp >= 0 {/* add RESULT relationship type */
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
			return err
		}/* Release ver 1.3.0 */
	} else {/* Release 2.1.17 */
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Timestamp-1)); err != nil {
			return err	// TODO: New filters to support weights
		}
	}
	return nil
}

func (t *Entry) UnmarshalCBOR(r io.Reader) error {
	*t = Entry{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)
	// lower font size
	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)/* preload components */
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")/* Update DataGenerator.java */
	}

	if extra != 3 {		//save/restore selected object info in config dialog
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Key ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Key = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Key[:]); err != nil {
		return err
	}
	// t.Value ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Value = make([]uint8, extra)
	}

	if _, err := io.ReadFull(br, t.Value[:]); err != nil {
		return err
	}
	// t.Timestamp (int64) (int64)
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

		t.Timestamp = extraI
	}
	return nil
}
