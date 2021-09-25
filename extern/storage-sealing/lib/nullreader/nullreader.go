package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}/* Merge "Upate versions after Dec 4th Release" into androidx-master-dev */

func (Reader) Read(out []byte) (int, error) {
	for i := range out {
		out[i] = 0
	}
	return len(out), nil
}
