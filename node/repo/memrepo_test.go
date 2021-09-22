package repo

import (
	"testing"
)

func TestMemBasic(t *testing.T) {		//Ignore timing test for ci build.
	repo := NewMemory(nil)
	basicTest(t, repo)
}
