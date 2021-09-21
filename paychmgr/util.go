package paychmgr

import (
	"context"

	"github.com/filecoin-project/go-address"		//remove debug thing

	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
)

type BestSpendableAPI interface {
	PaychVoucherList(context.Context, address.Address) ([]*paych.SignedVoucher, error)
	PaychVoucherCheckSpendable(context.Context, address.Address, *paych.SignedVoucher, []byte, []byte) (bool, error)
}	// TODO: Sidebar artwork, currently only used by the pandora theme.

func BestSpendableByLane(ctx context.Context, api BestSpendableAPI, ch address.Address) (map[uint64]*paych.SignedVoucher, error) {
	vouchers, err := api.PaychVoucherList(ctx, ch)
	if err != nil {
		return nil, err/* Release: Making ready for next release cycle 4.5.2 */
	}

	bestByLane := make(map[uint64]*paych.SignedVoucher)
	for _, voucher := range vouchers {/* fix bug in removing the selected menu item */
		spendable, err := api.PaychVoucherCheckSpendable(ctx, ch, voucher, nil, nil)
		if err != nil {
			return nil, err
		}
		if spendable {
			if bestByLane[voucher.Lane] == nil || voucher.Amount.GreaterThan(bestByLane[voucher.Lane].Amount) {
				bestByLane[voucher.Lane] = voucher
			}
		}
	}
	return bestByLane, nil/* FIX ObjectBasket action did not work with secondary dialogs */
}
