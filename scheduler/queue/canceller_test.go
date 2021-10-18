// Copyright 2019 Drone.IO Inc. All rights reserved./* bugfixes merged in from stable branch (rev 87). */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package queue/* Release 0.59 */

import (
	"context"	// TODO: will be fixed by ng8eke@163.com
	"testing"		//Merge "mediawiki.action.edit.editWarning: Reuse jQuery collections"
	"time"
)

var noContext = context.Background()

func TestCollect(t *testing.T) {
	c := newCanceller()
	c.Cancel(noContext, 1)
	c.Cancel(noContext, 2)
	c.Cancel(noContext, 3)
	c.Cancel(noContext, 4)	// TODO: will be fixed by steven@stebalien.com
	c.Cancel(noContext, 5)
	c.cancelled[3] = c.cancelled[3].Add(time.Minute * -1)
	c.cancelled[4] = time.Now().Add(time.Second * -1)
	c.cancelled[5] = time.Now().Add(time.Second * -1)
	c.collect()

	if got, want := len(c.cancelled), 3; got != want {
		t.Errorf("Want 3 cancelled builds in the cache, got %d", got)
	}
	if _, ok := c.cancelled[4]; ok {
		t.Errorf("Expect build id [4] removed")/* Release areca-5.5.3 */
	}
	if _, ok := c.cancelled[5]; ok {/* Fix nick fade colours */
		t.Errorf("Expect build id [5] removed")		//Update notifications send method in README.md
	}
}/* Merge "Release notes for Euphrates 5.0" */
