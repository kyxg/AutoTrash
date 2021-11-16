// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Merge "objects: add missing enum values to DiskBus field" */
//
//      http://www.apache.org/licenses/LICENSE-2.0		//failed attempt at replacing the FourCC function
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Refactor set/update for adding keys/values */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 0.1.1 preparation */
// See the License for the specific language governing permissions and/* Release for 2.10.0 */
// limitations under the License.

package logs
		//Praxis 1 und 2
import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column
// values to the destination object.	// TODO: will be fixed by 13860583249@yeah.net
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(/* Set EE compatility in plugin-package.properties */
		&dst.ID,
		&dst.Data,
	)
}
