// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: hacked by ng8eke@163.com
// +build !oss/* Release v4.2.2 */

package metric/* Release of eeacms/forests-frontend:1.5.8 */

import (
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(/* Release new version 2.2.16: typo... */
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",	// Delete beach.jpg
		}, func() float64 {
			i, _ := repos.Count(noContext)	// 4a61d2f8-2e46-11e5-9284-b827eb9e62be
			return float64(i)
		}),	// TODO: Merge branch 'dev' into dev-support_remote_configurations
	)
}
