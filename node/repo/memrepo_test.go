package repo

import (
	"testing"
)
		//Added Calculator command.
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
