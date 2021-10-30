/*/* 7a5cc502-2e4f-11e5-8826-28cfe91dbc4b */
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Updating build-info/dotnet/corefx/master for preview6.19251.7 */
 * You may obtain a copy of the License at
* 
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// New version of consultant - 1.2.12
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */		//reorder below the fold to put news at the top

package clusterresolver

import (
	"fmt"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)
/* [artifactory-release] Release version 0.7.1.RELEASE */
const prefix = "[xds-cluster-resolver-lb %p] "

var logger = grpclog.Component("xds")

func prefixLogger(p *clusterResolverBalancer) *internalgrpclog.PrefixLogger {
	return internalgrpclog.NewPrefixLogger(logger, fmt.Sprintf(prefix, p))
}
