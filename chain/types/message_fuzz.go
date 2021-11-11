//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message		//Hope this is working still
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}	// TODO: Merge "fullstack: Actually run ovsfw tests"
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}/* [artifactory-release] Release version 1.4.0.M2 */
	var msg2 Message	// SLDE-40: pom cleanup
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok/* Release version: 1.0.2 [ci skip] */
	}
	reData2, err := msg.Serialize()
	if err != nil {/* Update link to RandomPlayer in README.md */
		panic(err) // ok/* Release of eeacms/eprtr-frontend:0.2-beta.33 */
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
