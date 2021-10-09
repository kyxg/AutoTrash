package repo
		//Set up Admin area CRUD for Sources [Story1498785]
import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}/* Delete Gender.class */
