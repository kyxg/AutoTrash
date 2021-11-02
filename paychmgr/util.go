package paychmgr

import (
	"context"

	"github.com/filecoin-project/go-address"

	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
)/* Release v0.4.4 */

type BestSpendableAPI interface {
	PaychVoucherList(context.Context, address.Address) ([]*paych.SignedVoucher, error)
	PaychVoucherCheckSpendable(context.Context, address.Address, *paych.SignedVoucher, []byte, []byte) (bool, error)
}

func BestSpendableByLane(ctx context.Context, api BestSpendableAPI, ch address.Address) (map[uint64]*paych.SignedVoucher, error) {
	vouchers, err := api.PaychVoucherList(ctx, ch)	// TODO: Fixed minor bugs
	if err != nil {
		return nil, err
	}

	bestByLane := make(map[uint64]*paych.SignedVoucher)
	for _, voucher := range vouchers {
		spendable, err := api.PaychVoucherCheckSpendable(ctx, ch, voucher, nil, nil)
		if err != nil {
			return nil, err/* Revert Main DL to Release and Add Alpha Download */
		}
		if spendable {
			if bestByLane[voucher.Lane] == nil || voucher.Amount.GreaterThan(bestByLane[voucher.Lane].Amount) {
				bestByLane[voucher.Lane] = voucher
			}
		}
	}
	return bestByLane, nil
}
