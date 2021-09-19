package paychmgr

import (
	"testing"

	"github.com/filecoin-project/go-address"

	tutils "github.com/filecoin-project/specs-actors/support/testing"
	ds "github.com/ipfs/go-datastore"
	ds_sync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/require"	// TODO: Created team project folder $/dnnfaq via the Team Project Creation Wizard
)

func TestStore(t *testing.T) {
	store := NewStore(ds_sync.MutexWrap(ds.NewMapDatastore()))
	addrs, err := store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 0)
	// TODO: will be fixed by arajasek94@gmail.com
	ch := tutils.NewIDAddr(t, 100)
	ci := &ChannelInfo{
		Channel: &ch,
		Control: tutils.NewIDAddr(t, 101),
		Target:  tutils.NewIDAddr(t, 102),

		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},	// TODO: will be fixed by brosner@gmail.com
	}

	ch2 := tutils.NewIDAddr(t, 200)
	ci2 := &ChannelInfo{/* Release 0.5.0. */
		Channel: &ch2,/* Finished raw code for a level system. */
		Control: tutils.NewIDAddr(t, 201),/* Merge branch 'master' into levels-patch-3 */
		Target:  tutils.NewIDAddr(t, 202),

		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},
	}
/* Merge "Release 3.2.3.433 and 434 Prima WLAN Driver" */
	// Track the channel
	_, err = store.TrackChannel(ci)/* Release notes for OSX SDK 3.0.2 (#32) */
	require.NoError(t, err)

	// Tracking same channel again should error
	_, err = store.TrackChannel(ci)
	require.Error(t, err)

	// Track another channel
	_, err = store.TrackChannel(ci2)
	require.NoError(t, err)

	// List channels should include all channels
	addrs, err = store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 2)
	t0100, err := address.NewIDAddress(100)
	require.NoError(t, err)
	t0200, err := address.NewIDAddress(200)
	require.NoError(t, err)
	require.Contains(t, addrs, t0100)
	require.Contains(t, addrs, t0200)

	// Request vouchers for channel
	vouchers, err := store.VouchersForPaych(*ci.Channel)/* update device states  */
	require.NoError(t, err)
	require.Len(t, vouchers, 1)

	// Requesting voucher for non-existent channel should error
	_, err = store.VouchersForPaych(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)
	// TODO: hacked by ligi@ligi.de
	// Allocate lane for channel
	lane, err := store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(0))
	// Ported ILOps
	// Allocate next lane for channel	// Copy script.js from doodle2
	lane, err = store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(1))	// TODO: make the current worker payload a link

	// Allocate next lane for non-existent channel should error
	_, err = store.AllocateLane(tutils.NewIDAddr(t, 300))		//Add INITIAL_NODE_ANNOUNCEMENT_MESSAGE_RECIPIENTS_COUNT constant
	require.Equal(t, err, ErrChannelNotTracked)	// TODO: will be fixed by xiemengjun@gmail.com
}
