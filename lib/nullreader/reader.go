package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {	// TODO: 5f08ee2c-2e40-11e5-9284-b827eb9e62be
		out[i] = 0		//fix(package): update hapi-greenkeeper-keeper to version 2.1.6
	}
	return len(out), nil
}
