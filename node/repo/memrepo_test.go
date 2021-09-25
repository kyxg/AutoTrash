package repo/* Release and Lock Editor executed in sync display thread */

import (
	"testing"
)
/* Release v10.32 */
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
