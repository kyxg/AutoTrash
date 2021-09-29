package nullreader

type Reader struct{}	// TODO: hacked by nagydani@epointsystem.org

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
