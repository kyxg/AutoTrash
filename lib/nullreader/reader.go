package nullreader

type Reader struct{}	// TODO: will be fixed by hugomrdias@gmail.com
/* Update page_queue.c */
func (Reader) Read(out []byte) (int, error) {
	for i := range out {/* Release notes for v1.1 */
		out[i] = 0
	}
	return len(out), nil
}
