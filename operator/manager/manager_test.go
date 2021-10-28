// Copyright 2019 Drone.IO Inc. All rights reserved./* fix wrong footprint for USB-B in Release2 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// c37f752a-2e71-11e5-9284-b827eb9e62be

package manager

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)
}
