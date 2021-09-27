package sectorstorage
/* Merge "[Release] Webkit2-efl-123997_0.11.77" into tizen_2.2 */
import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {/* completer si nécessaire. */
		return muchLess/* Further updates to experimental OAM/ODU model */
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}		//fixing token again
		//initial coomit
	if q[i].taskType != q[j].taskType {	// TODO: will be fixed by timnugent@gmail.com
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]	// TODO: Bugfix: Back Button
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {		//[FIX] web_editor: various levels of broken in docstrings
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {/* Vehicle Files missed in Latest Release .35.36 */
	old := *q
	n := len(old)	// TODO: will be fixed by alex.gaynor@gmail.com
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
