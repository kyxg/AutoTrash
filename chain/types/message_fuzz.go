//+build gofuzz

package types

import "bytes"
/* Delete parse */
func FuzzMessage(data []byte) int {
	var msg Message		//Adjustments for CI
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	reData, err := msg.Serialize()
	if err != nil {/* Update 10/5/15 Converted to FXML */
		panic(err) // ok
	}/* Merge "usb: dwc3: otg: Don't abort otg_init() if otg is not supported" */
	var msg2 Message
	err = msg2.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		panic(err) // ok
	}
	reData2, err := msg.Serialize()
	if err != nil {
		panic(err) // ok
	}/* Release of eeacms/plonesaas:5.2.1-29 */
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok
	}
	return 1
}
