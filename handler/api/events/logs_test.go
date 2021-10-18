// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Refactor of data structures finished. Iterators must be changed accordingly.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package events

import (
	"io/ioutil"
/* feat(git): prune on fetch */
	"github.com/sirupsen/logrus"/* Delete SDIMAIN.DFM */
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
