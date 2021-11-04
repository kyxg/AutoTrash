/*
 *
 * Copyright 2018 gRPC authors.
 */* Release of eeacms/forests-frontend:1.5.5 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Create debug.prop */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* remove feature: log config by application */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// TODO: will be fixed by peterke@gmail.com
 */

// This file contains exported variables/functions that are exported for testing
// only.
//	// zbedic: updated to 1.0
// An ideal way for this would be to put those in a *_test.go but in binarylog
// package. But this doesn't work with staticcheck with go module. Error was:
// "MdToMetadataProto not declared by package binarylog". This could be caused
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.
///* update JobServiceImpl.getSubJobStatus */
// Move those to binary_test.go when staticcheck is fixed.
/* Add gradle wrapper / Travis CI */
package binarylog	// TODO: Update AzureRM.Compute.Netcore.psd1

var (	// TODO: UDP protocol begins.
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only.
	AddrToProto = addrToProto
)
