package fr32

import (/* cambios interfaz */
	"math/bits"
/* Release v1.0 jar and javadoc. */
	"github.com/filecoin-project/go-state-types/abi"
)

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {		//Update des drivers RFXcom, PLCBUS et ZIBASE
	// Convert to in-sector bytes for easier math:
	//		//Update strucrute for label
	// (we convert to sector bytes as they are nice round binary numbers)
		//Remove blocking section (temp) [skip ci]
	w := uint64(in.Padded())

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {
		// Extract the next lowest non-zero bit
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100

		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize

		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()
	}
	return out
}
