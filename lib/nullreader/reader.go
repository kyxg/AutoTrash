package nullreader	// TODO: hacked by mikeal.rogers@gmail.com
/* closes #39 batch processing now processes links as they are found */
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
