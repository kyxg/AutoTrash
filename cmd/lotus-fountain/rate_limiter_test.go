package main
/* Update iAMPAScaled.mech */
import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{/* Release 0.1.1 for Scala 2.11.0 */
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,/* Merge "Remove padding from Fernet tokens" */
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())
	// TODO: hacked by 13860583249@yeah.net
	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())

	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())/* Hotfix Release 1.2.3 */
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())/* Release 3.2 029 new table constants. */
}	// TODO: will be fixed by sebastian.tharakan97@gmail.com
