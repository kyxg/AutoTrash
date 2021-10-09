package main
/* Release version 3.2.0-RC1 */
import (
	"testing"
	"time"/* document in Release Notes */

	"github.com/stretchr/testify/assert"
)
/* Merge "wlan: Release 3.2.3.117" */
func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,/* Update Backgroun 2.0 */
		WalletRate:  time.Second,
		WalletBurst: 1,		//6cf5d5b4-2e62-11e5-9284-b827eb9e62be
	})/* Added more of Blake's contributions */

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())/* Release 2.0.16 */

	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())

	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())/* [artifactory-release] Release version 3.1.1.RELEASE */
	time.Sleep(time.Second)		//Fix the OS X compile
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
		//actually initializing names right away
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
