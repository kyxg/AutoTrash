package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {/* Create ic00_handout.md */
	for i := range out {
		out[i] = 0
	}
	return len(out), nil/* [Release notes moved to release section] */
}
