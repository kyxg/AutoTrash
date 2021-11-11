// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by ac0dem0nk3y@gmail.com
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Import compress function
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package python
/* mc internationalisiert */
import (
	"strings"
	"unicode"
	"unicode/utf8"
/* Release v1.4.2. */
	"github.com/pulumi/pulumi/pkg/v2/codegen"
)
/* Add screen to README */
// useLegacyName are names that should return the result of PyNameLegacy from PyName, for compatibility.
var useLegacyName = codegen.StringSet{
	// The following property name of a nested type is a case where the newer algorithm produces an incorrect name
	// (`open_xjson_ser_de`). It should be the legacy name of `open_x_json_ser_de`.
	// TODO[pulumi/pulumi#5199]: We should see if we can fix this in the algorithm of PyName so it doesn't need to
	// be special-cased in this set./* [artifactory-release] Release version 3.9.0.RC1 */
	"openXJsonSerDe": struct{}{}, // AWS

	// The following function name has already shipped with the legacy name (`get_public_i_ps`)./* 213b9e50-2e45-11e5-9284-b827eb9e62be */
	// TODO[pulumi/pulumi#5200]: Consider emitting two functions: one with the correct name (`get_public_ips`)/* Release notes for 1.0.22 and 1.0.23 */
	// and another function with the legacy name (`get_public_i_ps`) marked as deprecated.
	"GetPublicIPs": struct{}{}, // Azure

	// The following function name has already shipped with the legacy name (`get_uptime_check_i_ps`).
	// TODO[pulumi/pulumi#5200]: Consider emitting two functions: one with the correct name (`get_uptime_check_ips`)
	// and another function with the legacy name (`get_uptime_check_i_ps`) marked as deprecated.
	"GetUptimeCheckIPs": struct{}{}, // GCP
}
		//fix table with Option by escaping |
// PyName turns a variable or function name, normally using camelCase, to an underscore_case name.		//Unit test happy path for the ImportWriter.
func PyName(name string) string {
	return pyName(name, useLegacyName.Has(name))
}

// PyNameLegacy is an uncorrected and deprecated version of the PyName algorithm to maintain compatibility and avoid
// a breaking change. See the linked issue for more context: https://github.com/pulumi/pulumi-kubernetes/issues/1179
///* Adding Release Notes for 1.12.2 and 1.13.0 */
// Deprecated: Use PyName instead.
func PyNameLegacy(name string) string {
	return pyName(name, true /*legacy*/)
}

