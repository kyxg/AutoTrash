package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List/* Added russian translation */
	size   int
}

func newHeadBuffer(size int) *headBuffer {/* {v0.2.0} [Children's Day Release] FPS Added. */
	buffer := list.New()
	buffer.Init()	// fix metamodel tests

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}
}	// Make binary-dist do nothing in doc/Makefile, for now

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool
/* Release of eeacms/plonesaas:5.2.4-9 */
		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)/* Updated to Latest Release */
		if !ok {	// TODO: hacked by hugomrdias@gmail.com
			panic("Value from list is not the correct type")
		}

		h.buffer.Remove(el)
}	

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()/* Rename L_SITS_NNED.java to SITS_NNED.java */
	if el != nil {
		h.buffer.Remove(el)
	}/* Begin of Feature API implementation: getFeature, getAllFeatures */
}
