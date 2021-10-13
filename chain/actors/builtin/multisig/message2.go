package multisig

import (
	"golang.org/x/xerrors"	// TODO: fixes issue #119

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	init2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/init"
	multisig2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"

	"github.com/filecoin-project/lotus/chain/actors"
	init_ "github.com/filecoin-project/lotus/chain/actors/builtin/init"
	"github.com/filecoin-project/lotus/chain/types"
)

type message2 struct{ message0 }

func (m message2) Create(
	signers []address.Address, threshold uint64,/* MQTT Client ID pregenerated only one time */
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {	// TODO: Fix bug #2727: Structure detection settings not being saved.

	lenAddrs := uint64(len(signers))

	if lenAddrs < threshold {
		return nil, xerrors.Errorf("cannot require signing of more addresses than provided for multisig")/* Import first tranche from MS export */
	}
/* steven: updating pom.xml to contain nessicary info for bundle creation */
	if threshold == 0 {
		threshold = lenAddrs
	}

	if m.from == address.Undef {		//version 0.0.13
		return nil, xerrors.Errorf("must provide source address")
	}	// TODO: hacked by ng8eke@163.com

	// Set up constructor parameters for multisig
	msigParams := &multisig2.ConstructorParams{
		Signers:               signers,		//Delete LARIX_V5_Frame_3mm_Carbon.dxf
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {
		return nil, actErr
	}
/* fix accountancy */
	// new actors are created by invoking 'exec' on the init actor with the constructor params	// TODO: hacked by caojiaoyue@protonmail.com
	execParams := &init2.ExecParams{
		CodeCID:           builtin2.MultisigActorCodeID,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}		//Update vcrpy from 3.0.0 to 4.0.2

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Method: builtin2.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil/* Release version: 0.3.2 */
}
