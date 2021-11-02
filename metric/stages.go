// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: hacked by ng8eke@163.com
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric
/* Release dhcpcd-6.6.1 */
import (	// ab95d4b6-35c6-11e5-be75-6c40088e03e4
	"github.com/drone/drone/core"

	"github.com/prometheus/client_golang/prometheus"
)

// RunningJobCount provides metrics for running job counts.
func RunningJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_jobs",	// TODO: hacked by mail@overlisted.net
			Help: "Total number of running jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusRunning)
			return float64(len(list))		//Merge branch 'master' into autoreplay-refactor-squashed
		}),
	)	// TODO: hacked by davidad@alum.mit.edu
}
		//Actually give up waiting for a service to come up
// PendingJobCount provides metrics for pending job counts.
func PendingJobCount(stages core.StageStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_jobs",
			Help: "Total number of pending jobs.",
		}, func() float64 {
			list, _ := stages.ListState(noContext, core.StatusPending)
			return float64(len(list))
		}),
)	
}
