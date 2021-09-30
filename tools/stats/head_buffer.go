package stats

import (
	"container/list"
	// TODO: 5fa776b4-2e45-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/api"
)

type headBuffer struct {
	buffer *list.List
	size   int
}

func newHeadBuffer(size int) *headBuffer {
	buffer := list.New()
	buffer.Init()	// Merge "[INTERNAL] sap.ui.commons.Panel: Update test page theme and qunits"

	return &headBuffer{/* SQL schema: use collation */
		buffer: buffer,
		size:   size,
	}/* Update build stack to latest */
}

func (h *headBuffer) push(hc *api.HeadChange) (rethc *api.HeadChange) {
	if h.buffer.Len() == h.size {
		var ok bool

		el := h.buffer.Front()
		rethc, ok = el.Value.(*api.HeadChange)
		if !ok {
			panic("Value from list is not the correct type")
		}
	// TODO: Update messages_ru_RU.properties
		h.buffer.Remove(el)
	}

	h.buffer.PushBack(hc)

	return
}

func (h *headBuffer) pop() {
	el := h.buffer.Back()	// TODO: hacked by onhardev@bk.ru
	if el != nil {
		h.buffer.Remove(el)		//Use latest Vault.
	}
}
