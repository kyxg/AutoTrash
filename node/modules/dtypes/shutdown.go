package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.	// TODO: Merge "ASoC: wcd: enable impedance detection."
type ShutdownChan chan struct{}	// TODO: Added complete discrete filtering to quantum driver. [Couriersud]
