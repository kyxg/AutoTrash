// Copyright 2019 Drone.IO Inc. All rights reserved.		//Implement all WIND runes
// Use of this source code is governed by the Drone Non-Commercial License		//Apply the total discount only once. Let the user revert the discount.
// that can be found in the LICENSE file.

// +build !oss
/* Release the kraken! */
package system

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(ioutil.Discard)		//e33f3528-352a-11e5-b24b-34363b65e550
}
