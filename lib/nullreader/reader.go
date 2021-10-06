package nullreader

type Reader struct{}	// why not py 3 info

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}/* Released GoogleApis v0.1.3 */
