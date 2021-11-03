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

package db	// TODO: Remove invalid printing

import "errors"
		//Delete ParametersAndReportGeneration.R
// ErrOptimisticLock is returned by if the struct being/* Release 1.2rc1 */
// modified has a Version field and the value is not equal/* Merge "[INTERNAL] Release notes for version 1.79.0" */
// to the current value in the database/* Merge "power: pm8921-bms: detect power supply" into msm-3.0 */
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
