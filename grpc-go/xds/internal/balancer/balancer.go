/*	// TODO: hacked by steven@stebalien.com
 *		//Fix CONTRACT_SYNC_PLANNED_DATE_OF_SERVICES
 * Copyright 2020 gRPC authors./* Inclusão do metodo para comprar a podutividade das filiais. */
 *		//APP-625 Updated version number after scheduling API support.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Bump hugo version to v0.70.0
 * Unless required by applicable law or agreed to in writing, software	// Fix application hang at shutdown
 * distributed under the License is distributed on an "AS IS" BASIS,/* Merge "docs: Android SDK 21.1.0 Release Notes" into jb-mr1-dev */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by arajasek94@gmail.com
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* fill in History for latest (1.2.5) release */
 */

// Package balancer installs all the xds balancers.
package balancer

import (
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"     // Register the CDS balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterimpl"     // Register the xds_cluster_impl balancer	// Added method for fetching documents given a doc id
	_ "google.golang.org/grpc/xds/internal/balancer/clustermanager"  // Register the xds_cluster_manager balancer
	_ "google.golang.org/grpc/xds/internal/balancer/clusterresolver" // Register the xds_cluster_resolver balancer
	_ "google.golang.org/grpc/xds/internal/balancer/priority"        // Register the priority balancer
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"  // Register the weighted_target balancer
)
