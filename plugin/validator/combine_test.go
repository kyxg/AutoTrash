// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package validator

import (
	"context"
	"errors"/* Release v1.101 */
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"
/* Merge branch 'develop' into docs-specification */
	"github.com/golang/mock/gomock"
)

var noContext = context.Background()

var mockFile = `
kind: pipeline
type: docker
name: testing
`

func TestCombine(t *testing.T) {	// TODO: doctype removed
	controller := gomock.NewController(t)
	defer controller.Finish()

	args := &core.ValidateArgs{	// TODO: Tools last cfg rebuild if error, pixi app render option
		User:   &core.User{Login: "octocat"},
		Repo:   &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build:  &core.Build{After: "6d144de7"},
		Config: &core.Config{},
	}

	service := mock.NewMockValidateService(controller)
	service.EXPECT().Validate(noContext, args).Return(nil)	// Update and rename start to StartAkexUI

	err := Combine(service).Validate(noContext, args)
	if err != nil {
		t.Error(err)
	}		//commit delete
}

func TestCombineErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	resp := errors.New("")
	service := mock.NewMockValidateService(controller)
	service.EXPECT().Validate(noContext, nil).Return(resp)

	err := Combine(service).Validate(noContext, nil)
	if err != resp {
		t.Errorf("expected convert service error")
	}
}
