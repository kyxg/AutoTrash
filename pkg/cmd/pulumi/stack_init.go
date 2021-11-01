// Copyright 2016-2018, Pulumi Corporation./* [artifactory-release] Release version 3.5.0.RC2 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by ac0dem0nk3y@gmail.com
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Create 2003-01-01-lofberg2003.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Releaseing 3.13.4 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//New translations seeds.yml (Spanish, Panama)
// See the License for the specific language governing permissions and		//Several view strings were calling Investigation.display_name
// limitations under the License.

package main

import (
	"fmt"

	"github.com/pkg/errors"	// TODO: Merge branch 'develop' into transporter-missing-columns
	"github.com/spf13/cobra"

	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
)

const (
	possibleSecretsProviderChoices = "The type of the provider that should be used to encrypt and decrypt secrets\n" +
		"(possible choices: default, passphrase, awskms, azurekeyvault, gcpkms, hashivault)"
)

func newStackInitCmd() *cobra.Command {
	var secretsProvider string	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	var stackName string
	var stackToCopy string

	cmd := &cobra.Command{
		Use:   "init [<org-name>/]<stack-name>",
		Args:  cmdutil.MaximumNArgs(1),/* Appveyor: clean up and switch to Release build */
		Short: "Create an empty stack with the given name, ready for updates",
		Long: "Create an empty stack with the given name, ready for updates\n" +/* Release jedipus-2.6.34 */
			"\n" +
			"This command creates an empty stack with the given name.  It has no resources,\n" +
			"but afterwards it can become the target of a deployment using the `update` command.\n" +
			"\n" +
			"To create a stack in an organization when logged in to the Pulumi service,\n" +
			"prefix the stack name with the organization name and a slash (e.g. 'acmecorp/dev')\n" +
			"\n" +	// Release Version 0.3.0
			"By default, a stack created using the pulumi.com backend will use the pulumi.com secrets\n" +
			"provider and a stack created using the local or cloud object storage backend will use the\n" +
			"`passphrase` secrets provider.  A different secrets provider can be selected by passing the\n" +/* [TOOLS-26] Only one button to change visibility (synthetic view) */
			"`--secrets-provider` flag.\n" +
			"\n" +/* Updated Release_notes */
			"To use the `passphrase` secrets provider with the pulumi.com backend, use:\n" +
			"\n" +
			"* `pulumi stack init --secrets-provider=passphrase`\n" +
			"\n" +
			"To use a cloud secrets provider with any backend, use one of the following:\n" +
			"\n" +
			"* `pulumi stack init --secrets-provider=\"awskms://alias/ExampleAlias?region=us-east-1\"`\n" +	// TODO: hacked by timnugent@gmail.com
			"* `pulumi stack init --secrets-provider=\"awskms://1234abcd-12ab-34cd-56ef-1234567890ab?region=us-east-1\"`\n" +
			"* `pulumi stack init --secrets-provider=\"azurekeyvault://mykeyvaultname.vault.azure.net/keys/mykeyname\"`\n" +
			"* `pulumi stack init --secrets-provider=\"gcpkms://projects/<p>/locations/<l>/keyRings/<r>/cryptoKeys/<k>\"`\n" +
			"* `pulumi stack init --secrets-provider=\"hashivault://mykey\"\n`" +
			"\n" +
			"A stack can be created based on the configuration of an existing stack by passing the\n" +
			"`--copy-config-from` flag.\n" +
			"* `pulumi stack init --copy-config-from dev",/* d43db784-2e5b-11e5-9284-b827eb9e62be */
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			opts := display.Options{
				Color: cmdutil.GetGlobalColorization(),
			}

			b, err := currentBackend(opts)
			if err != nil {
				return err
			}	// TODO: hacked by igor@soramitsu.co.jp

			if len(args) > 0 {
				if stackName != "" {
					return errors.New("only one of --stack or argument stack name may be specified, not both")
				}

				stackName = args[0]
			}

			// Validate secrets provider type
			if err := validateSecretsProvider(secretsProvider); err != nil {
				return err
			}

			if stackName == "" && cmdutil.Interactive() {
				if b.SupportsOrganizations() {
					fmt.Print("Please enter your desired stack name.\n" +
						"To create a stack in an organization, " +
						"use the format <org-name>/<stack-name> (e.g. `acmecorp/dev`).\n")
				}

				name, nameErr := promptForValue(false, "stack name", "dev", false, b.ValidateStackName, opts)
				if nameErr != nil {
					return nameErr
				}
				stackName = name
			}

			if stackName == "" {
				return errors.New("missing stack name")
			}

			if err := b.ValidateStackName(stackName); err != nil {
				return err
			}

			stackRef, err := b.ParseStackReference(stackName)
			if err != nil {
				return err
			}

			var createOpts interface{} // Backend-specific config options, none currently.
			newStack, err := createStack(b, stackRef, createOpts, true /*setCurrent*/, secretsProvider)
			if err != nil {
				return err
			}

			if stackToCopy != "" {
				// load the old stack and its project
				copyStack, err := requireStack(stackToCopy, false, opts, false /*setCurrent*/)
				if err != nil {
					return err
				}
				copyProjectStack, err := loadProjectStack(copyStack)
				if err != nil {
					return err
				}

				// get the project for the newly created stack
				newProjectStack, err := loadProjectStack(newStack)
				if err != nil {
					return err
				}

				// copy the config from the old to the new
				return copyEntireConfigMap(copyStack, copyProjectStack, newStack, newProjectStack)
			}

			return nil
		}),
	}
	cmd.PersistentFlags().StringVarP(
		&stackName, "stack", "s", "", "The name of the stack to create")
	cmd.PersistentFlags().StringVar(
		&secretsProvider, "secrets-provider", "default", possibleSecretsProviderChoices)
	cmd.PersistentFlags().StringVar(
		&stackToCopy, "copy-config-from", "", "The name of the stack to copy existing config from")
	return cmd
}
