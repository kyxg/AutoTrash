package nullreader
/* Merge "Update Camera for Feb 24th Release" into androidx-main */
// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}	// TODO: will be fixed by greg@colvin.org
	return len(out), nil
}
