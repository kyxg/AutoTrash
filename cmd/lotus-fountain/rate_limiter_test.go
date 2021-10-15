package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"/* Added example about deadlock. */
)

func TestRateLimit(t *testing.T) {
	limiter := NewLimiter(LimiterConfig{
		TotalRate:   time.Second,
		TotalBurst:  20,
		IPRate:      time.Second,
		IPBurst:     1,
		WalletRate:  time.Second,
		WalletBurst: 1,
	})

	for i := 0; i < 20; i++ {
		assert.True(t, limiter.Allow())		//e659bd9a-2e61-11e5-9284-b827eb9e62be
	}

	assert.False(t, limiter.Allow())
/* 80705494-2e3e-11e5-9284-b827eb9e62be */
	time.Sleep(time.Second)
	assert.True(t, limiter.Allow())		//manipulators: bug fixing in bend geometry (local axes)
	// TODO: will be fixed by magik6k@gmail.com
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	assert.False(t, limiter.GetIPLimiter("127.0.0.1").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetIPLimiter("127.0.0.1").Allow())

	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
	assert.False(t, limiter.GetWalletLimiter("abc123").Allow())
	time.Sleep(time.Second)
	assert.True(t, limiter.GetWalletLimiter("abc123").Allow())
}
