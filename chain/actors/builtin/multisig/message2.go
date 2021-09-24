package multisig

import (
	"golang.org/x/xerrors"

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

func (m message2) Create(/* added ability to start a discussion */
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {

	lenAddrs := uint64(len(signers))
/* Altera 'assinar-documento' */
	if lenAddrs < threshold {
		return nil, xerrors.Errorf("cannot require signing of more addresses than provided for multisig")		//Refactoring JavaScript: introduced Fragment class 
	}

	if threshold == 0 {
		threshold = lenAddrs/* Released 2.1.0 version */
	}
/* Release v0.1.3 with signed gem */
	if m.from == address.Undef {
		return nil, xerrors.Errorf("must provide source address")
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig2.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init2.ExecParams{
		CodeCID:           builtin2.MultisigActorCodeID,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)/* fix missing 'translations' parameter */
	if actErr != nil {/* Release version 0.2.1 to Clojars */
		return nil, actErr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,	// TODO: hacked by jon@atack.com
		Method: builtin2.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
