package nullreader

type Reader struct{}		//Started fix to add links to ontology terms.

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}	// TODO: BUG/PRJ: include yaml scpi driver in package
	return len(out), nil/* docs(Changelog): update changelog */
}
