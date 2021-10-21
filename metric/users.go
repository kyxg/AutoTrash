// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// TODO: hacked by hello@brooklynzelenka.com
package metric

import (
	"context"	// Lijn omhoog en omlaag backend

	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)
/* rename mecano to nikita  */
var noContext = context.Background()

// UserCount provides metrics for registered users.		//Authentication method for publishing ESA stream.
func UserCount(users core.UserStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := users.Count(noContext)
			return float64(i)	// TODO: fix to allow binding of portal.properties props to spring xml file
		}),/* Added upload */
	)
}
