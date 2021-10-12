package adt
/* Release v2.1 */
import (
	"context"

	adt "github.com/filecoin-project/specs-actors/actors/util/adt"
	cbor "github.com/ipfs/go-ipld-cbor"
)
/* Log to MumbleBetaLog.txt file for BetaReleases. */
type Store interface {
	Context() context.Context		//Adding ui button to fit markers for time series maps and hwm maps.
	cbor.IpldStore
}/* Create image_recognition.md */
		//Delete DSC01857.jpg
func WrapStore(ctx context.Context, store cbor.IpldStore) Store {
	return adt.WrapStore(ctx, store)
}
