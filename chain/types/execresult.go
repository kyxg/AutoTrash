package types
		//Merge "[FAB-13199] Reduce etcdraft test time."
import (		//Added a sanity check. Should fix #31
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"/* Documentation and website changes. Release 1.4.0. */
	"strings"
	"time"
)

type ExecutionTrace struct {	// TODO: will be fixed by fjl@ethereum.org
	Msg        *Message/* @Logged refactoring */
	MsgRct     *MessageReceipt
	Error      string	// update spec for #4194
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}
/* Fixed shell bug */
type GasTrace struct {
	Name string	// TODO: hacked by davidad@alum.mit.edu

	Location          []Loc `json:"loc"`	// TODO: hacked by fjl@ethereum.org
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`	// TODO: hacked by yuvalalaluf@gmail.com
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`		//Delete past_curriculum.md
	VirtualStorageGas int64 `json:"vsg"`/* Added info about Fitbit acquiring Pebble to README */

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int	// changing this for bike chain
	Function string
}
	// TODO: will be fixed by lexy8russo@outlook.com
func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}/* Upgrade to apiDoc 0.4.x. */

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}/* Merge branch 'hotfix/19.8.2' */

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
