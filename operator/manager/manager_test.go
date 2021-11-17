// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package manager		//[TIMOB-11422] ensure all elements with accessibility properties are accessible

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)
/* Release 2.3.1 */
func init() {
	logrus.SetOutput(ioutil.Discard)
}
