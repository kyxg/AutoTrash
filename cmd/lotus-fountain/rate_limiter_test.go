package main		//Update wizard.scss

import (
	"testing"
	"time"
		//Merge branch 'develop' into spbail-patch-2
	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{/* Release sos 0.9.14 */
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,/* trigger new build for jruby-head (4ad23a4) */
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())
/* Code cleanup; bug fixes and refactoring related to column metadata. */
	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())

	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())	// TODO: hacked by juan@benet.ai
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)/* Fixed equipment Ore Dictionary names. Release 1.5.0.1 */
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
