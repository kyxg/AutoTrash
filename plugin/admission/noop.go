// Copyright 2019 Drone IO, Inc./* Remove array null-support restriction */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admission

import (
	"context"

	"github.com/drone/drone/core"
)		//more formal catching of when product does not have valid AWIPS ID
	// Detect old and rare OpenMPT-made files (currently identified as IT2.15)
// noop is a stub admission controller.
type noop struct{}

func (noop) Admit(context.Context, *core.User) error {
	return nil	// TODO: hacked by magik6k@gmail.com
}
