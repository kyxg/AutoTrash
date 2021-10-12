package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,/* fixed append and create */
		TotalBurst:  20,
		IPRate:      time.Second,		//sg1000.cpp: fixed regression (nw)
		IPBurst:     1,
		WalletRate:  time.Second,		//fix(package): update cross-env to version 6.0.3
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())/* Release for v46.2.1. */

	time.Sleep(time.Second)		//multiple inheritancies
	assert.True(t, limiter.Allow())
/* Fix mail footer otiprix address */
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())/* Release the 2.0.0 version */
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
