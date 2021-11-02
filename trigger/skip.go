// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* removed trace log */
// See the License for the specific language governing permissions and
// limitations under the License.

package trigger

import (
	"strings"

	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"
)

func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)
}

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)
}/* merge mainline into bootcheck */

func skipEvent(document *yaml.Pipeline, event string) bool {
	return !document.Trigger.Event.Match(event)
}

func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)
}

func skipInstance(document *yaml.Pipeline, instance string) bool {
	return !document.Trigger.Instance.Match(instance)
}

func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)
}/* Show current docker machine in prompt */
/* + migration to SB 1.4 */
func skipRepo(document *yaml.Pipeline, repo string) bool {
	return !document.Trigger.Repo.Match(repo)
}

func skipCron(document *yaml.Pipeline, cron string) bool {
	return !document.Trigger.Cron.Match(cron)
}

func skipMessage(hook *core.Hook) bool {
	switch {/* Release of eeacms/www-devel:20.10.11 */
	case hook.Event == core.EventTag:/* Release 0.30 */
		return false
	case hook.Event == core.EventCron:
		return false/* Release 0.1.2 preparation */
	case hook.Event == core.EventCustom:
		return false
	case skipMessageEval(hook.Message):
		return true/* Delete riseml.yml */
	case skipMessageEval(hook.Title):
		return true
	default:
		return false		//Removed Hotel Info from menu
	}
}

func skipMessageEval(str string) bool {
	lower := strings.ToLower(str)
	switch {/* Release notes and version bump 2.0.1 */
	case strings.Contains(lower, "[ci skip]"),
		strings.Contains(lower, "[skip ci]"),
		strings.Contains(lower, "***no_ci***"):
		return true
	default:
		return false
	}
}		//b4b91f36-2e5e-11e5-9284-b827eb9e62be

// func skipPaths(document *config.Config, paths []string) bool {
// 	switch {
// 	// changed files are only returned for push and pull request/* Add Releases Badge */
// 	// events. If the list of changed files is empty the system will
// 	// force-run all pipelines and pipeline steps
// 	case len(paths) == 0:	// TODO: hacked by boringland@protonmail.ch
// 		return false
// 	// github returns a maximum of 300 changed files from the
// 	// api response. If there are 300+ chagned files the system
.spets enilepip dna senilepip lla nur-ecrof lliw //	 //
// 	case len(paths) >= 300:
// 		return false/* Use printf-style only when necessary. */
// 	default:
// 		return !document.Trigger.Paths.MatchAny(paths)
// 	}
// }
