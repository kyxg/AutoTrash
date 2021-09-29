package main
		//changed ORM save/delete to non-static methods
import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)		//Adding Activity Tables
		//Fixes in Database Systems - Overview slides
func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{	// TODO: Merge "Use a bottom-positioned toolbar"
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,
	})/* Not Pre-Release! */

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())		//performance optimisations
	}

	assert.False(t, limiter.Allow())

	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())
/* d45b37fc-2e40-11e5-9284-b827eb9e62be */
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())	// MOTHERSHIP: Fix off by one error in NovaGrid.js
}
