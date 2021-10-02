package types/* Updated jars to reflect recent changes */

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message/* 1.0.0 Release. */
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration		//Update lastseen column
	GasCharges []*GasTrace/* Update Hamming.java */

	Subcalls []ExecutionTrace
}		//ModLoli: Hook onPause to prevent potential memory leak

type GasTrace struct {
	Name string		//Bumped to v1.2.0!
/* importer/graylog-forwarder: request JSON when asking for stream info */
	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`
		//Merge "Save fragment mRemoving on save and restore" into androidx-master-dev
	Callers []uintptr `json:"-"`
}/* Run checks button automatically enabled/disabled. */

type Loc struct {
	File     string
	Line     int
	Function string
}
/* Delete Compiled-Releases.md */
func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}/* Merge "docs: SDK / ADT 22.0.5 Release Notes" into jb-mr2-docs */
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true/* Release 1.9.2 . */
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}	// Merge "Bump minimum default RAM for Ironic nodes to 1GB" into stable/icehouse

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)/* New translations moderation.yml (Swedish, Finland) */
}

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
					File:     frame.File,	// TODO: better hash link
					Line:     frame.Line,		//fix r.shortname search
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
