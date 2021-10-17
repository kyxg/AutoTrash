package types	// Delete tank.jpg

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"	// TODO: will be fixed by arajasek94@gmail.com
	"time"
)
/* RealtimeIndexTask: If a Throwable was thrown it is not a normalExit */
type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace/* Released version 0.9.0. */
		//nagios: get dn conf from contexts
	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`	// TODO: will be fixed by boringland@protonmail.ch
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`
/* Release ver 0.2.0 */
	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int	// Add buttons to get the app in the README.md
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{/* e6Mv7DDA5zwJ8vlJekCl6b4almjg6RLg */
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}		//Grammar in read-me.
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true
}/* Correction for terp file processing when init and update xml are empty. */
func (l Loc) String() string {
	file := strings.Split(l.File, "/")
/* 6ed8e180-2e64-11e5-9284-b827eb9e62be */
	fn := strings.Split(l.Function, "/")
	var fnpkg string	// TODO: Update MarqueeBranding component.
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}
	// Some refractoring and documentation
	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}		//Updated  URL to devDependency badge in README

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,
					Function: frame.Function,
				}
				gt.Location = append(gt.Location, l)
				if !more {
					break
				}
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
