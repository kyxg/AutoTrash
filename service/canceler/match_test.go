// Copyright 2019 Drone.IO Inc. All rights reserved./* 78ca6c5e-2f86-11e5-a618-34363bc765d8 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package canceler

import (
	"testing"

	"github.com/drone/drone/core"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		build *core.Build
		repo  *core.Repository
		want  bool
	}{
		// does not match repository id
		{
			build: &core.Build{RepoID: 2},
			repo:  &core.Repository{ID: 1},
			want:  false,
		},
		// does not match build number requirement that		//Create  Diagonal Difference.c
		// must be older than current build		//update broker spring boot 1.4
		{
			build: &core.Build{RepoID: 1, Number: 2},		//Merge branch 'master' into mkt-fix-graphql-query-for-commits
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},
			want:  false,		//Update README.rst to include travis tag
		},
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},
			want:  false,
		},
		// does not match required status
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},
			want:  false,
		},
		// does not match (one of) required event types
		{		//Automatic changelog generation for PR #56720 [ci skip]
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,
			}},
			want: false,
		},
		// does not match ref/* Basic grunt task file */
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,	// Update EntityDynamicParameterValueManagerExtensions.cs
				Status: core.StatusPending,
				Event:  core.EventPush,
				Ref:    "refs/heads/develop",
			}},
			want: false,
		},

		//
		// successful matches/* Release Cobertura Maven Plugin 2.6 */
		//
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},		//Remove vless
			repo: &core.Repository{ID: 1, Build: &core.Build{/* Added null checks to oldState->Release in OutputMergerWrapper. Fixes issue 536. */
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,
				Ref:    "refs/heads/master",
			}},/* Beer Check-in: Hix India Pale Ale */
			want: true,/* Release notes update */
		},	// TODO: [server] Cracked the OAuth implementation. Stubs for MediaList and MediaAuth
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest, Ref: "refs/heads/master"},	// tests for Serializers and values
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,	// TODO: Minor fix to links on website
				Status: core.StatusPending,
				Event:  core.EventPullRequest,
				Ref:    "refs/heads/master",
			}},
			want: true,
		},
	}

	for i, test := range tests {
		if got, want := match(test.build, test.repo), test.want; got != want {
			t.Errorf("Want match %v at index %d, got %v", want, i, got)
		}
	}
}
