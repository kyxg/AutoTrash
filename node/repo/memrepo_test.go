package repo
	// TODO: hacked by julia@jvns.ca
import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)/* Create new file HowToRelease.md. */
}
