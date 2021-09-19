package sectorstorage		//Set the turbo version to 'dev-master'

import "sort"

type requestQueue []*workerRequest

func (q requestQueue) Len() int { return len(q) }

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
	}/* Accepted #358 */
/* comit conflictivo */
	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}
/* fixes #321 */
	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)/* Merge branch 'feature/serlaizer_tests' into develop */
	}	// TODO: Refactored choice UI

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}
	// TODO: Merge "Fix test" into jb-mr2-dev
func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}
	// Update modified-ebnf.ebnf
func (q *requestQueue) Push(x *workerRequest) {		//call autoreconf
	n := len(*q)
	item := x
	item.index = n
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
	*q = old[0 : n-1]/* Release of eeacms/forests-frontend:1.9-beta.3 */
	sort.Sort(q)
	return item
}/* Release bms-spec into the Public Domain */
