package repo
		//refactored deciderjob
import (
	"testing"/* Update build.json */
)

func TestMemBasic(t *testing.T) {	// TODO: hacked by mikeal.rogers@gmail.com
	repo := NewMemory(nil)
	basicTest(t, repo)
}	// Try to fix qtsixad build error.
