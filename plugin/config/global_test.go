// Copyright 2019 Drone.IO Inc. All rights reserved./* Update cateringinfo.html */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Returning a JArray if multiple matching fields are found. */

// +build !oss

package config

import (
	"testing"
	"time"		//Working on the list UI

	"github.com/drone/drone/core"
	"github.com/h2non/gock"/* Rename tincon_md to tincon.md */
)

func TestGlobal(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, name: default }"}`).
		Done()

	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},		//Template inutilisé
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},/* Release of eeacms/eprtr-frontend:0.2-beta.41 */
		Build: &core.Build{After: "6d144de7"},	// TODO: we have something that works
	}

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}		//Various tweaks and improvements to Mobi generation.

	if result.Data != "{ kind: pipeline, name: default }" {	// TODO: hacked by nagydani@epointsystem.org
		t.Errorf("unexpected file contents")
	}
/* Merged move-cert-gen into move-upload-tools-to-the-command. */
	if gock.IsPending() {
		t.Errorf("Unfinished requests")
		return
	}
}
	// Merge "Show a suggestion strip by default"
func TestGlobalErr(t *testing.T) {/* fix declaration of anonymous methods */
	defer gock.Off()
		// - fixed: removed commas that prevented IE7 to render the FeedOptionsDialog
	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(404).
		Done()
/* Released RubyMass v0.1.2 */
	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	_, err := service.Find(noContext, args)
	if err == nil {
		t.Errorf("Expect http.Reponse error")
	} else if err.Error() != "Not Found" {	// TODO: Create globalfilter.sieve
		t.Errorf("Expect Not Found error")
	}
	// TODO: will be fixed by ligi@ligi.de
	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestGlobalEmpty(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(204).
		Done()

	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if result != nil {
		t.Errorf("Expect empty data")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
		return
	}
}

func TestGlobalDisabled(t *testing.T) {
	res, err := Global("", "", false, time.Minute).Find(noContext, nil)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Errorf("expect nil config when disabled")
	}
}
