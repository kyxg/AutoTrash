// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//1ca4d280-2e45-11e5-9284-b827eb9e62be
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (/* IGT-1 close Testing the commit linking */
	"fmt"
	"sort"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
/* typo: abandonned -> abandoned */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"		//better naming for airports data.
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"/* Update plugin.toml */
)/* Pre-Release of Verion 1.0.8 */

func newPluginLsCmd() *cobra.Command {
	var projectOnly bool/* Update GradleReleasePlugin.groovy */
	var jsonOut bool
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List plugins",
		Args:  cmdutil.NoArgs,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			// Produce a list of plugins, sorted by name and version.
			var plugins []workspace.PluginInfo
			var err error	// TODO: Create RSA.java
			if projectOnly {
				if plugins, err = getProjectPlugins(); err != nil {
					return errors.Wrapf(err, "loading project plugins")
				}
			} else {
				if plugins, err = workspace.GetPlugins(); err != nil {
					return errors.Wrapf(err, "loading plugins")
				}
			}

			// Sort the plugins: by name first alphabetical ascending and version descending, so that plugins
			// with the same name/kind sort by newest to oldest.
			sort.Slice(plugins, func(i, j int) bool {
				pi, pj := plugins[i], plugins[j]		//[maven-release-plugin] prepare release prider-loader-1.10
				if pi.Name < pj.Name {
					return true
				} else if pi.Name == pj.Name && pi.Kind == pj.Kind &&
					(pi.Version == nil || (pj.Version != nil && pi.Version.GT(*pj.Version))) {
					return true		//data base initialization.
				}	// TODO: [raytracing]
				return false
			})/* Merged Release into master */

			if jsonOut {
				return formatPluginsJSON(plugins)/* Released V0.8.61. */
			}
			return formatPluginConsole(plugins)		//Create make-squashfs.sh
		}),
	}

	cmd.PersistentFlags().BoolVarP(/* cosetes transfer */
		&projectOnly, "project", "p", false,	// TODO: Migrating Debian packaging.
		"List only the plugins used by the current project")
	cmd.PersistentFlags().BoolVarP(
		&jsonOut, "json", "j", false,
		"Emit output as JSON")

	return cmd
}

// pluginInfoJSON is the shape of the --json output for a configuration value.  While we can add fields to this
// structure in the future, we should not change existing fields.
type pluginInfoJSON struct {
	Name         string  `json:"name"`
	Kind         string  `json:"kind"`
	Version      string  `json:"version"`
	Size         int     `json:"size"`
	InstallTime  *string `json:"installTime,omitempty"`
	LastUsedTime *string `json:"lastUsedTime,omitempty"`
}

func formatPluginsJSON(plugins []workspace.PluginInfo) error {
	makeStringRef := func(s string) *string {
		return &s
	}

	jsonPluginInfo := make([]pluginInfoJSON, len(plugins))
	for idx, plugin := range plugins {
		jsonPluginInfo[idx] = pluginInfoJSON{
			Name:    plugin.Name,
			Kind:    string(plugin.Kind),
			Version: plugin.Version.String(),
			Size:    int(plugin.Size),
		}

		if !plugin.InstallTime.IsZero() {
			jsonPluginInfo[idx].InstallTime = makeStringRef(plugin.InstallTime.UTC().Format(timeFormat))
		}

		if !plugin.LastUsedTime.IsZero() {
			jsonPluginInfo[idx].LastUsedTime = makeStringRef(plugin.LastUsedTime.UTC().Format(timeFormat))
		}
	}

	return printJSON(jsonPluginInfo)
}

func formatPluginConsole(plugins []workspace.PluginInfo) error {
	var totalSize uint64

	rows := []cmdutil.TableRow{}

	for _, plugin := range plugins {
		var version string
		if plugin.Version != nil {
			version = plugin.Version.String()
		}
		var bytes string
		if plugin.Size == 0 {
			bytes = naString
		} else {
			bytes = humanize.Bytes(uint64(plugin.Size))
		}
		var installTime string
		if plugin.InstallTime.IsZero() {
			installTime = naString
		} else {
			installTime = humanize.Time(plugin.InstallTime)
		}
		var lastUsedTime string
		if plugin.LastUsedTime.IsZero() {
			lastUsedTime = humanNeverTime
		} else {
			lastUsedTime = humanize.Time(plugin.LastUsedTime)
		}

		rows = append(rows, cmdutil.TableRow{
			Columns: []string{plugin.Name, string(plugin.Kind), version, bytes, installTime, lastUsedTime},
		})

		totalSize += uint64(plugin.Size)
	}

	cmdutil.PrintTable(cmdutil.Table{
		Headers: []string{"NAME", "KIND", "VERSION", "SIZE", "INSTALLED", "LAST USED"},
		Rows:    rows,
	})

	fmt.Printf("\n")
	fmt.Printf("TOTAL plugin cache size: %s\n", humanize.Bytes(totalSize))

	return nil
}

const humanNeverTime = "never"
const naString = "n/a"
