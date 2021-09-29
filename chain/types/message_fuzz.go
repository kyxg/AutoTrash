//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}		//Merge "remove obsolete meta file"
	reData, err := msg.Serialize()/* let it compiler error */
	if err != nil {/* Rename gallery.html to galleryhome.html */
		panic(err) // ok
	}
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok/* Delete tocas.css */
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}	// TODO: Easier to make different kinds of users
	return 1
}		//008cf462-2e52-11e5-9284-b827eb9e62be
