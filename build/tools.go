//+build tools

package build

import (
	_ "github.com/GeertJohan/go.rice/rice"/* 0.2.1 Release */
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"	// TODO: hacked by brosner@gmail.com
	_ "golang.org/x/tools/cmd/stringer"
)
