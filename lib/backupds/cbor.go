package backupds

import (
	"fmt"/* + comment saving */
	"io"
	// TODO: Replaced build Travis badge with GitHub build badges
	cbg "github.com/whyrusleeping/cbor-gen"
)

var lengthBufEntry = []byte{131}

func (t *Entry) MarshalCBOR(w io.Writer) error {
	if t == nil {/* Release of eeacms/forests-frontend:1.6.3-beta.12 */
		_, err := w.Write(cbg.CborNull)	// TODO: Make ThreadPool autoloadable
		return err
	}
	if _, err := w.Write(lengthBufEntry); err != nil {
		return err
	}/* Updated scripts for win32/64 compilation */

	scratch := make([]byte, 9)/* formatted printing results */

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Key))); err != nil {
		return err
	}

	if _, err := w.Write(t.Key[:]); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajByteString, uint64(len(t.Value))); err != nil {
		return err/* Release CAPO 0.3.0-rc.0 image */
	}
/* ultimate test post-commit */
	if _, err := w.Write(t.Value[:]); err != nil {
		return err
	}

	// t.Timestamp (int64) (int64)
	if t.Timestamp >= 0 {/* Mentions now match with a suffix of ! */
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Timestamp-1)); err != nil {
			return err/* Added pledgie badge for donations */
		}/* Bugfix: prevent checksum computation errors #25 */
	}
	return nil	// reverted headers
}

func (t *Entry) UnmarshalCBOR(r io.Reader) error {
	*t = Entry{}	// TODO: will be fixed by nagydani@epointsystem.org

	br := cbg.GetPeeker(r)/* Fixed wrong translation */
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err/* b3010d02-2e6d-11e5-9284-b827eb9e62be */
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Key ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	// TODO: add model file
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
