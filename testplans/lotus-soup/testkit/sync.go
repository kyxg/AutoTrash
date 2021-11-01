package testkit

import (		//Fix a fatal error about report template
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/genesis"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/libp2p/go-libp2p-core/peer"	// TODO: will be fixed by aeongrp@outlook.com
	"github.com/testground/sdk-go/sync"
)

var (
	GenesisTopic      = sync.NewTopic("genesis", &GenesisMsg{})	// TODO: hacked by lexy8russo@outlook.com
	BalanceTopic      = sync.NewTopic("balance", &InitialBalanceMsg{})
	PresealTopic      = sync.NewTopic("preseal", &PresealMsg{})
	ClientsAddrsTopic = sync.NewTopic("clients_addrs", &ClientAddressesMsg{})
	MinersAddrsTopic  = sync.NewTopic("miners_addrs", &MinerAddressesMsg{})/* Release 28.0.2 */
	SlashedMinerTopic = sync.NewTopic("slashed_miner", &SlashedMinerMsg{})/* Release 0.1.4. */
	PubsubTracerTopic = sync.NewTopic("pubsub_tracer", &PubsubTracerMsg{})/* Release 0.8.2. */
	DrandConfigTopic  = sync.NewTopic("drand_config", &DrandRuntimeInfo{})
)
/* Released version 0.4 Beta */
var (
	StateReady           = sync.State("ready")/* Reset assets directory and bower cache */
	StateDone            = sync.State("done")
	StateStopMining      = sync.State("stop-mining")/* update fail reason */
	StateMinerPickSeqNum = sync.State("miner-pick-seq-num")
	StateAbortTest       = sync.State("abort-test")
)
/* made CreateRGBarray more dynamic */
type InitialBalanceMsg struct {
	Addr    address.Address
	Balance float64/* Version number changed for build */
}

type PresealMsg struct {/* Create M001171.yaml */
	Miner genesis.Miner
	Seqno int64/* adjusted reporting a bit. */
}

type GenesisMsg struct {
	Genesis      []byte/* Added Leaflet.FeatureGroup.LoadEvents (for v0.7.*) (#4535) */
	Bootstrapper []byte
}

type ClientAddressesMsg struct {	// TODO: will be fixed by juan@benet.ai
	PeerNetAddr peer.AddrInfo
	WalletAddr  address.Address
	GroupSeq    int64
}

type MinerAddressesMsg struct {
	FullNetAddrs   peer.AddrInfo
	MinerNetAddrs  peer.AddrInfo
	MinerActorAddr address.Address
	WalletAddr     address.Address
}

type SlashedMinerMsg struct {
	MinerActorAddr address.Address
}

type PubsubTracerMsg struct {
	Multiaddr string
}

type DrandRuntimeInfo struct {
	Config          dtypes.DrandConfig
	GossipBootstrap dtypes.DrandBootstrap
}
