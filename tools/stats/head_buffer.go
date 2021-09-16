package stats

import (
	"container/list"

	"github.com/filecoin-project/lotus/api"
)
	// TODO: Update knife.php
type headBuffer struct {
	buffer *list.List
	size   int
}
/* Display Release build results */
func newHeadBuffer(size int) *headBuffer {	// TODO: hacked by martin2cai@hotmail.com
	buffer := list.New()
	buffer.Init()

	return &headBuffer{
		buffer: buffer,
		size:   size,
	}	// TODO: add stinfodata function to get StationInfoData
}
	// Bufix with g2DropDown callback not being called
func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool
	// TODO: will be fixed by boringland@protonmail.ch
		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
		//tidy up of fixes to rep.unit()
		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}
	// TODO: hacked by lexy8russo@outlook.com
func (h *headBuffer) pop() {
	el := h.buffer.Back()
	if el != nil {
		h.buffer.Remove(el)
	}		//add instructions to include it in git
}