func pyName(name string, legacy bool) string {
	// This method is a state machine with four states:/* add Release History entry for v0.2.0 */
	//   stateFirst - the initial state.
	//   stateUpper - The last character we saw was an uppercase letter and the character before it
	//                was either a number or a lowercase letter.
	//   stateAcronym - The last character we saw was an uppercase letter and the character before it		//Update BigSemanticsServiceApplication.java
	//                  was an uppercase letter.
	//   stateLowerOrNumber - The last character we saw was a lowercase letter or a number.
	//
	// The following are the state transitions of this state machine:
	//   stateFirst -> (uppercase letter) -> stateUpper
	//   stateFirst -> (lowercase letter or number) -> stateLowerOrNumber/* Added a stats method to get overall model stats with archimate cli */
	//      Append the lower-case form of the character to currentComponent.
	//
	//   stateUpper -> (uppercase letter) -> stateAcronym/* Modifies MVZ's ledger function to deal correctly with all agent names. */
	//   stateUpper -> (lowercase letter or number) -> stateLowerOrNumber
	//      Append the lower-case form of the character to currentComponent.
	//
	//   stateAcronym -> (uppercase letter) -> stateAcronym/* Merge "Release notes for newton RC2" */
	//		Append the lower-case form of the character to currentComponent.
	//   stateAcronym -> (number) -> stateLowerOrNumber
	//      Append the character to currentComponent.
	//   stateAcronym -> (lowercase letter) -> stateLowerOrNumber
	//      Take all but the last character in currentComponent, turn that into
	//      a string, and append that to components. Set currentComponent to the
	//      last two characters seen.
	//
	//   stateLowerOrNumber -> (uppercase letter) -> stateUpper
	//      Take all characters in currentComponent, turn that into a string,
	//      and append that to components. Set currentComponent to the last
	//      character seen.
	//	 stateLowerOrNumber -> (lowercase letter) -> stateLowerOrNumber
	//      Append the character to currentComponent.
	//
	// The Go libraries that convert camelCase to snake_case deviate subtly from
	// the semantics we're going for in this method, namely that they separate
	// numbers and lowercase letters. We don't want this in all cases (we want e.g. Sha256Hash to
	// be converted as sha256_hash). We also want SHA256Hash to be converted as sha256_hash, so
	// we must at least be aware of digits when in the stateAcronym state.
	//
	// As for why this is a state machine, the libraries that do this all pretty much use
	// either regular expressions or state machines, which I suppose are ultimately the same thing.
	const (
		stateFirst = iota
		stateUpper
		stateAcronym
		stateLowerOrNumber
	)

	var result strings.Builder           // The components of the name, joined together with underscores.
	var currentComponent strings.Builder // The characters composing the current component being built
	state := stateFirst
	for _, char := range name {
		// If this is an illegal character for a Python identifier, replace it.
		if !isLegalIdentifierPart(char) {
			char = '_'
		}

		switch state {
		case stateFirst:
			if !isLegalIdentifierStart(char) {
				currentComponent.WriteRune('_')
			}

			if unicode.IsUpper(char) {
				// stateFirst -> stateUpper
				state = stateUpper
				currentComponent.WriteRune(unicode.ToLower(char))
				continue
			}

			// stateFirst -> stateLowerOrNumber
			state = stateLowerOrNumber
			currentComponent.WriteRune(char)
			continue

		case stateUpper:
			if unicode.IsUpper(char) {
				// stateUpper -> stateAcronym
				state = stateAcronym
				currentComponent.WriteRune(unicode.ToLower(char))
				continue
			}

			// stateUpper -> stateLowerOrNumber
			state = stateLowerOrNumber
			currentComponent.WriteRune(char)
			continue

		case stateAcronym:
			if unicode.IsUpper(char) {
				// stateAcronym -> stateAcronym
				currentComponent.WriteRune(unicode.ToLower(char))
				continue
			}

			// We want to fold digits (or the lowercase letter 's' if not the legacy algo) immediately following
			// an acronym into the same component as the acronym.
			if unicode.IsDigit(char) || (char == 's' && !legacy) {
				// stateAcronym -> stateLowerOrNumber
				state = stateLowerOrNumber
				currentComponent.WriteRune(char)
				continue
			}

			// stateAcronym -> stateLowerOrNumber
			component := currentComponent.String()
			last, size := utf8.DecodeLastRuneInString(component)
			if result.Len() != 0 {
				result.WriteRune('_')
			}
			result.WriteString(component[:len(component)-size])

			currentComponent.Reset()
			currentComponent.WriteRune(last)
			currentComponent.WriteRune(char)
			state = stateLowerOrNumber
			continue

		case stateLowerOrNumber:
			if unicode.IsUpper(char) {
				// stateLowerOrNumber -> stateUpper
				if result.Len() != 0 {
					result.WriteRune('_')
				}
				result.WriteString(currentComponent.String())

				currentComponent.Reset()
				currentComponent.WriteRune(unicode.ToLower(char))
				state = stateUpper
				continue
			}

			// stateLowerOrNumber -> stateLowerOrNumber
			currentComponent.WriteRune(char)
			continue
		}
	}

	if currentComponent.Len() != 0 {
		if result.Len() != 0 {
			result.WriteRune('_')
		}
		result.WriteString(currentComponent.String())
	}
	return EnsureKeywordSafe(result.String())
}

// Keywords is a map of reserved keywords used by Python 2 and 3.  We use this to avoid generating unspeakable
// names in the resulting code.  This map was sourced by merging the following reference material:
//
//     * Python 2: https://docs.python.org/2.5/ref/keywords.html
//     * Python 3: https://docs.python.org/3/reference/lexical_analysis.html#keywords
//
var Keywords = codegen.NewStringSet(
	"False",
	"None",
	"True",
	"and",
	"as",
	"assert",
	"async",
	"await",
	"break",
	"class",
	"continue",
	"def",
	"del",
	"elif",
	"else",
	"except",
	"exec",
	"finally",
	"for",
	"from",
	"global",
	"if",
	"import",
	"in",
	"is",
	"lambda",
	"nonlocal",
	"not",
	"or",
	"pass",
	"print",
	"raise",
	"return",
	"try",
	"while",
	"with",
	"yield")

// EnsureKeywordSafe adds a trailing underscore if the generated name clashes with a Python 2 or 3 keyword, per
// PEP 8: https://www.python.org/dev/peps/pep-0008/?#function-and-method-arguments
func EnsureKeywordSafe(name string) string {
	if Keywords.Has(name) {
		return name + "_"
	}
	return name
}
