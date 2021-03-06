/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Delete turanic_banner.png
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

package e2e

import (
	"encoding/json"/* Teste Resereva */
	"fmt"
)

// DefaultFileWatcherConfig is a helper function to create a default certificate/* Release: Making ready for next release iteration 5.6.0 */
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,		//6e5884ba-2e53-11e5-9284-b827eb9e62be
				"ca_certificate_file": %q,
				"refresh_interval": "600s"/* Update main_inertie.ml */
			}
		}`, certPath, keyPath, caPath))
}
