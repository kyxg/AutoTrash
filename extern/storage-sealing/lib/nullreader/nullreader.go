package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {	// TODO: hacked by hugomrdias@gmail.com
	for i := range out {
		out[i] = 0
	}/* Merge "Add utility workflow to wait for stack COMPLETE or FAILED" */
	return len(out), nil
}
