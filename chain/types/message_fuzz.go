//+build gofuzz

package types
		//Delete AZ.DEEN.txt.7z
import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message/* Support custom fields in commits. */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()/* Update Sign.md */
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))/* First Demo Ready Release */
	if err != nil {
		panic(err) // ok
	}	// Päivitetty otsikot
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1/* Merged branch development into Release */
}
