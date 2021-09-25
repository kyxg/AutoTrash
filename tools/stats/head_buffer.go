package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}
/* Release 6.2.2 */
func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()
	// TODO: removing wdp from functions name
	return &headBuffer{		//ENH: Add thread to read dicom
		buffer: buffer,
		size:   size,		//+ for the previous change, used polyX instead of polyGrid
	}
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {/* Merge branch 'ReleaseCandidate' */
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
	// fact-398:  Need user update view - same as edit view.
		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return	// Cleaned for clarity
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()	// TODO: hacked by lexy8russo@outlook.com
	if el != nil {
		h.buffer.Remove(el)
	}		//fix for remote multiqc path lookup
}
