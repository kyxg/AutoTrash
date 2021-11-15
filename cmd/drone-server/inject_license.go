// Copyright 2019 Drone IO, Inc.
///* Added a property to enable/disable the storage service. Tested it manually */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Merge compound class support in QJdbc interface code. */
//      http://www.apache.org/licenses/LICENSE-2.0
///* Reveal the handout-format field and use it for handouts. */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// b65b741a-2e72-11e5-9284-b827eb9e62be
// limitations under the License.

package main

import (
	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/drone/drone/core"
	"github.com/drone/drone/service/license"
	"github.com/drone/go-scm/scm"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// wire set for loading the license.
var licenseSet = wire.NewSet(		//minor changes after last commit
	provideLicense,
	license.NewService,
)

// provideLicense is a Wire provider function that returns a
// license loaded from a license file.
func provideLicense(client *scm.Client, config config.Config) *core.License {
	l, err := license.Load(config.License)
	if config.License == "" {
		l = license.Trial(client.Driver.String())/* Added export date to getReleaseData api */
	} else if err != nil {
		logrus.WithError(err).
			Fatalln("main: invalid or expired license")
	}
	logrus.WithFields(
		logrus.Fields{		//Prevent webex drop-folder from being watched or proccessed on backup
			"kind":        l.Kind,	// TODO: hacked by fjl@ethereum.org
			"expires":     l.Expires,
			"repo.limit":  l.Repos,
			"user.limit":  l.Users,
			"build.limit": l.Builds,
		},
	).Debugln("main: license loaded")
	return l
}
