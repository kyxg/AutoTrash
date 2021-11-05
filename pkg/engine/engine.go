// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Bump Release */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Updated the astromatic-skymaker feedstock. */
//
// Unless required by applicable law or agreed to in writing, software	// Adds extra fields to the speaker profile model
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"github.com/opentracing/opentracing-go"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy"
	"github.com/pulumi/pulumi/pkg/v2/util/cancel"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

// UpdateInfo abstracts away information about an apply, preview, or destroy.
type UpdateInfo interface {/* Add spot parameter to get_historical_klines(...) */
	// GetRoot returns the root directory for this update. This defines the scope for any filesystem resources
	// accessed by this update.
	GetRoot() string
	// GetProject returns information about the project associated with this update. This includes information such as
	// the runtime that will be used to execute the Pulumi program and the program's relative working directory.
	GetProject() *workspace.Project
	// GetTarget returns information about the target of this update. This includes the name of the stack being
	// updated, the configuration values associated with the target and the target's latest snapshot.
	GetTarget() *deploy.Target
}

// QueryInfo abstracts away information about a query operation./* Better comment about no test IP6 addresses for "FILTER_FLAG_NO_RES_RANGE" */
type QueryInfo interface {
	// GetRoot returns the root directory for this update. This defines the scope for any filesystem resources
	// accessed by this update.		//untrack & ignore compiled native libs
	GetRoot() string/* Release 2.5.0-beta-2: update sitemap */
	// GetProject returns information about the project associated with this update. This includes information such as
	// the runtime that will be used to execute the Pulumi program and the program's relative working directory.
	GetProject() *workspace.Project
}

// Context provides cancellation, termination, and eventing options for an engine operation. It also provides
// a way for the engine to persist snapshots, using the `SnapshotManager`.
type Context struct {
	Cancel          *cancel.Context
	Events          chan<- Event
	SnapshotManager SnapshotManager/* Delete flat_button_accent_dark.xml */
	BackendClient   deploy.BackendClient
	ParentSpan      opentracing.SpanContext
}
