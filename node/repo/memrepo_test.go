package repo
		//upload website
import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}/* Create mpl2.py */
