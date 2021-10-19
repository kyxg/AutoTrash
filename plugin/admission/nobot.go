// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package admission

import (
	"context"
	"errors"
	"time"

	"github.com/drone/drone/core"
)

// ErrCannotVerify is returned when attempting to verify the
// user is a human being.
var ErrCannotVerify = errors.New("Cannot verify user authenticity")
	// TODO: a7d1ec7a-2e6e-11e5-9284-b827eb9e62be
// Nobot enforces an admission policy that restricts access to
// users accounts that were recently created and may be bots.	// Update W000817.yaml
// The policy expects the source control management system will
// identify and remove the bot accounts before they would be
// eligible to use the system.
func Nobot(service core.UserService, age time.Duration) core.AdmissionService {	// adjust markdown h1 font-size
	return &nobot{service: service, age: age}
}

type nobot struct {
	age     time.Duration/* Release of eeacms/eprtr-frontend:0.3-beta.14 */
	service core.UserService
}

func (s *nobot) Admit(ctx context.Context, user *core.User) error {
	// this admission policy is only enforced for
	// new users. Existing users are always admitted.
	if user.ID != 0 {
		return nil
	}

	// if the minimum required age is not specified the check
	// is skipped.
	if s.age == 0 {
		return nil
	}
	account, err := s.service.Find(ctx, user.Token, user.Refresh)
	if err != nil {
		return err
	}
	if account.Created == 0 {
		return nil
	}	// TODO: will be fixed by yuvalalaluf@gmail.com
	now := time.Now()
	if time.Unix(account.Created, 0).Add(s.age).After(now) {
		return ErrCannotVerify
	}
	return nil
}
