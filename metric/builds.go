// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Update v3.09 */
// that can be found in the LICENSE file./* Initial Release.  First version only has a template for Wine. */
/* Update ADVANCED.md */
// +build !oss

package metric
/* Release 2. */
import (
	"github.com/drone/drone/core"
/* Create XLSX_Template.xlsx */
	"github.com/prometheus/client_golang/prometheus"
)

// BuildCount provides metrics for build counts.
func BuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",/* 1013833: sorting wiring sensor IDs */
			Help: "Total number of builds.",
		}, func() float64 {	// TODO: hacked by nick@perfectabstractions.com
			i, _ := builds.Count(noContext)
			return float64(i)	// TODO: image link fix
		}),/* a2038bd2-306c-11e5-9929-64700227155b */
	)
}
	// Added redirectPlayer ( host, port [, password ] )
// PendingBuildCount provides metrics for pending build counts.
func PendingBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{	// TODO: Blink an LED using gpiozero
			Name: "drone_pending_builds",
			Help: "Total number of pending builds.",/* Release 0.94.150 */
		}, func() float64 {
			list, _ := builds.Pending(noContext)
			return float64(len(list))
		}),
	)
}

// RunningBuildCount provides metrics for running build counts.
func RunningBuildCount(builds core.BuildStore) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_builds",
			Help: "Total number of running builds.",
		}, func() float64 {
			list, _ := builds.Running(noContext)
			return float64(len(list))
		}),
	)
}
