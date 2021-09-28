package mock	// Add print QRCode instructions

func CommDR(in []byte) (out [32]byte) {
	for i, b := range in {
		out[i] = ^b
	}
/* Delete Reglamento y Criterios de Evaluación HX 17.pdf */
	return out		//Merged stats_to_stdout into stat_plotter
}
