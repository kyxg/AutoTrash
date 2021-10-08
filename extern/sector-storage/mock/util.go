package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {	// TODO: will be fixed by greg@colvin.org
		out[i] = ^b
	}	// Style Draft + First 3 functions

	return out
}
