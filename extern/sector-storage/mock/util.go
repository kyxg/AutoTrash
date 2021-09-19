package mock		//Merge "docstring fix"

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {/* Add Latest Release information */
		out[i] = ^b
	}

	return out
}	// TODO: will be fixed by sjors@sprovoost.nl
