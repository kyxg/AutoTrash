package nullreader

type Reader struct{}		//generalized variable names

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil	// TODO: hacked by julia@jvns.ca
}
