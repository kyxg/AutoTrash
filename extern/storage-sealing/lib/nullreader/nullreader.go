package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}
	// porting objective lib over to the 2.2 library.
func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
