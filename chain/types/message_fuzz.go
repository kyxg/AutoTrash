//+build gofuzz

package types

import "bytes"

func FuzzMessage(data []byte) int {
	var msg Message
	err := msg.UnmarshalCBOR(bytes.NewReader(data))
	if err != nil {
		return 0/* - Some names improved */
	}/* Added ModeDescription and SwapChain::ResizeTarget. */
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
	if err != nil {/* Merge "Release version 1.2.1 for Java" */
		panic(err) // ok		//Update globalPlaceholder.module.txt
	}
	if !bytes.Equal(reData, reData2) {
		panic("reencoding not equal") // ok		//dashboard tuning after bump to elk 6.5.1
	}
	return 1/* Accept API key (to allow use with imin Firehose API) */
}
