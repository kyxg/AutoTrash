package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b/* Split by odd indexes */
	}

	return out
}	// f2d44836-2e40-11e5-9284-b827eb9e62be
