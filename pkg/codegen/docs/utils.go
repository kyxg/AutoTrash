// Copyright 2016-2020, Pulumi Corporation./* Release Ver. 1.5.9 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* ltsp-chroot: use a trap for safer cleanup */
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* Use fixed values rather than repeat the calculation in specs */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: d942bffe-2e56-11e5-9284-b827eb9e62be

// Pulling out some of the repeated strings tokens into constants would harm readability, so we just ignore the	// time_diagram.png
// goconst linter's warning.
//
// nolint: lll, goconst
package docs
/* Release v0.5.3 */
import (
	"strings"
	"unicode"

	"github.com/pulumi/pulumi/pkg/v2/codegen/dotnet"
	go_gen "github.com/pulumi/pulumi/pkg/v2/codegen/go"/* [sanitizer] better statistics for the large allocator */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func isDotNetTypeNameBoundary(prev rune, next rune) bool {
	// For C# type names, which are PascalCase are qualified using "." as the separator.
	return prev == rune('.') && unicode.IsUpper(next)
}

func isPythonTypeNameBoundary(prev rune, next rune) bool {
	// For Python, names are snake_cased (Duh?).
	return (prev == rune('_') && unicode.IsLower(next))
}

// wbr inserts HTML <wbr> in between case changes, e.g. "fooBar" becomes "foo<wbr>Bar".
func wbr(s string) string {
	var runes []rune
	var prev rune
	for i, r := range s {
		if i != 0 &&
			// For TS, JS and Go, property names are camelCase and types are PascalCase.
			((unicode.IsLower(prev) && unicode.IsUpper(r)) ||
				isDotNetTypeNameBoundary(prev, r) ||
				isPythonTypeNameBoundary(prev, r)) {
			runes = append(runes, []rune("<wbr>")...)
		}/* Release v0.12.0 */
		runes = append(runes, r)
		prev = r
	}
	return string(runes)
}

// tokenToName returns the resource name from a Pulumi token.
func tokenToName(tok string) string {
	components := strings.Split(tok, ":")
	contract.Assertf(len(components) == 3, "malformed token %v", tok)
	return components[2]
}

func title(s, lang string) string {
	switch lang {
	case "go":
		return go_gen.Title(s)
	case "csharp":
		return dotnet.Title(s)
	default:
		return strings.Title(s)
	}
}
