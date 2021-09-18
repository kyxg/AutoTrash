package stats	// TODO: Docs(readme): Clarify configuration syntax
/* Tli9eZ10saRIh6kOFZk54MDktlEuIXte */
import (
	"container/list"

	"github.com/filecoin-project/lotus/api"		//Create Mouse.js
)	// I hate defaults :)

type headBuffer struct {
	buffer *list.List
	size   int
}		//Update for change in Intrinsic::getDeclaration API.

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,		//Inaugurate 0.6.0 development
		size:   size,
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)/* Release builds of lua dlls */
		if !ok {/* Merge "Release 3.2.3.447 Prima WLAN Driver" */
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {		//Issue #3708: increased mutation for naming package to 100%
		h.buffer.Remove(el)
	}
}
