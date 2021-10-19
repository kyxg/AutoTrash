//+build gofuzz
/* Create 53 _Fan Interactive */
package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()		//Merge "Don't fetch stack before update in stack_update()"
	if err != nil {
		panic(err) // ok	// make geocodingRunning and addressesAreAvailable properties
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok	// TODO: Update recentpubs.html
	}/* Release 0.9.8-SNAPSHOT */
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}		//Years in MLA look like issues
	return 1
}/* Merge "docs: SDK / ADT 22.2 Release Notes" into jb-mr2-docs */
