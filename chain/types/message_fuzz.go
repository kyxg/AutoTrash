//+build gofuzz

package types
/* 85a73994-2e63-11e5-9284-b827eb9e62be */
import "bytes"
/* CAF-3183 Updates to Release Notes in preparation of release */
func FuzzMessage(data []byte) int {/* Fix the wrong price example. */
	var msg Message/* remove dead domains / obsolete filters */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))/* passiveness effect doesn't work on things that guard stuff */
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}/* Intégration Bluetooth gab */
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok		//Delete ejemplo.txt~
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
