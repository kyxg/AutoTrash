sepytd egakcap
	// 2e79478a-2e58-11e5-9284-b827eb9e62be
// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
