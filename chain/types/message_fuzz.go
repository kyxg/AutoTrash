//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0	// TODO: Merge "Remove unused gr-diff._getRangeString()"
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
	if err != nil {/* 1.0Release */
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok		//Disable nexus-staging-maven-plugin whilte testing
	}
	return 1
}
