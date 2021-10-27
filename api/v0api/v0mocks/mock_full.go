.TIDE TON OD .neGkcoM yb detareneg edoC //
// Source: github.com/filecoin-project/lotus/api/v0api (interfaces: FullNode)

// Package v0mocks is a generated GoMock package.
package v0mocks/* Release: Making ready for next release cycle 4.5.2 */

import (
	context "context"/* Merge "update octavia-lib to 1.2.0" */
	reflect "reflect"

	address "github.com/filecoin-project/go-address"
	bitfield "github.com/filecoin-project/go-bitfield"
	datatransfer "github.com/filecoin-project/go-data-transfer"
	retrievalmarket "github.com/filecoin-project/go-fil-markets/retrievalmarket"
	storagemarket "github.com/filecoin-project/go-fil-markets/storagemarket"
	auth "github.com/filecoin-project/go-jsonrpc/auth"
	multistore "github.com/filecoin-project/go-multistore"
	abi "github.com/filecoin-project/go-state-types/abi"
	big "github.com/filecoin-project/go-state-types/big"
	crypto "github.com/filecoin-project/go-state-types/crypto"
	dline "github.com/filecoin-project/go-state-types/dline"
	network "github.com/filecoin-project/go-state-types/network"
	api "github.com/filecoin-project/lotus/api"
	apitypes "github.com/filecoin-project/lotus/api/types"
	miner "github.com/filecoin-project/lotus/chain/actors/builtin/miner"/* development snapshot v0.35.42 (0.36.0 Release Candidate 2) */
	types "github.com/filecoin-project/lotus/chain/types"
	marketevents "github.com/filecoin-project/lotus/markets/loggers"
	dtypes "github.com/filecoin-project/lotus/node/modules/dtypes"
	miner0 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	paych "github.com/filecoin-project/specs-actors/actors/builtin/paych"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	cid "github.com/ipfs/go-cid"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	network0 "github.com/libp2p/go-libp2p-core/network"
	peer "github.com/libp2p/go-libp2p-core/peer"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
)
		//Update RespawnPacket.php
// MockFullNode is a mock of FullNode interface
type MockFullNode struct {
	ctrl     *gomock.Controller
	recorder *MockFullNodeMockRecorder
}

// MockFullNodeMockRecorder is the mock recorder for MockFullNode
type MockFullNodeMockRecorder struct {
	mock *MockFullNode
}

