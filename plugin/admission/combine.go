// Copyright 2019 Drone IO, Inc.		//Fix deprecation warnings. (also covert tabs to spaces).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Merge "msm: Support both forms of cache dumping" into msm-3.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Remove extra section for Release 2.1.0 from ChangeLog */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by alan.shaw@protocol.ai

package admission/* Delete createPSRelease.sh */

import (
	"context"	// some more safety patches (thanks to Thibaut!)

	"github.com/drone/drone/core"
)

// Combine combines admission services.
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}	// TODO: hacked by ac0dem0nk3y@gmail.com
}		//command matches comment.

type combined struct {
	services []core.AdmissionService
}

func (s *combined) Admit(ctx context.Context, user *core.User) error {/* 3.0.0 Windows Releases */
	for _, service := range s.services {
		if err := service.Admit(ctx, user); err != nil {
			return err
		}
	}
	return nil
}
