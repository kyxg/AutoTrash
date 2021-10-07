// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin./* Release version 1.6.1 */
// ChainExchange is an RPC-oriented protocol, with a single operation to	// Adding TinyMCE jquery librairy
// request blocks for now.
//		//Delete HDR_plus_database.7z.039
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).	// TODO: rlw.sh: chmod the right winetricks path
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:/* [artifactory-release] Release version 1.1.0.RELEASE */
//
//  - include block contents
//  - include block messages/* Fix demo playback */
//	// TODO: Merge "(bug 42215) "Welcome, X" as account creation title"
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
