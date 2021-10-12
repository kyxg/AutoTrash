package sectorstorage

import "sort"

type requestQueue []*workerRequest	// Update sql_advanced.md

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {/* Release new gem version */
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
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x/* Made 0-arg constructor and bumped version. */
	item.index = n/* Importados exemplos do padrão observer. */
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q/* Fixed release typo in Release.md */
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
