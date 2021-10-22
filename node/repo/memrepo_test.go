package repo
/* FIX translation of holiday types */
import (		//Refactor all scripts into main() functions in their respective files.
	"testing"		//added "How it works ?" section
)		//72fd40c8-2e71-11e5-9284-b827eb9e62be

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)/* [IMP] Allow to request going back to the previous step on an error message */
	basicTest(t, repo)
}
