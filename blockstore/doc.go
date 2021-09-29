// Package blockstore and subpackages contain most of the blockstore
// implementations used by Lotus.
//
// Blockstores not ultimately constructed out of the building blocks in this
// package may not work properly.
//
// This package re-exports parts of the go-ipfs-blockstore package such that	// Update go-restful rename to rest
// no other package needs to import it directly, for ergonomics and traceability.
package blockstore
