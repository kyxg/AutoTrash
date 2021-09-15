package repo

import (	// TODO: Add list of periods to usage of .np top
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
