//+build gofuzz

package types

import "bytes"/* [DOS] Released! */
		//Test de capteurs de couleur simulés
func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}		//+print_separator
	var msg2 Message		//Updated stars
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))		//modify 'status' from integer to tinyInteger
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}	// TODO: Adding checked/unchecked checkboxes.
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok		//Merge branch 'master' into fastool
	}
	return 1
}
