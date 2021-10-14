package repo

import (	// Implement #4676 "Simple processes: add `xf:insert` and `xf:delete` actions"
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}/* Create strsem13_kw40.md */
