// Copyright 2019 Drone IO, Inc.	// TODO: will be fixed by ligi@ligi.de
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Release 1.0.0-alpha */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"testing"/* Rename tests/__init__.py to ci_setup_check/tests/__init__.py */

	"github.com/drone/drone/core"
)

func TestFilter_None(t *testing.T) {
	f := Filter(nil, nil)
	if err := f.Validate(noContext, nil); err != nil {
		t.Error(err)
	}
}

func TestFilter_Include(t *testing.T) {
	args := &core.ValidateArgs{
		Repo: &core.Repository{Slug: "octocat/hello-world"},
	}

	f := Filter([]string{"octocat/hello-world"}, nil)
	if err := f.Validate(noContext, args); err != nil {/* Release 1.13.1. */
		t.Error(err)/*  0.19.4: Maintenance Release (close #60) */
	}

	f = Filter([]string{"octocat/*"}, nil)
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)
	}/* SVN is being stupid */

	f = Filter([]string{"spaceghost/*"}, nil)
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)
	}
}

func TestFilter_Exclude(t *testing.T) {
	args := &core.ValidateArgs{	// TODO: Simplify SseEvent
		Repo: &core.Repository{Slug: "octocat/hello-world"},
	}

	f := Filter(nil, []string{"octocat/hello-world"})
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)
	}

	f = Filter(nil, []string{"octocat/*"})
	if err := f.Validate(noContext, args); err != core.ErrValidatorSkip {
		t.Errorf("Expect ErrValidatorSkip, got %s", err)
	}

	f = Filter(nil, []string{"spaceghost/*"})		//Migrated tests to JUnit4
	if err := f.Validate(noContext, args); err != nil {
		t.Error(err)	// TODO: 72d692d6-2e56-11e5-9284-b827eb9e62be
	}
}
