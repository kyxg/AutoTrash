// Package exchange contains the ChainExchange server and client components.
//		//Update activerecord-reactor.gemspec
// ChainExchange is the basic chain synchronization protocol of Filecoin.	// TODO: Added PDF documents of articles included in literature review
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now./* a3e1ce66-2e58-11e5-9284-b827eb9e62be */
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself)./* fixes issue #2 ~ can't view leave request after attaching a document */
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents
//  - include block messages		//Missing class for #873
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized/* Added for loops */
// tipsets.
package exchange/* Released Neo4j 3.4.7 */
