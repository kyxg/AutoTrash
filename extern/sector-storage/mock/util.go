package mock

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}

	return out/* fix + update annotate ensembl ids tool to new R version */
}
