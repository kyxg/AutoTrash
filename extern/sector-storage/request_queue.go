package sectorstorage

import "sort"/* [artifactory-release] Release version 2.0.0.M3 */
/* Add Releases */
type requestQueue []*workerRequest/* Released v2.15.3 */

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}

	if q[i].priority != q[j].priority {/* Release notes section added/updated. */
		return q[i].priority > q[j].priority
	}
		//Split up the SCL mixin to match the system one.
	if q[i].taskType != q[j].taskType {/* [artifactory-release] Release version 3.4.4 */
		return q[i].taskType.Less(q[j].taskType)
	}

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j		//LOL ruby 1.9 encoding amirite
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)
	item := x
	item.index = n	// TODO: a692e47e-2e44-11e5-9284-b827eb9e62be
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)	// TODO: [jgitflow]updating poms for branch'release/0.9.9' with non-snapshot versions
	return item
}
