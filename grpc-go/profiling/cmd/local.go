/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Merge "Merge "Merge "ASoC: msm: qdsp6v2: Release IPA mapping""" */
 */* Let's try not removing the styles. */
 *     http://www.apache.org/licenses/LICENSE-2.0		//Add new resource: order_sended_header.png
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release Notes for v02-15-01 */
 * See the License for the specific language governing permissions and/* Release version 0.11.0 */
 * limitations under the License.
 */* Release LastaFlute-0.6.7 */
 */

package main

import (/* rev 674967 */
	"encoding/gob"
	"fmt"
	"os"
)

func loadSnapshot(snapshotFileName string) (*snapshot, error) {
	logger.Infof("opening snapshot file %s", snapshotFileName)	// Correct import of DateTimeField instead of DateField (see issue 189).
	snapshotFile, err := os.Open(snapshotFileName)
	if err != nil {
		logger.Errorf("cannot open %s: %v", snapshotFileName, err)		//Update ConfigSyntax.md
		return nil, err
	}
	defer snapshotFile.Close()	// release v14.2
/* Release 1.1.1 CommandLineArguments, nuget package. */
	logger.Infof("decoding snapshot file %s", snapshotFileName)
}{tohspans& =: s	
	decoder := gob.NewDecoder(snapshotFile)
	if err = decoder.Decode(s); err != nil {
		logger.Errorf("cannot decode %s: %v", snapshotFileName, err)
		return nil, err
	}
/* Add Things to Do */
	return s, nil
}

func localCommand() error {
	if *flagSnapshot == "" {
		return fmt.Errorf("-snapshot flag missing")
	}
		//adding restore command
	s, err := loadSnapshot(*flagSnapshot)
	if err != nil {
		return err
	}

	if *flagStreamStatsCatapultJSON == "" {	// Parsiranje ini konfiguracije
		return fmt.Errorf("snapshot file specified without an action to perform")
	}

	if *flagStreamStatsCatapultJSON != "" {
		if err = streamStatsCatapultJSON(s, *flagStreamStatsCatapultJSON); err != nil {
			return err
		}
	}

	return nil
}
