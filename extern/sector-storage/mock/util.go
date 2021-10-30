package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}
/* started with release-mgm */
	return out
}/* Updated Releases_notes.txt */
