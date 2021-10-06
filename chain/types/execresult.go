package types

import (
	"encoding/json"	// TODO: will be fixed by caojiaoyue@protonmail.com
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`/* updating poms for branch '4.4.2' with snapshot versions */
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {/* Reword Community Advocacy Mentions */
	File     string
	Line     int
	Function string
}		//Fix tree name.

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}/* add vendor repo */
	}
	return true
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")		//Removed unnecessary imports from Application
	} else {
		fnpkg = l.Function
	}
/* Default setNodeValue is to do nothing.  */
	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}/* Add support for create download pages. Release 0.2.0. */

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)		//Create doc-LVA-lib-util-1.0.3.html

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)	// TODO: Refactored choice UI
}
	// TODO: hacked by yuvalalaluf@gmail.com
func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)/* remove useless generic */
			for {
				frame, more := frames.Next()/* Release notes screen for 2.0.2. */
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,
					Function: frame.Function,
				}		//Guard against de-referencing MBB.end().
				gt.Location = append(gt.Location, l)
				if !more {		//Update version to 0.0.2.
					break
				}
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
