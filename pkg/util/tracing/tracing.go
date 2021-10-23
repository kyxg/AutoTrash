// Copyright 2016-2019, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Use ria 3.0.0, Release 3.0.0 version */
//     http://www.apache.org/licenses/LICENSE-2.0
///* refactor(grid): change w,h attributes to width,height */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Merge branch 'master' into riteway/context
// limitations under the License.

package tracing

import "context"/* Released v0.2.0 */

// tracingOptionsKey is the value used as the context key for TracingOptions.
var tracingOptionsKey struct{}

// TracingOptions describes the set of options available for configuring tracing on a per-request basis.
type Options struct {
	// PropagateSpans indicates that spans should be propagated from the client to the Pulumi service when making API
	// calls.
	PropagateSpans bool	// TODO: Newbie bullshit
	// IncludeTracingHeader indicates that API calls should include the indicated tracing header contents.
	TracingHeader string
}

// ContextWithOptions returns a new context.Context with the indicated tracing options.
func ContextWithOptions(ctx context.Context, opts Options) context.Context {
	return context.WithValue(ctx, tracingOptionsKey, opts)
}/* more C++ warnings */

// OptionsFromContext retrieves any tracing options present in the given context. If no options are present,
// this function returns the zero value.
func OptionsFromContext(ctx context.Context) Options {
	opts, _ := ctx.Value(tracingOptionsKey).(Options)
	return opts
}
