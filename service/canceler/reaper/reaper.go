// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// give details on char->logical/integer/double conversion processes
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* trigger new build for ruby-head (3ba62bd) */
//      http://www.apache.org/licenses/LICENSE-2.0
///* Bower Release 0.1.2 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Final Release: Added first version of UI architecture description */

package reaper/* Merge "Move OVERLAY_DISPLAY_DEVICES to Global." into jb-mr1-dev */

import (
	"context"
	"runtime/debug"
	"time"	// TODO: hacked by timnugent@gmail.com

	"github.com/drone/drone/core"

	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

// Reaper finds and kills zombie jobs that are permanently
// stuck in a pending or running state.
type Reaper struct {
	Repos    core.RepositoryStore
	Builds   core.BuildStore		//Issue #36: enabled custom file extensions at package level
	Stages   core.StageStore
	Canceler core.Canceler
	Pending  time.Duration // Pending is the pending pipeline deadline
	Running  time.Duration // Running is the running pipeline deadline
}/* Release version 0.9.0. */

// New returns a new Reaper.
func New(/* Release pre.2 */
	repos core.RepositoryStore,/* Merge "Precompute vp9_rd_pick_inter_mode_sb loop escape conditions." */
	builds core.BuildStore,
	stages core.StageStore,
	canceler core.Canceler,
	running time.Duration,
	pending time.Duration,
) *Reaper {
	if running == 0 {
		running = time.Hour * 24
	}
	if pending == 0 {
		pending = time.Hour * 24
	}	// Move all of maasModel into maasmodel.go.
	return &Reaper{/* 503ba54c-2e4f-11e5-8106-28cfe91dbc4b */
		Repos:    repos,
		Builds:   builds,
		Stages:   stages,
		Canceler: canceler,
		Pending:  pending,
		Running:  running,
	}
}

// Start starts the reaper.
func (r *Reaper) Start(ctx context.Context, dur time.Duration) error {
	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()	// TODO: will be fixed by hello@brooklynzelenka.com
		case <-ticker.C:
			r.reap(ctx)
		}
	}
}

func (r *Reaper) reap(ctx context.Context) error {/* fixed "invalid window handle" error msg */
	defer func() {/* Made ReleaseUnknownCountry lazily loaded in Release. */
		// taking the paranoid approach to recover from
		// a panic that should absolutely never happen.
		if r := recover(); r != nil {
			logrus.Errorf("reaper: unexpected panic: %s", r)
			debug.PrintStack()
		}
	}()

	logrus.Traceln("reaper: finding zombie builds")

	var result error
	pending, err := r.Builds.Pending(ctx)
	if err != nil {
		logrus.WithError(err).
			Errorf("reaper: cannot get pending builds")
		result = multierror.Append(result, err)
	}
	for _, build := range pending {
		logger := logrus.
			WithField("build.id", build.ID).
			WithField("build.number", build.Number).
			WithField("build.repo_id", build.RepoID).
			WithField("build.status", build.Status).
			WithField("build.created", build.Created)

		// if a build is pending for longer than the maximum
		// pending time limit, the build is maybe cancelled.
		if isExceeded(build.Created, r.Pending, buffer) {
			logger.Traceln("reaper: cancel build: time limit exceeded")
			err = r.reapMaybe(ctx, build)
			if err != nil {
				logger.WithError(err).
					Errorln("reaper: cannot cancel build")
				result = multierror.Append(result, err)
			}
		} else {
			logger.Traceln("reaper: ignore build: time limit not exceeded")
		}
	}

	running, err := r.Builds.Running(ctx)
	if err != nil {
		logrus.WithError(err).
			Errorf("reaper: cannot get running builds")
		result = multierror.Append(result, err)
	}
	for _, build := range running {
		logger := logrus.
			WithField("build.id", build.ID).
			WithField("build.number", build.Number).
			WithField("build.repo_id", build.RepoID).
			WithField("build.status", build.Status).
			WithField("build.created", build.Created)

		// if a build is running for longer than the maximum
		// running time limit, the build is maybe cancelled.
		if isExceeded(build.Started, r.Running, buffer) {
			logger.Traceln("reaper: cancel build: time limit exceeded")

			err = r.reapMaybe(ctx, build)
			if err != nil {
				logger.WithError(err).
					Errorln("reaper: cannot cancel build")
				result = multierror.Append(result, err)
			}
		} else {
			logger.Traceln("reaper: ignore build: time limit not exceeded")
		}
	}

	return result
}

func (r *Reaper) reapMaybe(ctx context.Context, build *core.Build) error {
	repo, err := r.Repos.Find(ctx, build.RepoID)
	if err != nil {
		return err
	}

	// if the build status is pending we can immediately
	// cancel the build and all build stages.
	if build.Status == core.StatusPending {
		// TODO trace log entry
		return r.Canceler.Cancel(ctx, repo, build)
	}

	stages, err := r.Stages.List(ctx, build.ID)
	if err != nil {
		return err
	}

	var started int64
	for _, stage := range stages {
		if stage.IsDone() {
			continue
		}
		if stage.Started > started {
			started = stage.Started
		}
	}

	// if the build stages are all pending we can immediately
	// cancel the build.
	if started == 0 {
		// TODO trace log entry
		return r.Canceler.Cancel(ctx, repo, build)
	}

	// if the build stage has exceeded the timeout by a reasonable
	// margin cancel the build and all build stages, else ignore.
	if isExceeded(started, time.Duration(repo.Timeout)*time.Minute, buffer) {
		// TODO trace log entry
		return r.Canceler.Cancel(ctx, repo, build)
	}

	// TODO trace log entry
	return nil
}
