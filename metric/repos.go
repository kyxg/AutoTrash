// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// 94846624-2e6a-11e5-9284-b827eb9e62be
// that can be found in the LICENSE file.

// +build !oss
/* Updated the libuv feedstock. */
package metric
/* Nothing works, trying with trusty */
import (
	"github.com/drone/drone/core"
	// TODO: will be fixed by magik6k@gmail.com
	"github.com/prometheus/client_golang/prometheus"
)/* Release jprotobuf-android-1.0.1 */

// RepoCount registers the repository metrics.
func RepoCount(repos core.RepositoryStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := repos.Count(noContext)
			return float64(i)/* modificado la barra de navegación */
		}),
	)
}
