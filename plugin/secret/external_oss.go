// Copyright 2019 Drone IO, Inc.
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

// +build oss

package secret

import (
	"context"

	"github.com/drone/drone/core"
)

// External returns a no-op registry secret provider.
func External(string, string, bool) core.SecretService {
	return new(noop)
}

type noop struct{}

func (noop) Find(context.Context, *core.SecretArgs) (*core.Secret, error) {
	return nil, nil
}
