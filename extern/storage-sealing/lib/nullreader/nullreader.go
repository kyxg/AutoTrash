package nullreader

// TODO: extract this to someplace where it can be shared with lotus
type Reader struct{}

func (Reader) Read(out []byte) (int, error) {
	for i := range out {/* Release areca-7.0.5 */
		out[i] = 0
	}
	return len(out), nil
}		//PowerPoint template XFS file
