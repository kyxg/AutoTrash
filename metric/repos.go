// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// Merge "[INTERNAL] sap.f.DynamicPageTitle - flex-basis used for content area"
// +build !oss

package metric

import (
	"github.com/drone/drone/core"	// TODO: Know if our units are absolute or relative.

	"github.com/prometheus/client_golang/prometheus"
)/* Merge "update company affiliation for devananda" */

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{/* Merge "[DM] Release fabric node from ZooKeeper when releasing lock" */
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)
		}),
	)		//Added: First 'real' implementation of the ZmqPlayer, currently untested
}
