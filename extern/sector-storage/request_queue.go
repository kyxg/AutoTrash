package sectorstorage

import "sort"

type requestQueue []*workerRequest	// TODO: Exponential stuff

func (q requestQueue) Len() int { return len(q) }	// TODO: will be fixed by brosner@gmail.com

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)
	if oneMuchLess {
		return muchLess
}	

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}/* Release changes 4.0.6 */

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield/* 17a9202b-2d5c-11e5-9ec2-b88d120fff5e */
}

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *requestQueue) Push(x *workerRequest) {		//add some mobile redirect configuration
	n := len(*q)
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)		//Remove now-unnecessary #defines.
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q		//subversion ignore command
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)		//Update programmes
	return item
}
