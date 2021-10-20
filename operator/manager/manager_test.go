// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager
/* Release code under MIT License */
import (
	"io/ioutil"
	// Toggle the pinkynails properly, props goto10, fixes #17212
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)/* Update Release Notes for JIRA step */
}
