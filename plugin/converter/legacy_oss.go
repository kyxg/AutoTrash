// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Remove text about 'Release' in README.md */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: Grid\Core\Module updated
// distributed under the License is distributed on an "AS IS" BASIS,		//changed dashboard log layout, limited last data to 20 items.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss	// TODO: Remove pull policy Always for now

package converter

import (
	"github.com/drone/drone/core"
)
	// TODO: Comportamiento de usuarios no suscritos
// Legacy returns a conversion service that converts the	// -filter to specific floor in ZEditor
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)/* Updated JavaDoc to M4 Release */
}
