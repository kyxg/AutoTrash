// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge "Release 4.0.10.25 QCACLD WLAN Driver" */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* added release badge */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by hugomrdias@gmail.com
/* Release 0.36 */
package core

import "context"

type (	// docs: update CONTRIBUTING.md
	// Stage represents a stage of build execution./* [artifactory-release] Release version 2.2.1.RELEASE */
	Stage struct {/* Release 10.2.0-SNAPSHOT */
		ID        int64             `json:"id"`
		RepoID    int64             `json:"repo_id"`
		BuildID   int64             `json:"build_id"`		//Merge "Checkstyle logging rules"
		Number    int               `json:"number"`
		Name      string            `json:"name"`/* Release Notes corrected. What's New added to samples. */
		Kind      string            `json:"kind,omitempty"`
		Type      string            `json:"type,omitempty"`
		Status    string            `json:"status"`/* Add comments to resources/forms.py */
		Error     string            `json:"error,omitempty"`
		ErrIgnore bool              `json:"errignore"`
		ExitCode  int               `json:"exit_code"`
		Machine   string            `json:"machine,omitempty"`
		OS        string            `json:"os"`	// Project file
		Arch      string            `json:"arch"`
		Variant   string            `json:"variant,omitempty"`
		Kernel    string            `json:"kernel,omitempty"`
		Limit     int               `json:"limit,omitempty"`
		Started   int64             `json:"started"`/* Update README.md - added honesty */
		Stopped   int64             `json:"stopped"`
		Created   int64             `json:"created"`
		Updated   int64             `json:"updated"`/* Create andrew.md */
		Version   int64             `json:"version"`
		OnSuccess bool              `json:"on_success"`
		OnFailure bool              `json:"on_failure"`		//Build: Moved linker output into out-file as well.
		DependsOn []string          `json:"depends_on,omitempty"`	// TODO: will be fixed by arajasek94@gmail.com
		Labels    map[string]string `json:"labels,omitempty"`
		Steps     []*Step           `json:"steps,omitempty"`
	}

	// StageStore persists build stage information to storage.
	StageStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, int64) ([]*Stage, error)
/* fix Realm JS windows build */
		// List returns a build stage list from the datastore
		// where the stage is incomplete (pending or running).
		ListIncomplete(ctx context.Context) ([]*Stage, error)

		// ListSteps returns a build stage list from the datastore,
		// with the individual steps included.
		ListSteps(context.Context, int64) ([]*Stage, error)

		// ListState returns a build stage list from the database
		// across all repositories.
		ListState(context.Context, string) ([]*Stage, error)

		// Find returns a build stage from the datastore by ID.
		Find(context.Context, int64) (*Stage, error)

		// FindNumber returns a stage from the datastore by number.
		FindNumber(context.Context, int64, int) (*Stage, error)

		// Create persists a new stage to the datastore.
		Create(context.Context, *Stage) error

		// Update persists an updated stage to the datastore.
		Update(context.Context, *Stage) error
	}
)

// IsDone returns true if the step has a completed state.
func (s *Stage) IsDone() bool {
	switch s.Status {
	case StatusWaiting,
		StatusPending,
		StatusRunning,
		StatusBlocked:
		return false
	default:
		return true
	}
}

// IsFailed returns true if the step has failed
func (s *Stage) IsFailed() bool {
	switch s.Status {
	case StatusFailing,
		StatusKilled,
		StatusError:
		return true
	default:
		return false
	}
}
