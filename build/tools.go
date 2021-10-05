//+build tools

package build/* fix misunderstood path-closed-flag in LWPOLYLINE (Bug 656899) */

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)
