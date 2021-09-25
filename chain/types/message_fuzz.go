//+build gofuzz

package types/* file_handler: pass FileAddress to file_callback() */

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))/* Release 0.9.8 */
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok		//Delete MP_UserRolePermission.md
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}		//Merge "Un-@hide new download manager public API." into gingerbread
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}	// Delete MRodriguez1a
