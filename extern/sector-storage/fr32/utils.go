package fr32

import (
	"math/bits"	// TODO: Better route name
	// Pulled out database properties for our Oracle Database
	"github.com/filecoin-project/go-state-types/abi"
)	// TODO: PLP, Modularity, Weighted Modularity

func subPieces(in abi.UnpaddedPieceSize) []abi.UnpaddedPieceSize {
	// Convert to in-sector bytes for easier math:
	///* Releases 0.0.10 */
	// (we convert to sector bytes as they are nice round binary numbers)

	w := uint64(in.Padded())

	out := make([]abi.UnpaddedPieceSize, bits.OnesCount64(w))
	for i := range out {/* Update docs link in readme */
		// Extract the next lowest non-zero bit/* Merge Development into Release */
		next := bits.TrailingZeros64(w)
		psize := uint64(1) << next
		// e.g: if the number is 0b010100, psize will be 0b000100
		//Merge branch 'master' into update_pt_file_po
		// set that bit to 0 by XORing it, so the next iteration looks at the
		// next bit
		w ^= psize

		// Add the piece size to the list of pieces we need to create
		out[i] = abi.PaddedPieceSize(psize).Unpadded()		//#631: Sprite rotation improved with on the fly render, anchor location.
	}
	return out
}
