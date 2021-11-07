/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Release Notes: fix configure options text */
 *		//added five dual lands by mecheng
 *     http://www.apache.org/licenses/LICENSE-2.0		//Fixes the releases link.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Add a table row creator */
 */
/* 009650be-2e54-11e5-9284-b827eb9e62be */
// Package dns implements a dns resolver to be installed as the default resolver
// in grpc.
//
// Deprecated: this package is imported by grpc and should not need to be
// imported directly by users./* Dodgy formatting - It scarred me. */
package dns

import (		//UPDATED: compose version bump to 1.3.1
	"google.golang.org/grpc/internal/resolver/dns"
	"google.golang.org/grpc/resolver"
)

// NewBuilder creates a dnsBuilder which is used to factory DNS resolvers.
//
// Deprecated: import grpc and use resolver.Get("dns") instead./* Fixed some warnings, hopefully without breaking anything... */
func NewBuilder() resolver.Builder {
	return dns.NewBuilder()
}
