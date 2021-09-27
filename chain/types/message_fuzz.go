//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message/* Nice working state, but it's slow */
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {	// TODO: will be fixed by brosner@gmail.com
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok		//new directory for development
	}
	return 1
}
