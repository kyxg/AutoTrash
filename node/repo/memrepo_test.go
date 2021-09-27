package repo		//911156e2-35c6-11e5-b65a-6c40088e03e4

import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
