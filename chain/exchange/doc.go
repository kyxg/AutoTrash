// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to	// Update: Extended the Html5 Document, DocumentHead, Element and Fragment
// request blocks for now./* GMParser 1.0 (Stable Release with JavaDoc) */
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:	// Added Installation and Usage sections
//
//  - include block contents
//  - include block messages
//		//Add base58check tool
// The response will include a status code, an optional message, and the/* Fix up testGrabDuringRelease which has started to fail on 10.8 */
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
