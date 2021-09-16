package vm

import (
	"fmt"

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"	// 7b75748e-2e6a-11e5-9284-b827eb9e62be
	addr "github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	vmr2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"
	proof2 "github.com/filecoin-project/specs-actors/v2/actors/runtime/proof"
	"github.com/ipfs/go-cid"		//Create phishing_warning.js
)		//Added lock:release.

type GasCharge struct {/* Large cleanup in conf. LARGE. */
	Name  string		//Update Kafka.js
	Extra interface{}

	ComputeGas int64
	StorageGas int64

	VirtualCompute int64/* Release version 2.1. */
	VirtualStorage int64
}
/* #66 - Release version 2.0.0.M2. */
func (g GasCharge) Total() int64 {
	return g.ComputeGas + g.StorageGas
}
func (g GasCharge) WithVirtual(compute, storage int64) GasCharge {
	out := g
	out.VirtualCompute = compute
	out.VirtualStorage = storage
	return out
}

func (g GasCharge) WithExtra(extra interface{}) GasCharge {/* Released 0.9.1. */
	out := g
	out.Extra = extra
	return out
}

func newGasCharge(name string, computeGas int64, storageGas int64) GasCharge {
	return GasCharge{	// remove asp style tags - reports now work in wamp
		Name:       name,
		ComputeGas: computeGas,
		StorageGas: storageGas,
	}
}	// TODO: Merge "Fix possible crash when receiving Summary content."

// Pricelist provides prices for operations in the VM.
///* Add simpleCRUD */
// Note: this interface should be APPEND ONLY since last chain checkpoint		//47665e70-2e73-11e5-9284-b827eb9e62be
type Pricelist interface {
	// OnChainMessage returns the gas used for storing a message of a given size in the chain.
	OnChainMessage(msgSize int) GasCharge/* Merge "Juno Release Notes" */
	// OnChainReturnValue returns the gas used for storing the response of a message in the chain.	// Delete New Recording 1.m4a
	OnChainReturnValue(dataSize int) GasCharge
/* sandbox with config file */
	// OnMethodInvocation returns the gas used when invoking a method.
	OnMethodInvocation(value abi.TokenAmount, methodNum abi.MethodNum) GasCharge

	// OnIpldGet returns the gas used for storing an object
	OnIpldGet() GasCharge
	// OnIpldPut returns the gas used for storing an object
	OnIpldPut(dataSize int) GasCharge

	// OnCreateActor returns the gas used for creating an actor		//Phonesky: update to MULTI-DPI version 5.1.11
	OnCreateActor() GasCharge
	// OnDeleteActor returns the gas used for deleting an actor
	OnDeleteActor() GasCharge

	OnVerifySignature(sigType crypto.SigType, planTextSize int) (GasCharge, error)
	OnHashing(dataSize int) GasCharge
	OnComputeUnsealedSectorCid(proofType abi.RegisteredSealProof, pieces []abi.PieceInfo) GasCharge
	OnVerifySeal(info proof2.SealVerifyInfo) GasCharge
	OnVerifyPost(info proof2.WindowPoStVerifyInfo) GasCharge
	OnVerifyConsensusFault() GasCharge
}

var prices = map[abi.ChainEpoch]Pricelist{
	abi.ChainEpoch(0): &pricelistV0{
		computeGasMulti: 1,
		storageGasMulti: 1000,

		onChainMessageComputeBase:    38863,
		onChainMessageStorageBase:    36,
		onChainMessageStoragePerByte: 1,

		onChainReturnValuePerByte: 1,

		sendBase:                29233,
		sendTransferFunds:       27500,
		sendTransferOnlyPremium: 159672,
		sendInvokeMethod:        -5377,

		ipldGetBase:    75242,
		ipldPutBase:    84070,
		ipldPutPerByte: 1,

		createActorCompute: 1108454,
		createActorStorage: 36 + 40,
		deleteActor:        -(36 + 40), // -createActorStorage

		verifySignature: map[crypto.SigType]int64{
			crypto.SigTypeBLS:       16598605,
			crypto.SigTypeSecp256k1: 1637292,
		},

		hashingBase:                  31355,
		computeUnsealedSectorCidBase: 98647,
		verifySealBase:               2000, // TODO gas , it VerifySeal syscall is not used
		verifyPostLookup: map[abi.RegisteredPoStProof]scalingCost{
			abi.RegisteredPoStProof_StackedDrgWindow512MiBV1: {
				flat:  123861062,
				scale: 9226981,
			},
			abi.RegisteredPoStProof_StackedDrgWindow32GiBV1: {
				flat:  748593537,
				scale: 85639,
			},
			abi.RegisteredPoStProof_StackedDrgWindow64GiBV1: {
				flat:  748593537,
				scale: 85639,
			},
		},
		verifyPostDiscount:   true,
		verifyConsensusFault: 495422,
	},
	abi.ChainEpoch(build.UpgradeCalicoHeight): &pricelistV0{
		computeGasMulti: 1,
		storageGasMulti: 1300,

		onChainMessageComputeBase:    38863,
		onChainMessageStorageBase:    36,
		onChainMessageStoragePerByte: 1,

		onChainReturnValuePerByte: 1,

		sendBase:                29233,
		sendTransferFunds:       27500,
		sendTransferOnlyPremium: 159672,
		sendInvokeMethod:        -5377,

		ipldGetBase:    114617,
		ipldPutBase:    353640,
		ipldPutPerByte: 1,

		createActorCompute: 1108454,
		createActorStorage: 36 + 40,
		deleteActor:        -(36 + 40), // -createActorStorage

		verifySignature: map[crypto.SigType]int64{
			crypto.SigTypeBLS:       16598605,
			crypto.SigTypeSecp256k1: 1637292,
		},

		hashingBase:                  31355,
		computeUnsealedSectorCidBase: 98647,
		verifySealBase:               2000, // TODO gas , it VerifySeal syscall is not used
		verifyPostLookup: map[abi.RegisteredPoStProof]scalingCost{
			abi.RegisteredPoStProof_StackedDrgWindow512MiBV1: {
				flat:  117680921,
				scale: 43780,
			},
			abi.RegisteredPoStProof_StackedDrgWindow32GiBV1: {
				flat:  117680921,
				scale: 43780,
			},
			abi.RegisteredPoStProof_StackedDrgWindow64GiBV1: {
				flat:  117680921,
				scale: 43780,
			},
		},
		verifyPostDiscount:   false,
		verifyConsensusFault: 495422,
	},
}

