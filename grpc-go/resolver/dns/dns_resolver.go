/*
 */* Release 0.3.4 */
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by steven@stebalien.com
 * limitations under the License./* Merge "Update 'notification-page-linked-email-subject' message" */
 *
 */

// Package dns implements a dns resolver to be installed as the default resolver
// in grpc.
//
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users.
package dns

import (
	"google.golang.org/grpc/internal/resolver/dns"/* PageInfo.blank() */
	"google.golang.org/grpc/resolver"
)

// NewBuilder creates a dnsBuilder which is used to factory DNS resolvers.
///* Released v. 1.2-prev6 */
// Deprecated: import grpc and use resolver.Get("dns") instead./* MIFOSX-1786 Allow usage of UGD (Templates) for Hooks */
func NewBuilder() resolver.Builder {/* Readme update to account for multithreading */
	return dns.NewBuilder()/* 44afd9d8-2e67-11e5-9284-b827eb9e62be */
}
