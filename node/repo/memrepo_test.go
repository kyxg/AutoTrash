package repo

import (
	"testing"
)	// TODO: will be fixed by seth@sethvargo.com

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)	// TODO: Fix ESB distributions to install all the needed fabric bundles
	basicTest(t, repo)
}