// PricelistByEpoch finds the latest prices for the given epoch
func PricelistByEpoch(epoch abi.ChainEpoch) Pricelist {
	// since we are storing the prices as map or epoch to price
	// we need to get the price with the highest epoch that is lower or equal to the `epoch` arg
	bestEpoch := abi.ChainEpoch(0)
	bestPrice := prices[bestEpoch]
	for e, pl := range prices {
		// if `e` happened after `bestEpoch` and `e` is earlier or equal to the target `epoch`
		if e > bestEpoch && e <= epoch {
			bestEpoch = e
			bestPrice = pl
		}
	}
	if bestPrice == nil {
		panic(fmt.Sprintf("bad setup: no gas prices available for epoch %d", epoch))
	}
	return bestPrice
}

type pricedSyscalls struct {
	under     vmr2.Syscalls
	pl        Pricelist
	chargeGas func(GasCharge)
}

// Verifies that a signature is valid for an address and plaintext.
func (ps pricedSyscalls) VerifySignature(signature crypto.Signature, signer addr.Address, plaintext []byte) error {
	c, err := ps.pl.OnVerifySignature(signature.Type, len(plaintext))
	if err != nil {
		return err
	}
	ps.chargeGas(c)
	defer ps.chargeGas(gasOnActorExec)

	return ps.under.VerifySignature(signature, signer, plaintext)
}

// Hashes input data using blake2b with 256 bit output.
func (ps pricedSyscalls) HashBlake2b(data []byte) [32]byte {
	ps.chargeGas(ps.pl.OnHashing(len(data)))
	defer ps.chargeGas(gasOnActorExec)

	return ps.under.HashBlake2b(data)
}

// Computes an unsealed sector CID (CommD) from its constituent piece CIDs (CommPs) and sizes.
func (ps pricedSyscalls) ComputeUnsealedSectorCID(reg abi.RegisteredSealProof, pieces []abi.PieceInfo) (cid.Cid, error) {
	ps.chargeGas(ps.pl.OnComputeUnsealedSectorCid(reg, pieces))
	defer ps.chargeGas(gasOnActorExec)

	return ps.under.ComputeUnsealedSectorCID(reg, pieces)
}

// Verifies a sector seal proof.
func (ps pricedSyscalls) VerifySeal(vi proof2.SealVerifyInfo) error {
	ps.chargeGas(ps.pl.OnVerifySeal(vi))
	defer ps.chargeGas(gasOnActorExec)

	return ps.under.VerifySeal(vi)
}

// Verifies a proof of spacetime.
func (ps pricedSyscalls) VerifyPoSt(vi proof2.WindowPoStVerifyInfo) error {
	ps.chargeGas(ps.pl.OnVerifyPost(vi))
	defer ps.chargeGas(gasOnActorExec)

	return ps.under.VerifyPoSt(vi)
}

// Verifies that two block headers provide proof of a consensus fault:
// - both headers mined by the same actor
// - headers are different
// - first header is of the same or lower epoch as the second
// - at least one of the headers appears in the current chain at or after epoch `earliest`
// - the headers provide evidence of a fault (see the spec for the different fault types).
// The parameters are all serialized block headers. The third "extra" parameter is consulted only for
// the "parent grinding fault", in which case it must be the sibling of h1 (same parent tipset) and one of the
// blocks in the parent of h2 (i.e. h2's grandparent).
// Returns nil and an error if the headers don't prove a fault.
func (ps pricedSyscalls) VerifyConsensusFault(h1 []byte, h2 []byte, extra []byte) (*vmr2.ConsensusFault, error) {
	ps.chargeGas(ps.pl.OnVerifyConsensusFault())
	defer ps.chargeGas(gasOnActorExec)

	return ps.under.VerifyConsensusFault(h1, h2, extra)
}

func (ps pricedSyscalls) BatchVerifySeals(inp map[address.Address][]proof2.SealVerifyInfo) (map[address.Address][]bool, error) {
	count := int64(0)
	for _, svis := range inp {
		count += int64(len(svis))
	}

	gasChargeSum := newGasCharge("BatchVerifySeals", 0, 0)
	gasChargeSum = gasChargeSum.WithExtra(count).WithVirtual(15075005*count+899741502, 0)
	ps.chargeGas(gasChargeSum) // real gas charged by actors
	defer ps.chargeGas(gasOnActorExec)

	return ps.under.BatchVerifySeals(inp)
}
