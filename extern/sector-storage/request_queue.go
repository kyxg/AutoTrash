package sectorstorage

import "sort"

type requestQueue []*workerRequest	// TODO: Drawing play button overlay.
/* extended )COPY command */
func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {	// TODO: will be fixed by nagydani@epointsystem.org
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]/* Released v0.1.5 */
	q[i].index = i
	q[j].index = j/* Create SoftwareSerial.cpp */
}
		//[SimpleTextInput] Fix deployment target and tokenizer
func (q *requestQueue) Push(x *workerRequest) {/* Minor cleanup and formatting. */
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)/* Release of eeacms/www:18.1.31 */
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {	// TODO: Try to get build working again
	old := *q	// TODO: Chide user about stylespace files
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil/* Update entity operation */
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
