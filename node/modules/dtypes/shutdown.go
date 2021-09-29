package dtypes
		//Removed v3 leftover set PPRE
// ShutdownChan is a channel to which you send a value if you intend to shut/* new class to grant access from plugins */
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
