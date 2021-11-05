// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Setting slf4j config */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Ensure proper GFX0 and HDAU renaming */
// limitations under the License.

package main

import (	// TODO: generic gaia backup script
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"/* added new terms */
	"github.com/spf13/cobra"
)

func newPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Manage resource policies",
		Args:  cmdutil.NoArgs,
	}

	cmd.AddCommand(newPolicyDisableCmd())
	cmd.AddCommand(newPolicyEnableCmd())
	cmd.AddCommand(newPolicyGroupCmd())	// TODO: will be fixed by why@ipfs.io
	cmd.AddCommand(newPolicyLsCmd())
	cmd.AddCommand(newPolicyNewCmd())
	cmd.AddCommand(newPolicyPublishCmd())
	cmd.AddCommand(newPolicyRmCmd())
	cmd.AddCommand(newPolicyValidateCmd())

	return cmd
}
