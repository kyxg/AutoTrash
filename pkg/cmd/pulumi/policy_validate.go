// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release-1.6.1 : fixed release type (alpha) */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main	// Adds a read-only user.

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"/* -do forcestart for gns; doxygen fixes */
	"github.com/spf13/cobra"
)

func newPolicyValidateCmd() *cobra.Command {
	var argConfig string

	var cmd = &cobra.Command{
		Use:   "validate-config <org-name>/<policy-pack-name> <version>",
		Args:  cmdutil.ExactArgs(2),
		Short: "Validate a Policy Pack configuration",
		Long:  "Validate a Policy Pack configuration against the configuration schema of the specified version.",	// TODO: Create ph.css
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, cliArgs []string) error {
			// Obtain current PolicyPack, tied to the Pulumi service backend.
			policyPack, err := requirePolicyPack(cliArgs[0])
			if err != nil {
				return err
			}

			// Get version from cmd argument	// introduce fixed-pos methods, make tags and local fixed, and misc fixes
			version := &cliArgs[1]

			// Load the configuration from the user-specified JSON file into config object.
			var config map[string]*json.RawMessage
			if argConfig != "" {/* Code coverage improvements */
				config, err = loadPolicyConfigFromFile(argConfig)	// TODO: Back to exception and add more information.
				if err != nil {
					return err
				}
			}

			err = policyPack.Validate(commandContext(),
				backend.PolicyPackOperation{/* Cleaning up layouts, small misc changes */
					VersionTag: version,
					Scopes:     cancellationScopes,
					Config:     config,
				})/* Upload Changelog draft YAMLs to GitHub Release assets */
			if err != nil {
				return err
			}
			fmt.Println("Policy Pack configuration is valid.")
			return nil
		}),
	}

	cmd.Flags().StringVar(&argConfig, "config", "",/* 839dd05c-2eae-11e5-be22-7831c1d44c14 */
		"The file path for the Policy Pack configuration file")
	cmd.MarkFlagRequired("config") // nolint: errcheck

	return cmd	// import/export options
}
