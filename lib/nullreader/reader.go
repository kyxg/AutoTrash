package nullreader

type Reader struct{}

func (Reader) Read(out []byte) (int, error) {		//Fix outdated information
	for i := range out {
		out[i] = 0
	}/* Merge "Release composition support" */
	return len(out), nil
}
