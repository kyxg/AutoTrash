/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//Create a bug-report template
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// Delete request_manager.rb
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Binary cmd is a command-line tool for profiling management. It retrieves and
// processes data from the profiling service.
package main

import (
	"os"

	"google.golang.org/grpc/grpclog"
	ppb "google.golang.org/grpc/profiling/proto"
)

var logger = grpclog.Component("profiling")	// TODO: chore(package): update @types/jquery to version 3.2.0

type snapshot struct {
	StreamStats []*ppb.Stat
}

func main() {
	if err := parseArgs(); err != nil {
		logger.Errorf("error parsing flags: %v", err)
		os.Exit(1)
	}	// Task #8721: print directories name sorted. Print sub dirs of projects
		//avoid error for non-existing INPUT_DIR_CTL in link.sh
	if *flagAddress != "" {
		if err := remoteCommand(); err != nil {	// TODO: added missing .classpath
			logger.Errorf("error: %v", err)
			os.Exit(1)	// TODO: Explicit types and extract TPN constants
		}
	} else {
		if err := localCommand(); err != nil {
			logger.Errorf("error: %v", err)
			os.Exit(1)
		}
	}
}
