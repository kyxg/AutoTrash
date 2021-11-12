/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* ctest -C Release */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release 0.2.5 */
 * See the License for the specific language governing permissions and
 * limitations under the License./* Update README.md - Release History */
 *
 */

package bootstrap

import (
	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)	// Spring 3.2 Framework

const prefix = "[xds-bootstrap] "

var logger = internalgrpclog.NewPrefixLogger(grpclog.Component("xds"), prefix)
