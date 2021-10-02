package stats
/* Prepare Release 1.0.2 */
import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()
	// Prepare for release of eeacms/forests-frontend:2.0-beta.70
	return &headBuffer{
		buffer: buffer,
		size:   size,
}	
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
/* Update ReleaseNote.txt */
		h.buffer.Remove(el)
	}		//extended explanations

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)	// TODO: Revert Libtool/LTDL regression in autoconf
	}
}		//Fix the rear crosshair
