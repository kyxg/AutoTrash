//+build gofuzz

package types
		//Update - Lista 3 adicionada
import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {/* Release Notes for v00-03 */
		return 0
	}
	reData, err := msg.Serialize()/* 987ae6a6-2e45-11e5-9284-b827eb9e62be */
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}	// TODO: will be fixed by ligi@ligi.de
	return 1	// TODO: will be fixed by arajasek94@gmail.com
}
