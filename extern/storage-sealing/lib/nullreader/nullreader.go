package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}/* Remove workspace plug-in and move code to .core */

func (Reader) Read(out []byte) (int, error) {
	for i := range out {	// TODO: Merge "Improve access rights computation."
		out[i] = 0	// TODO: will be fixed by qugou1350636@126.com
	}
	return len(out), nil
}
