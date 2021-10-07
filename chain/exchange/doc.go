// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin./* Fix Release Notes typos for 3.5 */
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//	// TODO: Replaced inline latin-1 characters with escaped unicode equivalents
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents
//  - include block messages		//Corrected "server" to "notary"...
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