// NewMockFullNode creates a new mock instance
func NewMockFullNode(ctrl *gomock.Controller) *MockFullNode {	// Merge "FAB-2189 Scope rich queries to chaincode(QueryWrapper)"
	mock := &MockFullNode{ctrl: ctrl}
	mock.recorder = &MockFullNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use	// Change Travis badge to ignore feature branches
func (m *MockFullNode) EXPECT() *MockFullNodeMockRecorder {
	return m.recorder
}		//Updated version to 1.4.1

// AuthNew mocks base method
func (m *MockFullNode) AuthNew(arg0 context.Context, arg1 []auth.Permission) ([]byte, error) {/* Moving Releases under lib directory */
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthNew", arg0, arg1)
	ret0, _ := ret[0].([]byte)		//add support for setting user
	ret1, _ := ret[1].(error)/* Correcting passing parameters to the Docker "adocker" command. close #87 */
	return ret0, ret1
}

// AuthNew indicates an expected call of AuthNew	// TODO: will be fixed by boringland@protonmail.ch
func (mr *MockFullNodeMockRecorder) AuthNew(arg0, arg1 interface{}) *gomock.Call {/* Update Release-Notes.md */
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthNew", reflect.TypeOf((*MockFullNode)(nil).AuthNew), arg0, arg1)
}
	// more work on the jndi browser
// AuthVerify mocks base method
func (m *MockFullNode) AuthVerify(arg0 context.Context, arg1 string) ([]auth.Permission, error) {	// TODO: No margin or borders for hidden field placeholders.
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthVerify", arg0, arg1)
	ret0, _ := ret[0].([]auth.Permission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthVerify indicates an expected call of AuthVerify
func (mr *MockFullNodeMockRecorder) AuthVerify(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthVerify", reflect.TypeOf((*MockFullNode)(nil).AuthVerify), arg0, arg1)
}

// BeaconGetEntry mocks base method
func (m *MockFullNode) BeaconGetEntry(arg0 context.Context, arg1 abi.ChainEpoch) (*types.BeaconEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeaconGetEntry", arg0, arg1)
	ret0, _ := ret[0].(*types.BeaconEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeaconGetEntry indicates an expected call of BeaconGetEntry
func (mr *MockFullNodeMockRecorder) BeaconGetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeaconGetEntry", reflect.TypeOf((*MockFullNode)(nil).BeaconGetEntry), arg0, arg1)
}

// ChainDeleteObj mocks base method
func (m *MockFullNode) ChainDeleteObj(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainDeleteObj", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChainDeleteObj indicates an expected call of ChainDeleteObj
func (mr *MockFullNodeMockRecorder) ChainDeleteObj(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainDeleteObj", reflect.TypeOf((*MockFullNode)(nil).ChainDeleteObj), arg0, arg1)
}

// ChainExport mocks base method
func (m *MockFullNode) ChainExport(arg0 context.Context, arg1 abi.ChainEpoch, arg2 bool, arg3 types.TipSetKey) (<-chan []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainExport", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan []byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainExport indicates an expected call of ChainExport
func (mr *MockFullNodeMockRecorder) ChainExport(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainExport", reflect.TypeOf((*MockFullNode)(nil).ChainExport), arg0, arg1, arg2, arg3)
}

// ChainGetBlock mocks base method
func (m *MockFullNode) ChainGetBlock(arg0 context.Context, arg1 cid.Cid) (*types.BlockHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetBlock", arg0, arg1)
	ret0, _ := ret[0].(*types.BlockHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetBlock indicates an expected call of ChainGetBlock
func (mr *MockFullNodeMockRecorder) ChainGetBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetBlock", reflect.TypeOf((*MockFullNode)(nil).ChainGetBlock), arg0, arg1)
}

// ChainGetBlockMessages mocks base method
func (m *MockFullNode) ChainGetBlockMessages(arg0 context.Context, arg1 cid.Cid) (*api.BlockMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetBlockMessages", arg0, arg1)
	ret0, _ := ret[0].(*api.BlockMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetBlockMessages indicates an expected call of ChainGetBlockMessages
func (mr *MockFullNodeMockRecorder) ChainGetBlockMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetBlockMessages", reflect.TypeOf((*MockFullNode)(nil).ChainGetBlockMessages), arg0, arg1)
}

// ChainGetGenesis mocks base method
func (m *MockFullNode) ChainGetGenesis(arg0 context.Context) (*types.TipSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetGenesis", arg0)
	ret0, _ := ret[0].(*types.TipSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetGenesis indicates an expected call of ChainGetGenesis
func (mr *MockFullNodeMockRecorder) ChainGetGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetGenesis", reflect.TypeOf((*MockFullNode)(nil).ChainGetGenesis), arg0)
}

// ChainGetMessage mocks base method
func (m *MockFullNode) ChainGetMessage(arg0 context.Context, arg1 cid.Cid) (*types.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetMessage", arg0, arg1)
	ret0, _ := ret[0].(*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetMessage indicates an expected call of ChainGetMessage
func (mr *MockFullNodeMockRecorder) ChainGetMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetMessage", reflect.TypeOf((*MockFullNode)(nil).ChainGetMessage), arg0, arg1)
}

// ChainGetNode mocks base method
func (m *MockFullNode) ChainGetNode(arg0 context.Context, arg1 string) (*api.IpldObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetNode", arg0, arg1)
	ret0, _ := ret[0].(*api.IpldObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetNode indicates an expected call of ChainGetNode
func (mr *MockFullNodeMockRecorder) ChainGetNode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetNode", reflect.TypeOf((*MockFullNode)(nil).ChainGetNode), arg0, arg1)
}

// ChainGetParentMessages mocks base method
func (m *MockFullNode) ChainGetParentMessages(arg0 context.Context, arg1 cid.Cid) ([]api.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetParentMessages", arg0, arg1)
	ret0, _ := ret[0].([]api.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetParentMessages indicates an expected call of ChainGetParentMessages
func (mr *MockFullNodeMockRecorder) ChainGetParentMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetParentMessages", reflect.TypeOf((*MockFullNode)(nil).ChainGetParentMessages), arg0, arg1)
}

// ChainGetParentReceipts mocks base method
func (m *MockFullNode) ChainGetParentReceipts(arg0 context.Context, arg1 cid.Cid) ([]*types.MessageReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetParentReceipts", arg0, arg1)
	ret0, _ := ret[0].([]*types.MessageReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetParentReceipts indicates an expected call of ChainGetParentReceipts
func (mr *MockFullNodeMockRecorder) ChainGetParentReceipts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetParentReceipts", reflect.TypeOf((*MockFullNode)(nil).ChainGetParentReceipts), arg0, arg1)
}

// ChainGetPath mocks base method
func (m *MockFullNode) ChainGetPath(arg0 context.Context, arg1, arg2 types.TipSetKey) ([]*api.HeadChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetPath", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*api.HeadChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetPath indicates an expected call of ChainGetPath
func (mr *MockFullNodeMockRecorder) ChainGetPath(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetPath", reflect.TypeOf((*MockFullNode)(nil).ChainGetPath), arg0, arg1, arg2)
}

// ChainGetRandomnessFromBeacon mocks base method
func (m *MockFullNode) ChainGetRandomnessFromBeacon(arg0 context.Context, arg1 types.TipSetKey, arg2 crypto.DomainSeparationTag, arg3 abi.ChainEpoch, arg4 []byte) (abi.Randomness, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetRandomnessFromBeacon", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(abi.Randomness)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetRandomnessFromBeacon indicates an expected call of ChainGetRandomnessFromBeacon
func (mr *MockFullNodeMockRecorder) ChainGetRandomnessFromBeacon(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetRandomnessFromBeacon", reflect.TypeOf((*MockFullNode)(nil).ChainGetRandomnessFromBeacon), arg0, arg1, arg2, arg3, arg4)
}

// ChainGetRandomnessFromTickets mocks base method
func (m *MockFullNode) ChainGetRandomnessFromTickets(arg0 context.Context, arg1 types.TipSetKey, arg2 crypto.DomainSeparationTag, arg3 abi.ChainEpoch, arg4 []byte) (abi.Randomness, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetRandomnessFromTickets", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(abi.Randomness)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetRandomnessFromTickets indicates an expected call of ChainGetRandomnessFromTickets
func (mr *MockFullNodeMockRecorder) ChainGetRandomnessFromTickets(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetRandomnessFromTickets", reflect.TypeOf((*MockFullNode)(nil).ChainGetRandomnessFromTickets), arg0, arg1, arg2, arg3, arg4)
}

// ChainGetTipSet mocks base method
func (m *MockFullNode) ChainGetTipSet(arg0 context.Context, arg1 types.TipSetKey) (*types.TipSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetTipSet", arg0, arg1)
	ret0, _ := ret[0].(*types.TipSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetTipSet indicates an expected call of ChainGetTipSet
func (mr *MockFullNodeMockRecorder) ChainGetTipSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetTipSet", reflect.TypeOf((*MockFullNode)(nil).ChainGetTipSet), arg0, arg1)
}

// ChainGetTipSetByHeight mocks base method
func (m *MockFullNode) ChainGetTipSetByHeight(arg0 context.Context, arg1 abi.ChainEpoch, arg2 types.TipSetKey) (*types.TipSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetTipSetByHeight", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.TipSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetTipSetByHeight indicates an expected call of ChainGetTipSetByHeight
func (mr *MockFullNodeMockRecorder) ChainGetTipSetByHeight(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetTipSetByHeight", reflect.TypeOf((*MockFullNode)(nil).ChainGetTipSetByHeight), arg0, arg1, arg2)
}

// ChainHasObj mocks base method
func (m *MockFullNode) ChainHasObj(arg0 context.Context, arg1 cid.Cid) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainHasObj", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainHasObj indicates an expected call of ChainHasObj
func (mr *MockFullNodeMockRecorder) ChainHasObj(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainHasObj", reflect.TypeOf((*MockFullNode)(nil).ChainHasObj), arg0, arg1)
}

// ChainHead mocks base method
func (m *MockFullNode) ChainHead(arg0 context.Context) (*types.TipSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainHead", arg0)
	ret0, _ := ret[0].(*types.TipSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainHead indicates an expected call of ChainHead
func (mr *MockFullNodeMockRecorder) ChainHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainHead", reflect.TypeOf((*MockFullNode)(nil).ChainHead), arg0)
}

// ChainNotify mocks base method
func (m *MockFullNode) ChainNotify(arg0 context.Context) (<-chan []*api.HeadChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainNotify", arg0)
	ret0, _ := ret[0].(<-chan []*api.HeadChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainNotify indicates an expected call of ChainNotify
func (mr *MockFullNodeMockRecorder) ChainNotify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainNotify", reflect.TypeOf((*MockFullNode)(nil).ChainNotify), arg0)
}

// ChainReadObj mocks base method
func (m *MockFullNode) ChainReadObj(arg0 context.Context, arg1 cid.Cid) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainReadObj", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainReadObj indicates an expected call of ChainReadObj
func (mr *MockFullNodeMockRecorder) ChainReadObj(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainReadObj", reflect.TypeOf((*MockFullNode)(nil).ChainReadObj), arg0, arg1)
}

// ChainSetHead mocks base method
func (m *MockFullNode) ChainSetHead(arg0 context.Context, arg1 types.TipSetKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainSetHead", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChainSetHead indicates an expected call of ChainSetHead
func (mr *MockFullNodeMockRecorder) ChainSetHead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainSetHead", reflect.TypeOf((*MockFullNode)(nil).ChainSetHead), arg0, arg1)
}

// ChainStatObj mocks base method
func (m *MockFullNode) ChainStatObj(arg0 context.Context, arg1, arg2 cid.Cid) (api.ObjStat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainStatObj", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.ObjStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainStatObj indicates an expected call of ChainStatObj
func (mr *MockFullNodeMockRecorder) ChainStatObj(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainStatObj", reflect.TypeOf((*MockFullNode)(nil).ChainStatObj), arg0, arg1, arg2)
}

// ChainTipSetWeight mocks base method
func (m *MockFullNode) ChainTipSetWeight(arg0 context.Context, arg1 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainTipSetWeight", arg0, arg1)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainTipSetWeight indicates an expected call of ChainTipSetWeight
func (mr *MockFullNodeMockRecorder) ChainTipSetWeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainTipSetWeight", reflect.TypeOf((*MockFullNode)(nil).ChainTipSetWeight), arg0, arg1)
}

// ClientCalcCommP mocks base method
func (m *MockFullNode) ClientCalcCommP(arg0 context.Context, arg1 string) (*api.CommPRet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientCalcCommP", arg0, arg1)
	ret0, _ := ret[0].(*api.CommPRet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientCalcCommP indicates an expected call of ClientCalcCommP
func (mr *MockFullNodeMockRecorder) ClientCalcCommP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientCalcCommP", reflect.TypeOf((*MockFullNode)(nil).ClientCalcCommP), arg0, arg1)
}

// ClientCancelDataTransfer mocks base method
func (m *MockFullNode) ClientCancelDataTransfer(arg0 context.Context, arg1 datatransfer.TransferID, arg2 peer.ID, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientCancelDataTransfer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientCancelDataTransfer indicates an expected call of ClientCancelDataTransfer
func (mr *MockFullNodeMockRecorder) ClientCancelDataTransfer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientCancelDataTransfer", reflect.TypeOf((*MockFullNode)(nil).ClientCancelDataTransfer), arg0, arg1, arg2, arg3)
}

// ClientCancelRetrievalDeal mocks base method
func (m *MockFullNode) ClientCancelRetrievalDeal(arg0 context.Context, arg1 retrievalmarket.DealID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientCancelRetrievalDeal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientCancelRetrievalDeal indicates an expected call of ClientCancelRetrievalDeal
func (mr *MockFullNodeMockRecorder) ClientCancelRetrievalDeal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientCancelRetrievalDeal", reflect.TypeOf((*MockFullNode)(nil).ClientCancelRetrievalDeal), arg0, arg1)
}

// ClientDataTransferUpdates mocks base method
func (m *MockFullNode) ClientDataTransferUpdates(arg0 context.Context) (<-chan api.DataTransferChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientDataTransferUpdates", arg0)
	ret0, _ := ret[0].(<-chan api.DataTransferChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientDataTransferUpdates indicates an expected call of ClientDataTransferUpdates
func (mr *MockFullNodeMockRecorder) ClientDataTransferUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientDataTransferUpdates", reflect.TypeOf((*MockFullNode)(nil).ClientDataTransferUpdates), arg0)
}

// ClientDealPieceCID mocks base method
func (m *MockFullNode) ClientDealPieceCID(arg0 context.Context, arg1 cid.Cid) (api.DataCIDSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientDealPieceCID", arg0, arg1)
	ret0, _ := ret[0].(api.DataCIDSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientDealPieceCID indicates an expected call of ClientDealPieceCID
func (mr *MockFullNodeMockRecorder) ClientDealPieceCID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientDealPieceCID", reflect.TypeOf((*MockFullNode)(nil).ClientDealPieceCID), arg0, arg1)
}

// ClientDealSize mocks base method
func (m *MockFullNode) ClientDealSize(arg0 context.Context, arg1 cid.Cid) (api.DataSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientDealSize", arg0, arg1)
	ret0, _ := ret[0].(api.DataSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientDealSize indicates an expected call of ClientDealSize
func (mr *MockFullNodeMockRecorder) ClientDealSize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientDealSize", reflect.TypeOf((*MockFullNode)(nil).ClientDealSize), arg0, arg1)
}

// ClientFindData mocks base method
func (m *MockFullNode) ClientFindData(arg0 context.Context, arg1 cid.Cid, arg2 *cid.Cid) ([]api.QueryOffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientFindData", arg0, arg1, arg2)
	ret0, _ := ret[0].([]api.QueryOffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientFindData indicates an expected call of ClientFindData
func (mr *MockFullNodeMockRecorder) ClientFindData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientFindData", reflect.TypeOf((*MockFullNode)(nil).ClientFindData), arg0, arg1, arg2)
}

// ClientGenCar mocks base method
func (m *MockFullNode) ClientGenCar(arg0 context.Context, arg1 api.FileRef, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGenCar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientGenCar indicates an expected call of ClientGenCar
func (mr *MockFullNodeMockRecorder) ClientGenCar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGenCar", reflect.TypeOf((*MockFullNode)(nil).ClientGenCar), arg0, arg1, arg2)
}

// ClientGetDealInfo mocks base method
func (m *MockFullNode) ClientGetDealInfo(arg0 context.Context, arg1 cid.Cid) (*api.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetDealInfo", arg0, arg1)
	ret0, _ := ret[0].(*api.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientGetDealInfo indicates an expected call of ClientGetDealInfo
func (mr *MockFullNodeMockRecorder) ClientGetDealInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetDealInfo", reflect.TypeOf((*MockFullNode)(nil).ClientGetDealInfo), arg0, arg1)
}

// ClientGetDealStatus mocks base method
func (m *MockFullNode) ClientGetDealStatus(arg0 context.Context, arg1 uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetDealStatus", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientGetDealStatus indicates an expected call of ClientGetDealStatus
func (mr *MockFullNodeMockRecorder) ClientGetDealStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetDealStatus", reflect.TypeOf((*MockFullNode)(nil).ClientGetDealStatus), arg0, arg1)
}

// ClientGetDealUpdates mocks base method
func (m *MockFullNode) ClientGetDealUpdates(arg0 context.Context) (<-chan api.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetDealUpdates", arg0)
	ret0, _ := ret[0].(<-chan api.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientGetDealUpdates indicates an expected call of ClientGetDealUpdates
func (mr *MockFullNodeMockRecorder) ClientGetDealUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetDealUpdates", reflect.TypeOf((*MockFullNode)(nil).ClientGetDealUpdates), arg0)
}

// ClientHasLocal mocks base method
func (m *MockFullNode) ClientHasLocal(arg0 context.Context, arg1 cid.Cid) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientHasLocal", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientHasLocal indicates an expected call of ClientHasLocal
func (mr *MockFullNodeMockRecorder) ClientHasLocal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientHasLocal", reflect.TypeOf((*MockFullNode)(nil).ClientHasLocal), arg0, arg1)
}

// ClientImport mocks base method
func (m *MockFullNode) ClientImport(arg0 context.Context, arg1 api.FileRef) (*api.ImportRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientImport", arg0, arg1)
	ret0, _ := ret[0].(*api.ImportRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientImport indicates an expected call of ClientImport
func (mr *MockFullNodeMockRecorder) ClientImport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientImport", reflect.TypeOf((*MockFullNode)(nil).ClientImport), arg0, arg1)
}

// ClientListDataTransfers mocks base method
func (m *MockFullNode) ClientListDataTransfers(arg0 context.Context) ([]api.DataTransferChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientListDataTransfers", arg0)
	ret0, _ := ret[0].([]api.DataTransferChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientListDataTransfers indicates an expected call of ClientListDataTransfers
func (mr *MockFullNodeMockRecorder) ClientListDataTransfers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientListDataTransfers", reflect.TypeOf((*MockFullNode)(nil).ClientListDataTransfers), arg0)
}

// ClientListDeals mocks base method
func (m *MockFullNode) ClientListDeals(arg0 context.Context) ([]api.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientListDeals", arg0)
	ret0, _ := ret[0].([]api.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientListDeals indicates an expected call of ClientListDeals
func (mr *MockFullNodeMockRecorder) ClientListDeals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientListDeals", reflect.TypeOf((*MockFullNode)(nil).ClientListDeals), arg0)
}

// ClientListImports mocks base method
func (m *MockFullNode) ClientListImports(arg0 context.Context) ([]api.Import, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientListImports", arg0)
	ret0, _ := ret[0].([]api.Import)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientListImports indicates an expected call of ClientListImports
func (mr *MockFullNodeMockRecorder) ClientListImports(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientListImports", reflect.TypeOf((*MockFullNode)(nil).ClientListImports), arg0)
}

// ClientMinerQueryOffer mocks base method
func (m *MockFullNode) ClientMinerQueryOffer(arg0 context.Context, arg1 address.Address, arg2 cid.Cid, arg3 *cid.Cid) (api.QueryOffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientMinerQueryOffer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(api.QueryOffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientMinerQueryOffer indicates an expected call of ClientMinerQueryOffer
func (mr *MockFullNodeMockRecorder) ClientMinerQueryOffer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientMinerQueryOffer", reflect.TypeOf((*MockFullNode)(nil).ClientMinerQueryOffer), arg0, arg1, arg2, arg3)
}

// ClientQueryAsk mocks base method
func (m *MockFullNode) ClientQueryAsk(arg0 context.Context, arg1 peer.ID, arg2 address.Address) (*storagemarket.StorageAsk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientQueryAsk", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storagemarket.StorageAsk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientQueryAsk indicates an expected call of ClientQueryAsk
func (mr *MockFullNodeMockRecorder) ClientQueryAsk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientQueryAsk", reflect.TypeOf((*MockFullNode)(nil).ClientQueryAsk), arg0, arg1, arg2)
}

// ClientRemoveImport mocks base method
func (m *MockFullNode) ClientRemoveImport(arg0 context.Context, arg1 multistore.StoreID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRemoveImport", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRemoveImport indicates an expected call of ClientRemoveImport
func (mr *MockFullNodeMockRecorder) ClientRemoveImport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRemoveImport", reflect.TypeOf((*MockFullNode)(nil).ClientRemoveImport), arg0, arg1)
}

// ClientRestartDataTransfer mocks base method
func (m *MockFullNode) ClientRestartDataTransfer(arg0 context.Context, arg1 datatransfer.TransferID, arg2 peer.ID, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRestartDataTransfer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRestartDataTransfer indicates an expected call of ClientRestartDataTransfer
func (mr *MockFullNodeMockRecorder) ClientRestartDataTransfer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRestartDataTransfer", reflect.TypeOf((*MockFullNode)(nil).ClientRestartDataTransfer), arg0, arg1, arg2, arg3)
}

// ClientRetrieve mocks base method
func (m *MockFullNode) ClientRetrieve(arg0 context.Context, arg1 api.RetrievalOrder, arg2 *api.FileRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRetrieve", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRetrieve indicates an expected call of ClientRetrieve
func (mr *MockFullNodeMockRecorder) ClientRetrieve(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRetrieve", reflect.TypeOf((*MockFullNode)(nil).ClientRetrieve), arg0, arg1, arg2)
}

// ClientRetrieveTryRestartInsufficientFunds mocks base method
func (m *MockFullNode) ClientRetrieveTryRestartInsufficientFunds(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRetrieveTryRestartInsufficientFunds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRetrieveTryRestartInsufficientFunds indicates an expected call of ClientRetrieveTryRestartInsufficientFunds
func (mr *MockFullNodeMockRecorder) ClientRetrieveTryRestartInsufficientFunds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRetrieveTryRestartInsufficientFunds", reflect.TypeOf((*MockFullNode)(nil).ClientRetrieveTryRestartInsufficientFunds), arg0, arg1)
}

// ClientRetrieveWithEvents mocks base method
func (m *MockFullNode) ClientRetrieveWithEvents(arg0 context.Context, arg1 api.RetrievalOrder, arg2 *api.FileRef) (<-chan marketevents.RetrievalEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRetrieveWithEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan marketevents.RetrievalEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientRetrieveWithEvents indicates an expected call of ClientRetrieveWithEvents
func (mr *MockFullNodeMockRecorder) ClientRetrieveWithEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRetrieveWithEvents", reflect.TypeOf((*MockFullNode)(nil).ClientRetrieveWithEvents), arg0, arg1, arg2)
}

// ClientStartDeal mocks base method
func (m *MockFullNode) ClientStartDeal(arg0 context.Context, arg1 *api.StartDealParams) (*cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientStartDeal", arg0, arg1)
	ret0, _ := ret[0].(*cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientStartDeal indicates an expected call of ClientStartDeal
func (mr *MockFullNodeMockRecorder) ClientStartDeal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientStartDeal", reflect.TypeOf((*MockFullNode)(nil).ClientStartDeal), arg0, arg1)
}

// Closing mocks base method
func (m *MockFullNode) Closing(arg0 context.Context) (<-chan struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Closing", arg0)
	ret0, _ := ret[0].(<-chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Closing indicates an expected call of Closing
func (mr *MockFullNodeMockRecorder) Closing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Closing", reflect.TypeOf((*MockFullNode)(nil).Closing), arg0)
}

// CreateBackup mocks base method
func (m *MockFullNode) CreateBackup(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBackup indicates an expected call of CreateBackup
func (mr *MockFullNodeMockRecorder) CreateBackup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackup", reflect.TypeOf((*MockFullNode)(nil).CreateBackup), arg0, arg1)
}

// Discover mocks base method
func (m *MockFullNode) Discover(arg0 context.Context) (apitypes.OpenRPCDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discover", arg0)
	ret0, _ := ret[0].(apitypes.OpenRPCDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Discover indicates an expected call of Discover
func (mr *MockFullNodeMockRecorder) Discover(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discover", reflect.TypeOf((*MockFullNode)(nil).Discover), arg0)
}

// GasEstimateFeeCap mocks base method
func (m *MockFullNode) GasEstimateFeeCap(arg0 context.Context, arg1 *types.Message, arg2 int64, arg3 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasEstimateFeeCap", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasEstimateFeeCap indicates an expected call of GasEstimateFeeCap
func (mr *MockFullNodeMockRecorder) GasEstimateFeeCap(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasEstimateFeeCap", reflect.TypeOf((*MockFullNode)(nil).GasEstimateFeeCap), arg0, arg1, arg2, arg3)
}

// GasEstimateGasLimit mocks base method
func (m *MockFullNode) GasEstimateGasLimit(arg0 context.Context, arg1 *types.Message, arg2 types.TipSetKey) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasEstimateGasLimit", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasEstimateGasLimit indicates an expected call of GasEstimateGasLimit
func (mr *MockFullNodeMockRecorder) GasEstimateGasLimit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasEstimateGasLimit", reflect.TypeOf((*MockFullNode)(nil).GasEstimateGasLimit), arg0, arg1, arg2)
}

// GasEstimateGasPremium mocks base method
func (m *MockFullNode) GasEstimateGasPremium(arg0 context.Context, arg1 uint64, arg2 address.Address, arg3 int64, arg4 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasEstimateGasPremium", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasEstimateGasPremium indicates an expected call of GasEstimateGasPremium
func (mr *MockFullNodeMockRecorder) GasEstimateGasPremium(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasEstimateGasPremium", reflect.TypeOf((*MockFullNode)(nil).GasEstimateGasPremium), arg0, arg1, arg2, arg3, arg4)
}

// GasEstimateMessageGas mocks base method
func (m *MockFullNode) GasEstimateMessageGas(arg0 context.Context, arg1 *types.Message, arg2 *api.MessageSendSpec, arg3 types.TipSetKey) (*types.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasEstimateMessageGas", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasEstimateMessageGas indicates an expected call of GasEstimateMessageGas
func (mr *MockFullNodeMockRecorder) GasEstimateMessageGas(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasEstimateMessageGas", reflect.TypeOf((*MockFullNode)(nil).GasEstimateMessageGas), arg0, arg1, arg2, arg3)
}

// ID mocks base method
func (m *MockFullNode) ID(arg0 context.Context) (peer.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID", arg0)
	ret0, _ := ret[0].(peer.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID
func (mr *MockFullNodeMockRecorder) ID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockFullNode)(nil).ID), arg0)
}

// LogList mocks base method
func (m *MockFullNode) LogList(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogList", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogList indicates an expected call of LogList
func (mr *MockFullNodeMockRecorder) LogList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogList", reflect.TypeOf((*MockFullNode)(nil).LogList), arg0)
}

// LogSetLevel mocks base method
func (m *MockFullNode) LogSetLevel(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogSetLevel", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogSetLevel indicates an expected call of LogSetLevel
func (mr *MockFullNodeMockRecorder) LogSetLevel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogSetLevel", reflect.TypeOf((*MockFullNode)(nil).LogSetLevel), arg0, arg1, arg2)
}

// MarketAddBalance mocks base method
func (m *MockFullNode) MarketAddBalance(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketAddBalance", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketAddBalance indicates an expected call of MarketAddBalance
func (mr *MockFullNodeMockRecorder) MarketAddBalance(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketAddBalance", reflect.TypeOf((*MockFullNode)(nil).MarketAddBalance), arg0, arg1, arg2, arg3)
}

// MarketGetReserved mocks base method
func (m *MockFullNode) MarketGetReserved(arg0 context.Context, arg1 address.Address) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketGetReserved", arg0, arg1)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketGetReserved indicates an expected call of MarketGetReserved
func (mr *MockFullNodeMockRecorder) MarketGetReserved(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketGetReserved", reflect.TypeOf((*MockFullNode)(nil).MarketGetReserved), arg0, arg1)
}

// MarketReleaseFunds mocks base method
func (m *MockFullNode) MarketReleaseFunds(arg0 context.Context, arg1 address.Address, arg2 big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketReleaseFunds", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarketReleaseFunds indicates an expected call of MarketReleaseFunds
func (mr *MockFullNodeMockRecorder) MarketReleaseFunds(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketReleaseFunds", reflect.TypeOf((*MockFullNode)(nil).MarketReleaseFunds), arg0, arg1, arg2)
}

// MarketReserveFunds mocks base method
func (m *MockFullNode) MarketReserveFunds(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketReserveFunds", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketReserveFunds indicates an expected call of MarketReserveFunds
func (mr *MockFullNodeMockRecorder) MarketReserveFunds(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketReserveFunds", reflect.TypeOf((*MockFullNode)(nil).MarketReserveFunds), arg0, arg1, arg2, arg3)
}

// MarketWithdraw mocks base method
func (m *MockFullNode) MarketWithdraw(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketWithdraw", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketWithdraw indicates an expected call of MarketWithdraw
func (mr *MockFullNodeMockRecorder) MarketWithdraw(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketWithdraw", reflect.TypeOf((*MockFullNode)(nil).MarketWithdraw), arg0, arg1, arg2, arg3)
}

// MinerCreateBlock mocks base method
func (m *MockFullNode) MinerCreateBlock(arg0 context.Context, arg1 *api.BlockTemplate) (*types.BlockMsg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinerCreateBlock", arg0, arg1)
	ret0, _ := ret[0].(*types.BlockMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MinerCreateBlock indicates an expected call of MinerCreateBlock
func (mr *MockFullNodeMockRecorder) MinerCreateBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinerCreateBlock", reflect.TypeOf((*MockFullNode)(nil).MinerCreateBlock), arg0, arg1)
}

// MinerGetBaseInfo mocks base method
func (m *MockFullNode) MinerGetBaseInfo(arg0 context.Context, arg1 address.Address, arg2 abi.ChainEpoch, arg3 types.TipSetKey) (*api.MiningBaseInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinerGetBaseInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.MiningBaseInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MinerGetBaseInfo indicates an expected call of MinerGetBaseInfo
func (mr *MockFullNodeMockRecorder) MinerGetBaseInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinerGetBaseInfo", reflect.TypeOf((*MockFullNode)(nil).MinerGetBaseInfo), arg0, arg1, arg2, arg3)
}

// MpoolBatchPush mocks base method
func (m *MockFullNode) MpoolBatchPush(arg0 context.Context, arg1 []*types.SignedMessage) ([]cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolBatchPush", arg0, arg1)
	ret0, _ := ret[0].([]cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolBatchPush indicates an expected call of MpoolBatchPush
func (mr *MockFullNodeMockRecorder) MpoolBatchPush(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolBatchPush", reflect.TypeOf((*MockFullNode)(nil).MpoolBatchPush), arg0, arg1)
}

// MpoolBatchPushMessage mocks base method
func (m *MockFullNode) MpoolBatchPushMessage(arg0 context.Context, arg1 []*types.Message, arg2 *api.MessageSendSpec) ([]*types.SignedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolBatchPushMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolBatchPushMessage indicates an expected call of MpoolBatchPushMessage
func (mr *MockFullNodeMockRecorder) MpoolBatchPushMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolBatchPushMessage", reflect.TypeOf((*MockFullNode)(nil).MpoolBatchPushMessage), arg0, arg1, arg2)
}

// MpoolBatchPushUntrusted mocks base method
func (m *MockFullNode) MpoolBatchPushUntrusted(arg0 context.Context, arg1 []*types.SignedMessage) ([]cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolBatchPushUntrusted", arg0, arg1)
	ret0, _ := ret[0].([]cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolBatchPushUntrusted indicates an expected call of MpoolBatchPushUntrusted
func (mr *MockFullNodeMockRecorder) MpoolBatchPushUntrusted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolBatchPushUntrusted", reflect.TypeOf((*MockFullNode)(nil).MpoolBatchPushUntrusted), arg0, arg1)
}

// MpoolClear mocks base method
func (m *MockFullNode) MpoolClear(arg0 context.Context, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolClear", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MpoolClear indicates an expected call of MpoolClear
func (mr *MockFullNodeMockRecorder) MpoolClear(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolClear", reflect.TypeOf((*MockFullNode)(nil).MpoolClear), arg0, arg1)
}

// MpoolGetConfig mocks base method
func (m *MockFullNode) MpoolGetConfig(arg0 context.Context) (*types.MpoolConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolGetConfig", arg0)
	ret0, _ := ret[0].(*types.MpoolConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolGetConfig indicates an expected call of MpoolGetConfig
func (mr *MockFullNodeMockRecorder) MpoolGetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolGetConfig", reflect.TypeOf((*MockFullNode)(nil).MpoolGetConfig), arg0)
}

// MpoolGetNonce mocks base method
func (m *MockFullNode) MpoolGetNonce(arg0 context.Context, arg1 address.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolGetNonce", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolGetNonce indicates an expected call of MpoolGetNonce
func (mr *MockFullNodeMockRecorder) MpoolGetNonce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolGetNonce", reflect.TypeOf((*MockFullNode)(nil).MpoolGetNonce), arg0, arg1)
}

// MpoolPending mocks base method
func (m *MockFullNode) MpoolPending(arg0 context.Context, arg1 types.TipSetKey) ([]*types.SignedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolPending", arg0, arg1)
	ret0, _ := ret[0].([]*types.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolPending indicates an expected call of MpoolPending
func (mr *MockFullNodeMockRecorder) MpoolPending(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolPending", reflect.TypeOf((*MockFullNode)(nil).MpoolPending), arg0, arg1)
}

// MpoolPush mocks base method
func (m *MockFullNode) MpoolPush(arg0 context.Context, arg1 *types.SignedMessage) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolPush", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolPush indicates an expected call of MpoolPush
func (mr *MockFullNodeMockRecorder) MpoolPush(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolPush", reflect.TypeOf((*MockFullNode)(nil).MpoolPush), arg0, arg1)
}

// MpoolPushMessage mocks base method
func (m *MockFullNode) MpoolPushMessage(arg0 context.Context, arg1 *types.Message, arg2 *api.MessageSendSpec) (*types.SignedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolPushMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolPushMessage indicates an expected call of MpoolPushMessage
func (mr *MockFullNodeMockRecorder) MpoolPushMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolPushMessage", reflect.TypeOf((*MockFullNode)(nil).MpoolPushMessage), arg0, arg1, arg2)
}

// MpoolPushUntrusted mocks base method
func (m *MockFullNode) MpoolPushUntrusted(arg0 context.Context, arg1 *types.SignedMessage) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolPushUntrusted", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolPushUntrusted indicates an expected call of MpoolPushUntrusted
func (mr *MockFullNodeMockRecorder) MpoolPushUntrusted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolPushUntrusted", reflect.TypeOf((*MockFullNode)(nil).MpoolPushUntrusted), arg0, arg1)
}

// MpoolSelect mocks base method
func (m *MockFullNode) MpoolSelect(arg0 context.Context, arg1 types.TipSetKey, arg2 float64) ([]*types.SignedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolSelect", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*types.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolSelect indicates an expected call of MpoolSelect
func (mr *MockFullNodeMockRecorder) MpoolSelect(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolSelect", reflect.TypeOf((*MockFullNode)(nil).MpoolSelect), arg0, arg1, arg2)
}

// MpoolSetConfig mocks base method
func (m *MockFullNode) MpoolSetConfig(arg0 context.Context, arg1 *types.MpoolConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolSetConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MpoolSetConfig indicates an expected call of MpoolSetConfig
func (mr *MockFullNodeMockRecorder) MpoolSetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolSetConfig", reflect.TypeOf((*MockFullNode)(nil).MpoolSetConfig), arg0, arg1)
}

// MpoolSub mocks base method
func (m *MockFullNode) MpoolSub(arg0 context.Context) (<-chan api.MpoolUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MpoolSub", arg0)
	ret0, _ := ret[0].(<-chan api.MpoolUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MpoolSub indicates an expected call of MpoolSub
func (mr *MockFullNodeMockRecorder) MpoolSub(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MpoolSub", reflect.TypeOf((*MockFullNode)(nil).MpoolSub), arg0)
}

// MsigAddApprove mocks base method
func (m *MockFullNode) MsigAddApprove(arg0 context.Context, arg1, arg2 address.Address, arg3 uint64, arg4, arg5 address.Address, arg6 bool) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigAddApprove", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigAddApprove indicates an expected call of MsigAddApprove
func (mr *MockFullNodeMockRecorder) MsigAddApprove(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigAddApprove", reflect.TypeOf((*MockFullNode)(nil).MsigAddApprove), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MsigAddCancel mocks base method
func (m *MockFullNode) MsigAddCancel(arg0 context.Context, arg1, arg2 address.Address, arg3 uint64, arg4 address.Address, arg5 bool) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigAddCancel", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigAddCancel indicates an expected call of MsigAddCancel
func (mr *MockFullNodeMockRecorder) MsigAddCancel(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigAddCancel", reflect.TypeOf((*MockFullNode)(nil).MsigAddCancel), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MsigAddPropose mocks base method
func (m *MockFullNode) MsigAddPropose(arg0 context.Context, arg1, arg2, arg3 address.Address, arg4 bool) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigAddPropose", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigAddPropose indicates an expected call of MsigAddPropose
func (mr *MockFullNodeMockRecorder) MsigAddPropose(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigAddPropose", reflect.TypeOf((*MockFullNode)(nil).MsigAddPropose), arg0, arg1, arg2, arg3, arg4)
}

// MsigApprove mocks base method
func (m *MockFullNode) MsigApprove(arg0 context.Context, arg1 address.Address, arg2 uint64, arg3 address.Address) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigApprove", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigApprove indicates an expected call of MsigApprove
func (mr *MockFullNodeMockRecorder) MsigApprove(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigApprove", reflect.TypeOf((*MockFullNode)(nil).MsigApprove), arg0, arg1, arg2, arg3)
}

// MsigApproveTxnHash mocks base method
func (m *MockFullNode) MsigApproveTxnHash(arg0 context.Context, arg1 address.Address, arg2 uint64, arg3, arg4 address.Address, arg5 big.Int, arg6 address.Address, arg7 uint64, arg8 []byte) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigApproveTxnHash", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigApproveTxnHash indicates an expected call of MsigApproveTxnHash
func (mr *MockFullNodeMockRecorder) MsigApproveTxnHash(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigApproveTxnHash", reflect.TypeOf((*MockFullNode)(nil).MsigApproveTxnHash), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// MsigCancel mocks base method
func (m *MockFullNode) MsigCancel(arg0 context.Context, arg1 address.Address, arg2 uint64, arg3 address.Address, arg4 big.Int, arg5 address.Address, arg6 uint64, arg7 []byte) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigCancel", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigCancel indicates an expected call of MsigCancel
func (mr *MockFullNodeMockRecorder) MsigCancel(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigCancel", reflect.TypeOf((*MockFullNode)(nil).MsigCancel), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// MsigCreate mocks base method
func (m *MockFullNode) MsigCreate(arg0 context.Context, arg1 uint64, arg2 []address.Address, arg3 abi.ChainEpoch, arg4 big.Int, arg5 address.Address, arg6 big.Int) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigCreate", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigCreate indicates an expected call of MsigCreate
func (mr *MockFullNodeMockRecorder) MsigCreate(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigCreate", reflect.TypeOf((*MockFullNode)(nil).MsigCreate), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MsigGetAvailableBalance mocks base method
func (m *MockFullNode) MsigGetAvailableBalance(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigGetAvailableBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigGetAvailableBalance indicates an expected call of MsigGetAvailableBalance
func (mr *MockFullNodeMockRecorder) MsigGetAvailableBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigGetAvailableBalance", reflect.TypeOf((*MockFullNode)(nil).MsigGetAvailableBalance), arg0, arg1, arg2)
}

// MsigGetPending mocks base method
func (m *MockFullNode) MsigGetPending(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) ([]*api.MsigTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigGetPending", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*api.MsigTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigGetPending indicates an expected call of MsigGetPending
func (mr *MockFullNodeMockRecorder) MsigGetPending(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigGetPending", reflect.TypeOf((*MockFullNode)(nil).MsigGetPending), arg0, arg1, arg2)
}

// MsigGetVested mocks base method
func (m *MockFullNode) MsigGetVested(arg0 context.Context, arg1 address.Address, arg2, arg3 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigGetVested", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigGetVested indicates an expected call of MsigGetVested
func (mr *MockFullNodeMockRecorder) MsigGetVested(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigGetVested", reflect.TypeOf((*MockFullNode)(nil).MsigGetVested), arg0, arg1, arg2, arg3)
}

// MsigGetVestingSchedule mocks base method
func (m *MockFullNode) MsigGetVestingSchedule(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (api.MsigVesting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigGetVestingSchedule", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.MsigVesting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigGetVestingSchedule indicates an expected call of MsigGetVestingSchedule
func (mr *MockFullNodeMockRecorder) MsigGetVestingSchedule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigGetVestingSchedule", reflect.TypeOf((*MockFullNode)(nil).MsigGetVestingSchedule), arg0, arg1, arg2)
}

// MsigPropose mocks base method
func (m *MockFullNode) MsigPropose(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int, arg4 address.Address, arg5 uint64, arg6 []byte) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigPropose", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigPropose indicates an expected call of MsigPropose
func (mr *MockFullNodeMockRecorder) MsigPropose(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigPropose", reflect.TypeOf((*MockFullNode)(nil).MsigPropose), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MsigRemoveSigner mocks base method
func (m *MockFullNode) MsigRemoveSigner(arg0 context.Context, arg1, arg2, arg3 address.Address, arg4 bool) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigRemoveSigner", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigRemoveSigner indicates an expected call of MsigRemoveSigner
func (mr *MockFullNodeMockRecorder) MsigRemoveSigner(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigRemoveSigner", reflect.TypeOf((*MockFullNode)(nil).MsigRemoveSigner), arg0, arg1, arg2, arg3, arg4)
}

// MsigSwapApprove mocks base method
func (m *MockFullNode) MsigSwapApprove(arg0 context.Context, arg1, arg2 address.Address, arg3 uint64, arg4, arg5, arg6 address.Address) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigSwapApprove", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigSwapApprove indicates an expected call of MsigSwapApprove
func (mr *MockFullNodeMockRecorder) MsigSwapApprove(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigSwapApprove", reflect.TypeOf((*MockFullNode)(nil).MsigSwapApprove), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MsigSwapCancel mocks base method
func (m *MockFullNode) MsigSwapCancel(arg0 context.Context, arg1, arg2 address.Address, arg3 uint64, arg4, arg5 address.Address) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigSwapCancel", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigSwapCancel indicates an expected call of MsigSwapCancel
func (mr *MockFullNodeMockRecorder) MsigSwapCancel(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigSwapCancel", reflect.TypeOf((*MockFullNode)(nil).MsigSwapCancel), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MsigSwapPropose mocks base method
func (m *MockFullNode) MsigSwapPropose(arg0 context.Context, arg1, arg2, arg3, arg4 address.Address) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MsigSwapPropose", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MsigSwapPropose indicates an expected call of MsigSwapPropose
func (mr *MockFullNodeMockRecorder) MsigSwapPropose(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsigSwapPropose", reflect.TypeOf((*MockFullNode)(nil).MsigSwapPropose), arg0, arg1, arg2, arg3, arg4)
}

// NetAddrsListen mocks base method
func (m *MockFullNode) NetAddrsListen(arg0 context.Context) (peer.AddrInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetAddrsListen", arg0)
	ret0, _ := ret[0].(peer.AddrInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetAddrsListen indicates an expected call of NetAddrsListen
func (mr *MockFullNodeMockRecorder) NetAddrsListen(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetAddrsListen", reflect.TypeOf((*MockFullNode)(nil).NetAddrsListen), arg0)
}

// NetAgentVersion mocks base method
func (m *MockFullNode) NetAgentVersion(arg0 context.Context, arg1 peer.ID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetAgentVersion", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetAgentVersion indicates an expected call of NetAgentVersion
func (mr *MockFullNodeMockRecorder) NetAgentVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetAgentVersion", reflect.TypeOf((*MockFullNode)(nil).NetAgentVersion), arg0, arg1)
}

// NetAutoNatStatus mocks base method
func (m *MockFullNode) NetAutoNatStatus(arg0 context.Context) (api.NatInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetAutoNatStatus", arg0)
	ret0, _ := ret[0].(api.NatInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetAutoNatStatus indicates an expected call of NetAutoNatStatus
func (mr *MockFullNodeMockRecorder) NetAutoNatStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetAutoNatStatus", reflect.TypeOf((*MockFullNode)(nil).NetAutoNatStatus), arg0)
}

// NetBandwidthStats mocks base method
func (m *MockFullNode) NetBandwidthStats(arg0 context.Context) (metrics.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetBandwidthStats", arg0)
	ret0, _ := ret[0].(metrics.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetBandwidthStats indicates an expected call of NetBandwidthStats
func (mr *MockFullNodeMockRecorder) NetBandwidthStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetBandwidthStats", reflect.TypeOf((*MockFullNode)(nil).NetBandwidthStats), arg0)
}

// NetBandwidthStatsByPeer mocks base method
func (m *MockFullNode) NetBandwidthStatsByPeer(arg0 context.Context) (map[string]metrics.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetBandwidthStatsByPeer", arg0)
	ret0, _ := ret[0].(map[string]metrics.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetBandwidthStatsByPeer indicates an expected call of NetBandwidthStatsByPeer
func (mr *MockFullNodeMockRecorder) NetBandwidthStatsByPeer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetBandwidthStatsByPeer", reflect.TypeOf((*MockFullNode)(nil).NetBandwidthStatsByPeer), arg0)
}

// NetBandwidthStatsByProtocol mocks base method
func (m *MockFullNode) NetBandwidthStatsByProtocol(arg0 context.Context) (map[protocol.ID]metrics.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetBandwidthStatsByProtocol", arg0)
	ret0, _ := ret[0].(map[protocol.ID]metrics.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetBandwidthStatsByProtocol indicates an expected call of NetBandwidthStatsByProtocol
func (mr *MockFullNodeMockRecorder) NetBandwidthStatsByProtocol(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetBandwidthStatsByProtocol", reflect.TypeOf((*MockFullNode)(nil).NetBandwidthStatsByProtocol), arg0)
}

// NetBlockAdd mocks base method
func (m *MockFullNode) NetBlockAdd(arg0 context.Context, arg1 api.NetBlockList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetBlockAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetBlockAdd indicates an expected call of NetBlockAdd
func (mr *MockFullNodeMockRecorder) NetBlockAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetBlockAdd", reflect.TypeOf((*MockFullNode)(nil).NetBlockAdd), arg0, arg1)
}

// NetBlockList mocks base method
func (m *MockFullNode) NetBlockList(arg0 context.Context) (api.NetBlockList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetBlockList", arg0)
	ret0, _ := ret[0].(api.NetBlockList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetBlockList indicates an expected call of NetBlockList
func (mr *MockFullNodeMockRecorder) NetBlockList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetBlockList", reflect.TypeOf((*MockFullNode)(nil).NetBlockList), arg0)
}

// NetBlockRemove mocks base method
func (m *MockFullNode) NetBlockRemove(arg0 context.Context, arg1 api.NetBlockList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetBlockRemove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetBlockRemove indicates an expected call of NetBlockRemove
func (mr *MockFullNodeMockRecorder) NetBlockRemove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetBlockRemove", reflect.TypeOf((*MockFullNode)(nil).NetBlockRemove), arg0, arg1)
}

// NetConnect mocks base method
func (m *MockFullNode) NetConnect(arg0 context.Context, arg1 peer.AddrInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetConnect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetConnect indicates an expected call of NetConnect
func (mr *MockFullNodeMockRecorder) NetConnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetConnect", reflect.TypeOf((*MockFullNode)(nil).NetConnect), arg0, arg1)
}

// NetConnectedness mocks base method
func (m *MockFullNode) NetConnectedness(arg0 context.Context, arg1 peer.ID) (network0.Connectedness, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetConnectedness", arg0, arg1)
	ret0, _ := ret[0].(network0.Connectedness)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetConnectedness indicates an expected call of NetConnectedness
func (mr *MockFullNodeMockRecorder) NetConnectedness(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetConnectedness", reflect.TypeOf((*MockFullNode)(nil).NetConnectedness), arg0, arg1)
}

// NetDisconnect mocks base method
func (m *MockFullNode) NetDisconnect(arg0 context.Context, arg1 peer.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetDisconnect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetDisconnect indicates an expected call of NetDisconnect
func (mr *MockFullNodeMockRecorder) NetDisconnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetDisconnect", reflect.TypeOf((*MockFullNode)(nil).NetDisconnect), arg0, arg1)
}

// NetFindPeer mocks base method
func (m *MockFullNode) NetFindPeer(arg0 context.Context, arg1 peer.ID) (peer.AddrInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetFindPeer", arg0, arg1)
	ret0, _ := ret[0].(peer.AddrInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetFindPeer indicates an expected call of NetFindPeer
func (mr *MockFullNodeMockRecorder) NetFindPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetFindPeer", reflect.TypeOf((*MockFullNode)(nil).NetFindPeer), arg0, arg1)
}

// NetPeerInfo mocks base method
func (m *MockFullNode) NetPeerInfo(arg0 context.Context, arg1 peer.ID) (*api.ExtendedPeerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetPeerInfo", arg0, arg1)
	ret0, _ := ret[0].(*api.ExtendedPeerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetPeerInfo indicates an expected call of NetPeerInfo
func (mr *MockFullNodeMockRecorder) NetPeerInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetPeerInfo", reflect.TypeOf((*MockFullNode)(nil).NetPeerInfo), arg0, arg1)
}

// NetPeers mocks base method
func (m *MockFullNode) NetPeers(arg0 context.Context) ([]peer.AddrInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetPeers", arg0)
	ret0, _ := ret[0].([]peer.AddrInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetPeers indicates an expected call of NetPeers
func (mr *MockFullNodeMockRecorder) NetPeers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetPeers", reflect.TypeOf((*MockFullNode)(nil).NetPeers), arg0)
}

// NetPubsubScores mocks base method
func (m *MockFullNode) NetPubsubScores(arg0 context.Context) ([]api.PubsubScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetPubsubScores", arg0)
	ret0, _ := ret[0].([]api.PubsubScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetPubsubScores indicates an expected call of NetPubsubScores
func (mr *MockFullNodeMockRecorder) NetPubsubScores(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetPubsubScores", reflect.TypeOf((*MockFullNode)(nil).NetPubsubScores), arg0)
}

// PaychAllocateLane mocks base method
func (m *MockFullNode) PaychAllocateLane(arg0 context.Context, arg1 address.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychAllocateLane", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychAllocateLane indicates an expected call of PaychAllocateLane
func (mr *MockFullNodeMockRecorder) PaychAllocateLane(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychAllocateLane", reflect.TypeOf((*MockFullNode)(nil).PaychAllocateLane), arg0, arg1)
}

// PaychAvailableFunds mocks base method
func (m *MockFullNode) PaychAvailableFunds(arg0 context.Context, arg1 address.Address) (*api.ChannelAvailableFunds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychAvailableFunds", arg0, arg1)
	ret0, _ := ret[0].(*api.ChannelAvailableFunds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychAvailableFunds indicates an expected call of PaychAvailableFunds
func (mr *MockFullNodeMockRecorder) PaychAvailableFunds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychAvailableFunds", reflect.TypeOf((*MockFullNode)(nil).PaychAvailableFunds), arg0, arg1)
}

// PaychAvailableFundsByFromTo mocks base method
func (m *MockFullNode) PaychAvailableFundsByFromTo(arg0 context.Context, arg1, arg2 address.Address) (*api.ChannelAvailableFunds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychAvailableFundsByFromTo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.ChannelAvailableFunds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychAvailableFundsByFromTo indicates an expected call of PaychAvailableFundsByFromTo
func (mr *MockFullNodeMockRecorder) PaychAvailableFundsByFromTo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychAvailableFundsByFromTo", reflect.TypeOf((*MockFullNode)(nil).PaychAvailableFundsByFromTo), arg0, arg1, arg2)
}

// PaychCollect mocks base method
func (m *MockFullNode) PaychCollect(arg0 context.Context, arg1 address.Address) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychCollect", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychCollect indicates an expected call of PaychCollect
func (mr *MockFullNodeMockRecorder) PaychCollect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychCollect", reflect.TypeOf((*MockFullNode)(nil).PaychCollect), arg0, arg1)
}

// PaychGet mocks base method
func (m *MockFullNode) PaychGet(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int) (*api.ChannelInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychGet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.ChannelInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychGet indicates an expected call of PaychGet
func (mr *MockFullNodeMockRecorder) PaychGet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychGet", reflect.TypeOf((*MockFullNode)(nil).PaychGet), arg0, arg1, arg2, arg3)
}

// PaychGetWaitReady mocks base method
func (m *MockFullNode) PaychGetWaitReady(arg0 context.Context, arg1 cid.Cid) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychGetWaitReady", arg0, arg1)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychGetWaitReady indicates an expected call of PaychGetWaitReady
func (mr *MockFullNodeMockRecorder) PaychGetWaitReady(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychGetWaitReady", reflect.TypeOf((*MockFullNode)(nil).PaychGetWaitReady), arg0, arg1)
}

// PaychList mocks base method
func (m *MockFullNode) PaychList(arg0 context.Context) ([]address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychList", arg0)
	ret0, _ := ret[0].([]address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychList indicates an expected call of PaychList
func (mr *MockFullNodeMockRecorder) PaychList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychList", reflect.TypeOf((*MockFullNode)(nil).PaychList), arg0)
}

// PaychNewPayment mocks base method
func (m *MockFullNode) PaychNewPayment(arg0 context.Context, arg1, arg2 address.Address, arg3 []api.VoucherSpec) (*api.PaymentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychNewPayment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.PaymentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychNewPayment indicates an expected call of PaychNewPayment
func (mr *MockFullNodeMockRecorder) PaychNewPayment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychNewPayment", reflect.TypeOf((*MockFullNode)(nil).PaychNewPayment), arg0, arg1, arg2, arg3)
}

// PaychSettle mocks base method
func (m *MockFullNode) PaychSettle(arg0 context.Context, arg1 address.Address) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychSettle", arg0, arg1)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychSettle indicates an expected call of PaychSettle
func (mr *MockFullNodeMockRecorder) PaychSettle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychSettle", reflect.TypeOf((*MockFullNode)(nil).PaychSettle), arg0, arg1)
}

// PaychStatus mocks base method
func (m *MockFullNode) PaychStatus(arg0 context.Context, arg1 address.Address) (*api.PaychStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychStatus", arg0, arg1)
	ret0, _ := ret[0].(*api.PaychStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychStatus indicates an expected call of PaychStatus
func (mr *MockFullNodeMockRecorder) PaychStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychStatus", reflect.TypeOf((*MockFullNode)(nil).PaychStatus), arg0, arg1)
}

// PaychVoucherAdd mocks base method
func (m *MockFullNode) PaychVoucherAdd(arg0 context.Context, arg1 address.Address, arg2 *paych.SignedVoucher, arg3 []byte, arg4 big.Int) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychVoucherAdd", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychVoucherAdd indicates an expected call of PaychVoucherAdd
func (mr *MockFullNodeMockRecorder) PaychVoucherAdd(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychVoucherAdd", reflect.TypeOf((*MockFullNode)(nil).PaychVoucherAdd), arg0, arg1, arg2, arg3, arg4)
}

// PaychVoucherCheckSpendable mocks base method
func (m *MockFullNode) PaychVoucherCheckSpendable(arg0 context.Context, arg1 address.Address, arg2 *paych.SignedVoucher, arg3, arg4 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychVoucherCheckSpendable", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychVoucherCheckSpendable indicates an expected call of PaychVoucherCheckSpendable
func (mr *MockFullNodeMockRecorder) PaychVoucherCheckSpendable(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychVoucherCheckSpendable", reflect.TypeOf((*MockFullNode)(nil).PaychVoucherCheckSpendable), arg0, arg1, arg2, arg3, arg4)
}

// PaychVoucherCheckValid mocks base method
func (m *MockFullNode) PaychVoucherCheckValid(arg0 context.Context, arg1 address.Address, arg2 *paych.SignedVoucher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychVoucherCheckValid", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PaychVoucherCheckValid indicates an expected call of PaychVoucherCheckValid
func (mr *MockFullNodeMockRecorder) PaychVoucherCheckValid(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychVoucherCheckValid", reflect.TypeOf((*MockFullNode)(nil).PaychVoucherCheckValid), arg0, arg1, arg2)
}

// PaychVoucherCreate mocks base method
func (m *MockFullNode) PaychVoucherCreate(arg0 context.Context, arg1 address.Address, arg2 big.Int, arg3 uint64) (*api.VoucherCreateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychVoucherCreate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.VoucherCreateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychVoucherCreate indicates an expected call of PaychVoucherCreate
func (mr *MockFullNodeMockRecorder) PaychVoucherCreate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychVoucherCreate", reflect.TypeOf((*MockFullNode)(nil).PaychVoucherCreate), arg0, arg1, arg2, arg3)
}

// PaychVoucherList mocks base method
func (m *MockFullNode) PaychVoucherList(arg0 context.Context, arg1 address.Address) ([]*paych.SignedVoucher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychVoucherList", arg0, arg1)
	ret0, _ := ret[0].([]*paych.SignedVoucher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychVoucherList indicates an expected call of PaychVoucherList
func (mr *MockFullNodeMockRecorder) PaychVoucherList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychVoucherList", reflect.TypeOf((*MockFullNode)(nil).PaychVoucherList), arg0, arg1)
}

// PaychVoucherSubmit mocks base method
func (m *MockFullNode) PaychVoucherSubmit(arg0 context.Context, arg1 address.Address, arg2 *paych.SignedVoucher, arg3, arg4 []byte) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaychVoucherSubmit", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaychVoucherSubmit indicates an expected call of PaychVoucherSubmit
func (mr *MockFullNodeMockRecorder) PaychVoucherSubmit(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaychVoucherSubmit", reflect.TypeOf((*MockFullNode)(nil).PaychVoucherSubmit), arg0, arg1, arg2, arg3, arg4)
}

// Session mocks base method
func (m *MockFullNode) Session(arg0 context.Context) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", arg0)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Session indicates an expected call of Session
func (mr *MockFullNodeMockRecorder) Session(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockFullNode)(nil).Session), arg0)
}

// Shutdown mocks base method
func (m *MockFullNode) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockFullNodeMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockFullNode)(nil).Shutdown), arg0)
}

// StateAccountKey mocks base method
func (m *MockFullNode) StateAccountKey(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAccountKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAccountKey indicates an expected call of StateAccountKey
func (mr *MockFullNodeMockRecorder) StateAccountKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAccountKey", reflect.TypeOf((*MockFullNode)(nil).StateAccountKey), arg0, arg1, arg2)
}

// StateAllMinerFaults mocks base method
func (m *MockFullNode) StateAllMinerFaults(arg0 context.Context, arg1 abi.ChainEpoch, arg2 types.TipSetKey) ([]*api.Fault, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateAllMinerFaults", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*api.Fault)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateAllMinerFaults indicates an expected call of StateAllMinerFaults
func (mr *MockFullNodeMockRecorder) StateAllMinerFaults(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateAllMinerFaults", reflect.TypeOf((*MockFullNode)(nil).StateAllMinerFaults), arg0, arg1, arg2)
}

// StateCall mocks base method
func (m *MockFullNode) StateCall(arg0 context.Context, arg1 *types.Message, arg2 types.TipSetKey) (*api.InvocResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateCall", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.InvocResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateCall indicates an expected call of StateCall
func (mr *MockFullNodeMockRecorder) StateCall(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateCall", reflect.TypeOf((*MockFullNode)(nil).StateCall), arg0, arg1, arg2)
}

// StateChangedActors mocks base method
func (m *MockFullNode) StateChangedActors(arg0 context.Context, arg1, arg2 cid.Cid) (map[string]types.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateChangedActors", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]types.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateChangedActors indicates an expected call of StateChangedActors
func (mr *MockFullNodeMockRecorder) StateChangedActors(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateChangedActors", reflect.TypeOf((*MockFullNode)(nil).StateChangedActors), arg0, arg1, arg2)
}

// StateCirculatingSupply mocks base method
func (m *MockFullNode) StateCirculatingSupply(arg0 context.Context, arg1 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateCirculatingSupply", arg0, arg1)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateCirculatingSupply indicates an expected call of StateCirculatingSupply
func (mr *MockFullNodeMockRecorder) StateCirculatingSupply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateCirculatingSupply", reflect.TypeOf((*MockFullNode)(nil).StateCirculatingSupply), arg0, arg1)
}

// StateCompute mocks base method
func (m *MockFullNode) StateCompute(arg0 context.Context, arg1 abi.ChainEpoch, arg2 []*types.Message, arg3 types.TipSetKey) (*api.ComputeStateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateCompute", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.ComputeStateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateCompute indicates an expected call of StateCompute
func (mr *MockFullNodeMockRecorder) StateCompute(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateCompute", reflect.TypeOf((*MockFullNode)(nil).StateCompute), arg0, arg1, arg2, arg3)
}

// StateDealProviderCollateralBounds mocks base method
func (m *MockFullNode) StateDealProviderCollateralBounds(arg0 context.Context, arg1 abi.PaddedPieceSize, arg2 bool, arg3 types.TipSetKey) (api.DealCollateralBounds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateDealProviderCollateralBounds", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(api.DealCollateralBounds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateDealProviderCollateralBounds indicates an expected call of StateDealProviderCollateralBounds
func (mr *MockFullNodeMockRecorder) StateDealProviderCollateralBounds(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateDealProviderCollateralBounds", reflect.TypeOf((*MockFullNode)(nil).StateDealProviderCollateralBounds), arg0, arg1, arg2, arg3)
}

// StateDecodeParams mocks base method
func (m *MockFullNode) StateDecodeParams(arg0 context.Context, arg1 address.Address, arg2 abi.MethodNum, arg3 []byte, arg4 types.TipSetKey) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateDecodeParams", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateDecodeParams indicates an expected call of StateDecodeParams
func (mr *MockFullNodeMockRecorder) StateDecodeParams(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateDecodeParams", reflect.TypeOf((*MockFullNode)(nil).StateDecodeParams), arg0, arg1, arg2, arg3, arg4)
}

// StateGetActor mocks base method
func (m *MockFullNode) StateGetActor(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (*types.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateGetActor", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateGetActor indicates an expected call of StateGetActor
func (mr *MockFullNodeMockRecorder) StateGetActor(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateGetActor", reflect.TypeOf((*MockFullNode)(nil).StateGetActor), arg0, arg1, arg2)
}

// StateGetReceipt mocks base method
func (m *MockFullNode) StateGetReceipt(arg0 context.Context, arg1 cid.Cid, arg2 types.TipSetKey) (*types.MessageReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateGetReceipt", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.MessageReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateGetReceipt indicates an expected call of StateGetReceipt
func (mr *MockFullNodeMockRecorder) StateGetReceipt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateGetReceipt", reflect.TypeOf((*MockFullNode)(nil).StateGetReceipt), arg0, arg1, arg2)
}

// StateListActors mocks base method
func (m *MockFullNode) StateListActors(arg0 context.Context, arg1 types.TipSetKey) ([]address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateListActors", arg0, arg1)
	ret0, _ := ret[0].([]address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateListActors indicates an expected call of StateListActors
func (mr *MockFullNodeMockRecorder) StateListActors(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateListActors", reflect.TypeOf((*MockFullNode)(nil).StateListActors), arg0, arg1)
}

// StateListMessages mocks base method
func (m *MockFullNode) StateListMessages(arg0 context.Context, arg1 *api.MessageMatch, arg2 types.TipSetKey, arg3 abi.ChainEpoch) ([]cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateListMessages", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateListMessages indicates an expected call of StateListMessages
func (mr *MockFullNodeMockRecorder) StateListMessages(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateListMessages", reflect.TypeOf((*MockFullNode)(nil).StateListMessages), arg0, arg1, arg2, arg3)
}

// StateListMiners mocks base method
func (m *MockFullNode) StateListMiners(arg0 context.Context, arg1 types.TipSetKey) ([]address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateListMiners", arg0, arg1)
	ret0, _ := ret[0].([]address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateListMiners indicates an expected call of StateListMiners
func (mr *MockFullNodeMockRecorder) StateListMiners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateListMiners", reflect.TypeOf((*MockFullNode)(nil).StateListMiners), arg0, arg1)
}

// StateLookupID mocks base method
func (m *MockFullNode) StateLookupID(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateLookupID", arg0, arg1, arg2)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateLookupID indicates an expected call of StateLookupID
func (mr *MockFullNodeMockRecorder) StateLookupID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateLookupID", reflect.TypeOf((*MockFullNode)(nil).StateLookupID), arg0, arg1, arg2)
}

// StateMarketBalance mocks base method
func (m *MockFullNode) StateMarketBalance(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (api.MarketBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMarketBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.MarketBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMarketBalance indicates an expected call of StateMarketBalance
func (mr *MockFullNodeMockRecorder) StateMarketBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMarketBalance", reflect.TypeOf((*MockFullNode)(nil).StateMarketBalance), arg0, arg1, arg2)
}

// StateMarketDeals mocks base method
func (m *MockFullNode) StateMarketDeals(arg0 context.Context, arg1 types.TipSetKey) (map[string]api.MarketDeal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMarketDeals", arg0, arg1)
	ret0, _ := ret[0].(map[string]api.MarketDeal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMarketDeals indicates an expected call of StateMarketDeals
func (mr *MockFullNodeMockRecorder) StateMarketDeals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMarketDeals", reflect.TypeOf((*MockFullNode)(nil).StateMarketDeals), arg0, arg1)
}

// StateMarketParticipants mocks base method
func (m *MockFullNode) StateMarketParticipants(arg0 context.Context, arg1 types.TipSetKey) (map[string]api.MarketBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMarketParticipants", arg0, arg1)
	ret0, _ := ret[0].(map[string]api.MarketBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMarketParticipants indicates an expected call of StateMarketParticipants
func (mr *MockFullNodeMockRecorder) StateMarketParticipants(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMarketParticipants", reflect.TypeOf((*MockFullNode)(nil).StateMarketParticipants), arg0, arg1)
}

// StateMarketStorageDeal mocks base method
func (m *MockFullNode) StateMarketStorageDeal(arg0 context.Context, arg1 abi.DealID, arg2 types.TipSetKey) (*api.MarketDeal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMarketStorageDeal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.MarketDeal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMarketStorageDeal indicates an expected call of StateMarketStorageDeal
func (mr *MockFullNodeMockRecorder) StateMarketStorageDeal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMarketStorageDeal", reflect.TypeOf((*MockFullNode)(nil).StateMarketStorageDeal), arg0, arg1, arg2)
}

// StateMinerActiveSectors mocks base method
func (m *MockFullNode) StateMinerActiveSectors(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerActiveSectors", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*miner.SectorOnChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerActiveSectors indicates an expected call of StateMinerActiveSectors
func (mr *MockFullNodeMockRecorder) StateMinerActiveSectors(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerActiveSectors", reflect.TypeOf((*MockFullNode)(nil).StateMinerActiveSectors), arg0, arg1, arg2)
}

// StateMinerAvailableBalance mocks base method
func (m *MockFullNode) StateMinerAvailableBalance(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerAvailableBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerAvailableBalance indicates an expected call of StateMinerAvailableBalance
func (mr *MockFullNodeMockRecorder) StateMinerAvailableBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerAvailableBalance", reflect.TypeOf((*MockFullNode)(nil).StateMinerAvailableBalance), arg0, arg1, arg2)
}

// StateMinerDeadlines mocks base method
func (m *MockFullNode) StateMinerDeadlines(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) ([]api.Deadline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerDeadlines", arg0, arg1, arg2)
	ret0, _ := ret[0].([]api.Deadline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerDeadlines indicates an expected call of StateMinerDeadlines
func (mr *MockFullNodeMockRecorder) StateMinerDeadlines(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerDeadlines", reflect.TypeOf((*MockFullNode)(nil).StateMinerDeadlines), arg0, arg1, arg2)
}

// StateMinerFaults mocks base method
func (m *MockFullNode) StateMinerFaults(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (bitfield.BitField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerFaults", arg0, arg1, arg2)
	ret0, _ := ret[0].(bitfield.BitField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerFaults indicates an expected call of StateMinerFaults
func (mr *MockFullNodeMockRecorder) StateMinerFaults(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerFaults", reflect.TypeOf((*MockFullNode)(nil).StateMinerFaults), arg0, arg1, arg2)
}

// StateMinerInfo mocks base method
func (m *MockFullNode) StateMinerInfo(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (miner.MinerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(miner.MinerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerInfo indicates an expected call of StateMinerInfo
func (mr *MockFullNodeMockRecorder) StateMinerInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerInfo", reflect.TypeOf((*MockFullNode)(nil).StateMinerInfo), arg0, arg1, arg2)
}

// StateMinerInitialPledgeCollateral mocks base method
func (m *MockFullNode) StateMinerInitialPledgeCollateral(arg0 context.Context, arg1 address.Address, arg2 miner0.SectorPreCommitInfo, arg3 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerInitialPledgeCollateral", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerInitialPledgeCollateral indicates an expected call of StateMinerInitialPledgeCollateral
func (mr *MockFullNodeMockRecorder) StateMinerInitialPledgeCollateral(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerInitialPledgeCollateral", reflect.TypeOf((*MockFullNode)(nil).StateMinerInitialPledgeCollateral), arg0, arg1, arg2, arg3)
}

// StateMinerPartitions mocks base method
func (m *MockFullNode) StateMinerPartitions(arg0 context.Context, arg1 address.Address, arg2 uint64, arg3 types.TipSetKey) ([]api.Partition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerPartitions", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]api.Partition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerPartitions indicates an expected call of StateMinerPartitions
func (mr *MockFullNodeMockRecorder) StateMinerPartitions(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerPartitions", reflect.TypeOf((*MockFullNode)(nil).StateMinerPartitions), arg0, arg1, arg2, arg3)
}

// StateMinerPower mocks base method
func (m *MockFullNode) StateMinerPower(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (*api.MinerPower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerPower", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.MinerPower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerPower indicates an expected call of StateMinerPower
func (mr *MockFullNodeMockRecorder) StateMinerPower(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerPower", reflect.TypeOf((*MockFullNode)(nil).StateMinerPower), arg0, arg1, arg2)
}

// StateMinerPreCommitDepositForPower mocks base method
func (m *MockFullNode) StateMinerPreCommitDepositForPower(arg0 context.Context, arg1 address.Address, arg2 miner0.SectorPreCommitInfo, arg3 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerPreCommitDepositForPower", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerPreCommitDepositForPower indicates an expected call of StateMinerPreCommitDepositForPower
func (mr *MockFullNodeMockRecorder) StateMinerPreCommitDepositForPower(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerPreCommitDepositForPower", reflect.TypeOf((*MockFullNode)(nil).StateMinerPreCommitDepositForPower), arg0, arg1, arg2, arg3)
}

// StateMinerProvingDeadline mocks base method
func (m *MockFullNode) StateMinerProvingDeadline(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (*dline.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerProvingDeadline", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dline.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerProvingDeadline indicates an expected call of StateMinerProvingDeadline
func (mr *MockFullNodeMockRecorder) StateMinerProvingDeadline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerProvingDeadline", reflect.TypeOf((*MockFullNode)(nil).StateMinerProvingDeadline), arg0, arg1, arg2)
}

// StateMinerRecoveries mocks base method
func (m *MockFullNode) StateMinerRecoveries(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (bitfield.BitField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerRecoveries", arg0, arg1, arg2)
	ret0, _ := ret[0].(bitfield.BitField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerRecoveries indicates an expected call of StateMinerRecoveries
func (mr *MockFullNodeMockRecorder) StateMinerRecoveries(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerRecoveries", reflect.TypeOf((*MockFullNode)(nil).StateMinerRecoveries), arg0, arg1, arg2)
}

// StateMinerSectorAllocated mocks base method
func (m *MockFullNode) StateMinerSectorAllocated(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerSectorAllocated", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerSectorAllocated indicates an expected call of StateMinerSectorAllocated
func (mr *MockFullNodeMockRecorder) StateMinerSectorAllocated(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerSectorAllocated", reflect.TypeOf((*MockFullNode)(nil).StateMinerSectorAllocated), arg0, arg1, arg2, arg3)
}

// StateMinerSectorCount mocks base method
func (m *MockFullNode) StateMinerSectorCount(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (api.MinerSectors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerSectorCount", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.MinerSectors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerSectorCount indicates an expected call of StateMinerSectorCount
func (mr *MockFullNodeMockRecorder) StateMinerSectorCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerSectorCount", reflect.TypeOf((*MockFullNode)(nil).StateMinerSectorCount), arg0, arg1, arg2)
}

// StateMinerSectors mocks base method
func (m *MockFullNode) StateMinerSectors(arg0 context.Context, arg1 address.Address, arg2 *bitfield.BitField, arg3 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerSectors", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*miner.SectorOnChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerSectors indicates an expected call of StateMinerSectors
func (mr *MockFullNodeMockRecorder) StateMinerSectors(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerSectors", reflect.TypeOf((*MockFullNode)(nil).StateMinerSectors), arg0, arg1, arg2, arg3)
}

// StateNetworkName mocks base method
func (m *MockFullNode) StateNetworkName(arg0 context.Context) (dtypes.NetworkName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateNetworkName", arg0)
	ret0, _ := ret[0].(dtypes.NetworkName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateNetworkName indicates an expected call of StateNetworkName
func (mr *MockFullNodeMockRecorder) StateNetworkName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateNetworkName", reflect.TypeOf((*MockFullNode)(nil).StateNetworkName), arg0)
}

// StateNetworkVersion mocks base method
func (m *MockFullNode) StateNetworkVersion(arg0 context.Context, arg1 types.TipSetKey) (network.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateNetworkVersion", arg0, arg1)
	ret0, _ := ret[0].(network.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateNetworkVersion indicates an expected call of StateNetworkVersion
func (mr *MockFullNodeMockRecorder) StateNetworkVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateNetworkVersion", reflect.TypeOf((*MockFullNode)(nil).StateNetworkVersion), arg0, arg1)
}

// StateReadState mocks base method
func (m *MockFullNode) StateReadState(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (*api.ActorState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateReadState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.ActorState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateReadState indicates an expected call of StateReadState
func (mr *MockFullNodeMockRecorder) StateReadState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateReadState", reflect.TypeOf((*MockFullNode)(nil).StateReadState), arg0, arg1, arg2)
}

// StateReplay mocks base method
func (m *MockFullNode) StateReplay(arg0 context.Context, arg1 types.TipSetKey, arg2 cid.Cid) (*api.InvocResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateReplay", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.InvocResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateReplay indicates an expected call of StateReplay
func (mr *MockFullNodeMockRecorder) StateReplay(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateReplay", reflect.TypeOf((*MockFullNode)(nil).StateReplay), arg0, arg1, arg2)
}

// StateSearchMsg mocks base method
func (m *MockFullNode) StateSearchMsg(arg0 context.Context, arg1 cid.Cid) (*api.MsgLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSearchMsg", arg0, arg1)
	ret0, _ := ret[0].(*api.MsgLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSearchMsg indicates an expected call of StateSearchMsg
func (mr *MockFullNodeMockRecorder) StateSearchMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSearchMsg", reflect.TypeOf((*MockFullNode)(nil).StateSearchMsg), arg0, arg1)
}

// StateSearchMsgLimited mocks base method
func (m *MockFullNode) StateSearchMsgLimited(arg0 context.Context, arg1 cid.Cid, arg2 abi.ChainEpoch) (*api.MsgLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSearchMsgLimited", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.MsgLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSearchMsgLimited indicates an expected call of StateSearchMsgLimited
func (mr *MockFullNodeMockRecorder) StateSearchMsgLimited(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSearchMsgLimited", reflect.TypeOf((*MockFullNode)(nil).StateSearchMsgLimited), arg0, arg1, arg2)
}

// StateSectorExpiration mocks base method
func (m *MockFullNode) StateSectorExpiration(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (*miner.SectorExpiration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorExpiration", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*miner.SectorExpiration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorExpiration indicates an expected call of StateSectorExpiration
func (mr *MockFullNodeMockRecorder) StateSectorExpiration(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorExpiration", reflect.TypeOf((*MockFullNode)(nil).StateSectorExpiration), arg0, arg1, arg2, arg3)
}

// StateSectorGetInfo mocks base method
func (m *MockFullNode) StateSectorGetInfo(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (*miner.SectorOnChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorGetInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*miner.SectorOnChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorGetInfo indicates an expected call of StateSectorGetInfo
func (mr *MockFullNodeMockRecorder) StateSectorGetInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorGetInfo", reflect.TypeOf((*MockFullNode)(nil).StateSectorGetInfo), arg0, arg1, arg2, arg3)
}

// StateSectorPartition mocks base method
func (m *MockFullNode) StateSectorPartition(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (*miner.SectorLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorPartition", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*miner.SectorLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorPartition indicates an expected call of StateSectorPartition
func (mr *MockFullNodeMockRecorder) StateSectorPartition(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorPartition", reflect.TypeOf((*MockFullNode)(nil).StateSectorPartition), arg0, arg1, arg2, arg3)
}

// StateSectorPreCommitInfo mocks base method
func (m *MockFullNode) StateSectorPreCommitInfo(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (miner.SectorPreCommitOnChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorPreCommitInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(miner.SectorPreCommitOnChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorPreCommitInfo indicates an expected call of StateSectorPreCommitInfo
func (mr *MockFullNodeMockRecorder) StateSectorPreCommitInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorPreCommitInfo", reflect.TypeOf((*MockFullNode)(nil).StateSectorPreCommitInfo), arg0, arg1, arg2, arg3)
}

// StateVMCirculatingSupplyInternal mocks base method
func (m *MockFullNode) StateVMCirculatingSupplyInternal(arg0 context.Context, arg1 types.TipSetKey) (api.CirculatingSupply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateVMCirculatingSupplyInternal", arg0, arg1)
	ret0, _ := ret[0].(api.CirculatingSupply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateVMCirculatingSupplyInternal indicates an expected call of StateVMCirculatingSupplyInternal
func (mr *MockFullNodeMockRecorder) StateVMCirculatingSupplyInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateVMCirculatingSupplyInternal", reflect.TypeOf((*MockFullNode)(nil).StateVMCirculatingSupplyInternal), arg0, arg1)
}

// StateVerifiedClientStatus mocks base method
func (m *MockFullNode) StateVerifiedClientStatus(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateVerifiedClientStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateVerifiedClientStatus indicates an expected call of StateVerifiedClientStatus
func (mr *MockFullNodeMockRecorder) StateVerifiedClientStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateVerifiedClientStatus", reflect.TypeOf((*MockFullNode)(nil).StateVerifiedClientStatus), arg0, arg1, arg2)
}

// StateVerifiedRegistryRootKey mocks base method
func (m *MockFullNode) StateVerifiedRegistryRootKey(arg0 context.Context, arg1 types.TipSetKey) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateVerifiedRegistryRootKey", arg0, arg1)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateVerifiedRegistryRootKey indicates an expected call of StateVerifiedRegistryRootKey
func (mr *MockFullNodeMockRecorder) StateVerifiedRegistryRootKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateVerifiedRegistryRootKey", reflect.TypeOf((*MockFullNode)(nil).StateVerifiedRegistryRootKey), arg0, arg1)
}

// StateVerifierStatus mocks base method
func (m *MockFullNode) StateVerifierStatus(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateVerifierStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateVerifierStatus indicates an expected call of StateVerifierStatus
func (mr *MockFullNodeMockRecorder) StateVerifierStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateVerifierStatus", reflect.TypeOf((*MockFullNode)(nil).StateVerifierStatus), arg0, arg1, arg2)
}

// StateWaitMsg mocks base method
func (m *MockFullNode) StateWaitMsg(arg0 context.Context, arg1 cid.Cid, arg2 uint64) (*api.MsgLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateWaitMsg", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.MsgLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateWaitMsg indicates an expected call of StateWaitMsg
func (mr *MockFullNodeMockRecorder) StateWaitMsg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateWaitMsg", reflect.TypeOf((*MockFullNode)(nil).StateWaitMsg), arg0, arg1, arg2)
}

// StateWaitMsgLimited mocks base method
func (m *MockFullNode) StateWaitMsgLimited(arg0 context.Context, arg1 cid.Cid, arg2 uint64, arg3 abi.ChainEpoch) (*api.MsgLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateWaitMsgLimited", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.MsgLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateWaitMsgLimited indicates an expected call of StateWaitMsgLimited
func (mr *MockFullNodeMockRecorder) StateWaitMsgLimited(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateWaitMsgLimited", reflect.TypeOf((*MockFullNode)(nil).StateWaitMsgLimited), arg0, arg1, arg2, arg3)
}

// SyncCheckBad mocks base method
func (m *MockFullNode) SyncCheckBad(arg0 context.Context, arg1 cid.Cid) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncCheckBad", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncCheckBad indicates an expected call of SyncCheckBad
func (mr *MockFullNodeMockRecorder) SyncCheckBad(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncCheckBad", reflect.TypeOf((*MockFullNode)(nil).SyncCheckBad), arg0, arg1)
}

// SyncCheckpoint mocks base method
func (m *MockFullNode) SyncCheckpoint(arg0 context.Context, arg1 types.TipSetKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncCheckpoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncCheckpoint indicates an expected call of SyncCheckpoint
func (mr *MockFullNodeMockRecorder) SyncCheckpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncCheckpoint", reflect.TypeOf((*MockFullNode)(nil).SyncCheckpoint), arg0, arg1)
}

// SyncIncomingBlocks mocks base method
func (m *MockFullNode) SyncIncomingBlocks(arg0 context.Context) (<-chan *types.BlockHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncIncomingBlocks", arg0)
	ret0, _ := ret[0].(<-chan *types.BlockHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncIncomingBlocks indicates an expected call of SyncIncomingBlocks
func (mr *MockFullNodeMockRecorder) SyncIncomingBlocks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncIncomingBlocks", reflect.TypeOf((*MockFullNode)(nil).SyncIncomingBlocks), arg0)
}

// SyncMarkBad mocks base method
func (m *MockFullNode) SyncMarkBad(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncMarkBad", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncMarkBad indicates an expected call of SyncMarkBad
func (mr *MockFullNodeMockRecorder) SyncMarkBad(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncMarkBad", reflect.TypeOf((*MockFullNode)(nil).SyncMarkBad), arg0, arg1)
}

// SyncState mocks base method
func (m *MockFullNode) SyncState(arg0 context.Context) (*api.SyncState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncState", arg0)
	ret0, _ := ret[0].(*api.SyncState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncState indicates an expected call of SyncState
func (mr *MockFullNodeMockRecorder) SyncState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncState", reflect.TypeOf((*MockFullNode)(nil).SyncState), arg0)
}

// SyncSubmitBlock mocks base method
func (m *MockFullNode) SyncSubmitBlock(arg0 context.Context, arg1 *types.BlockMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncSubmitBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncSubmitBlock indicates an expected call of SyncSubmitBlock
func (mr *MockFullNodeMockRecorder) SyncSubmitBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncSubmitBlock", reflect.TypeOf((*MockFullNode)(nil).SyncSubmitBlock), arg0, arg1)
}

// SyncUnmarkAllBad mocks base method
func (m *MockFullNode) SyncUnmarkAllBad(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncUnmarkAllBad", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncUnmarkAllBad indicates an expected call of SyncUnmarkAllBad
func (mr *MockFullNodeMockRecorder) SyncUnmarkAllBad(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncUnmarkAllBad", reflect.TypeOf((*MockFullNode)(nil).SyncUnmarkAllBad), arg0)
}

// SyncUnmarkBad mocks base method
func (m *MockFullNode) SyncUnmarkBad(arg0 context.Context, arg1 cid.Cid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncUnmarkBad", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncUnmarkBad indicates an expected call of SyncUnmarkBad
func (mr *MockFullNodeMockRecorder) SyncUnmarkBad(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncUnmarkBad", reflect.TypeOf((*MockFullNode)(nil).SyncUnmarkBad), arg0, arg1)
}

// SyncValidateTipset mocks base method
func (m *MockFullNode) SyncValidateTipset(arg0 context.Context, arg1 types.TipSetKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncValidateTipset", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncValidateTipset indicates an expected call of SyncValidateTipset
func (mr *MockFullNodeMockRecorder) SyncValidateTipset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncValidateTipset", reflect.TypeOf((*MockFullNode)(nil).SyncValidateTipset), arg0, arg1)
}

// Version mocks base method
func (m *MockFullNode) Version(arg0 context.Context) (api.APIVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(api.APIVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockFullNodeMockRecorder) Version(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockFullNode)(nil).Version), arg0)
}

// WalletBalance mocks base method
func (m *MockFullNode) WalletBalance(arg0 context.Context, arg1 address.Address) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletBalance", arg0, arg1)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletBalance indicates an expected call of WalletBalance
func (mr *MockFullNodeMockRecorder) WalletBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletBalance", reflect.TypeOf((*MockFullNode)(nil).WalletBalance), arg0, arg1)
}

// WalletDefaultAddress mocks base method
func (m *MockFullNode) WalletDefaultAddress(arg0 context.Context) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletDefaultAddress", arg0)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletDefaultAddress indicates an expected call of WalletDefaultAddress
func (mr *MockFullNodeMockRecorder) WalletDefaultAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletDefaultAddress", reflect.TypeOf((*MockFullNode)(nil).WalletDefaultAddress), arg0)
}

// WalletDelete mocks base method
func (m *MockFullNode) WalletDelete(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalletDelete indicates an expected call of WalletDelete
func (mr *MockFullNodeMockRecorder) WalletDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletDelete", reflect.TypeOf((*MockFullNode)(nil).WalletDelete), arg0, arg1)
}

// WalletExport mocks base method
func (m *MockFullNode) WalletExport(arg0 context.Context, arg1 address.Address) (*types.KeyInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletExport", arg0, arg1)
	ret0, _ := ret[0].(*types.KeyInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletExport indicates an expected call of WalletExport
func (mr *MockFullNodeMockRecorder) WalletExport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletExport", reflect.TypeOf((*MockFullNode)(nil).WalletExport), arg0, arg1)
}

// WalletHas mocks base method
func (m *MockFullNode) WalletHas(arg0 context.Context, arg1 address.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletHas", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletHas indicates an expected call of WalletHas
func (mr *MockFullNodeMockRecorder) WalletHas(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletHas", reflect.TypeOf((*MockFullNode)(nil).WalletHas), arg0, arg1)
}

// WalletImport mocks base method
func (m *MockFullNode) WalletImport(arg0 context.Context, arg1 *types.KeyInfo) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletImport", arg0, arg1)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletImport indicates an expected call of WalletImport
func (mr *MockFullNodeMockRecorder) WalletImport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletImport", reflect.TypeOf((*MockFullNode)(nil).WalletImport), arg0, arg1)
}

// WalletList mocks base method
func (m *MockFullNode) WalletList(arg0 context.Context) ([]address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletList", arg0)
	ret0, _ := ret[0].([]address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletList indicates an expected call of WalletList
func (mr *MockFullNodeMockRecorder) WalletList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletList", reflect.TypeOf((*MockFullNode)(nil).WalletList), arg0)
}

// WalletNew mocks base method
func (m *MockFullNode) WalletNew(arg0 context.Context, arg1 types.KeyType) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletNew", arg0, arg1)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletNew indicates an expected call of WalletNew
func (mr *MockFullNodeMockRecorder) WalletNew(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletNew", reflect.TypeOf((*MockFullNode)(nil).WalletNew), arg0, arg1)
}

// WalletSetDefault mocks base method
func (m *MockFullNode) WalletSetDefault(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletSetDefault", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalletSetDefault indicates an expected call of WalletSetDefault
func (mr *MockFullNodeMockRecorder) WalletSetDefault(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletSetDefault", reflect.TypeOf((*MockFullNode)(nil).WalletSetDefault), arg0, arg1)
}

// WalletSign mocks base method
func (m *MockFullNode) WalletSign(arg0 context.Context, arg1 address.Address, arg2 []byte) (*crypto.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletSign", arg0, arg1, arg2)
	ret0, _ := ret[0].(*crypto.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletSign indicates an expected call of WalletSign
func (mr *MockFullNodeMockRecorder) WalletSign(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletSign", reflect.TypeOf((*MockFullNode)(nil).WalletSign), arg0, arg1, arg2)
}

// WalletSignMessage mocks base method
func (m *MockFullNode) WalletSignMessage(arg0 context.Context, arg1 address.Address, arg2 *types.Message) (*types.SignedMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletSignMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletSignMessage indicates an expected call of WalletSignMessage
func (mr *MockFullNodeMockRecorder) WalletSignMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletSignMessage", reflect.TypeOf((*MockFullNode)(nil).WalletSignMessage), arg0, arg1, arg2)
}

// WalletValidateAddress mocks base method
func (m *MockFullNode) WalletValidateAddress(arg0 context.Context, arg1 string) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletValidateAddress", arg0, arg1)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletValidateAddress indicates an expected call of WalletValidateAddress
func (mr *MockFullNodeMockRecorder) WalletValidateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletValidateAddress", reflect.TypeOf((*MockFullNode)(nil).WalletValidateAddress), arg0, arg1)
}

// WalletVerify mocks base method
func (m *MockFullNode) WalletVerify(arg0 context.Context, arg1 address.Address, arg2 []byte, arg3 *crypto.Signature) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletVerify", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletVerify indicates an expected call of WalletVerify
func (mr *MockFullNodeMockRecorder) WalletVerify(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletVerify", reflect.TypeOf((*MockFullNode)(nil).WalletVerify), arg0, arg1, arg2, arg3)
}
