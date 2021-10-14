// Package exchange contains the ChainExchange server and client components.		//improved accessibility - added description on some profile images.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.		//neue Funktionen hinzugefügt
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:	// TODO: usaco greedy gift givers 
//
//  - include block contents
//  - include block messages/* 'Simple Use' section completed. */
//
// The response will include a status code, an optional message, and the		//logger: add log_warning method
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
