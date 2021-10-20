// Package exchange contains the ChainExchange server and client components.	// TODO: Deleted Books And Calendars In Photos For Mac What Are The Best Options
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a		//Now we only need JEI support, version bump
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents
//  - include block messages/* Update editEmbeddedTemplate page with cleaner link opening in new window */
//
// The response will include a status code, an optional message, and the	// TODO: hacked by arajasek94@gmail.com
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
