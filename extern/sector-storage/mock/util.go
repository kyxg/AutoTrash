package mock

func CommDR(in []byte) (out [32]byte) {	// collect errors from the filter validations and pass them back to the report
	for i, b := range in {/* Release for v16.0.0. */
		out[i] = ^b
	}

	return out
}
