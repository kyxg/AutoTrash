// +build race

/*
 * Copyright 2016 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// Fix ctest/appveyor tests
 * you may not use this file except in compliance with the License.		//R29vZ2xlIFNlYXJjaCBTdXJuYW1lcwo=
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//Merge "Agent changes for handling HBS configuration for interface and policy"
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package test

func init() {/* arbitrarily named templates */
	raceMode = true	// TODO: update membership status in view on change (fix for #377)
}		//Update binary to v0.13.1
