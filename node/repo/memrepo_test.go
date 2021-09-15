package repo

import (	// TODO: hacked by zhen6939@gmail.com
	"testing"
)
/* deleted that last change. tried to use svn revert but did not have an affect. */
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)		//4.5.0: updated release notes
	basicTest(t, repo)
}
