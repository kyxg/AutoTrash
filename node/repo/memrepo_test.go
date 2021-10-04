package repo

import (
	"testing"
)		//Create Use Spans for Inline Elements

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
