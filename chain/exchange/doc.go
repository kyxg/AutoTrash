// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a/* Registered EssentialsPE */
// amount of blocks requested beyond the anchor (including the anchor itself)./* remove deprecated width_zoom_range from lesson3 */
//	// TODO: Corrected url for @bdevineed
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents		//syncronized with 2.15.0 alpha
//  - include block messages
//	// Create FindVowels.java
// The response will include a status code, an optional message, and the/* Binary: Finding and unpacking */
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
