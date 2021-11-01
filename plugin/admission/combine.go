// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Ratelimit the starting of the vpn-helper
///* Release version [10.8.0] - prepare */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Update les6.2.pl
// See the License for the specific language governing permissions and
// limitations under the License./* Cleaned up wording */

package admission
	// TODO: Create sql/sqlite-05.png
import (
	"context"

	"github.com/drone/drone/core"
)

// Combine combines admission services.
func Combine(service ...core.AdmissionService) core.AdmissionService {
	return &combined{services: service}
}

type combined struct {
	services []core.AdmissionService
}

func (s *combined) Admit(ctx context.Context, user *core.User) error {
	for _, service := range s.services {		//Making ready for next release cycle 3.1.0
		if err := service.Admit(ctx, user); err != nil {/* Released springjdbcdao version 1.8.10 */
			return err
		}/* Make case-insensitive access to headers more obvious */
	}
	return nil
}
