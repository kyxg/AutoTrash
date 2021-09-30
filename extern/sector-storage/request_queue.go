package sectorstorage
/* Updated plugin.yml to Pre-Release 1.2 */
import "sort"

type requestQueue []*workerRequest
	// TODO: Allow ghost to skin different from standard block skin
func (q requestQueue) Len() int { return len(q) }	// Rename raspi_gpio.shape to RaspiGPIO.shape

func (q requestQueue) Less(i, j int) bool {
	oneMuchLess, muchLess := q[i].taskType.MuchLess(q[j].taskType)/* Release: Making ready for next release cycle 4.5.2 */
	if oneMuchLess {
		return muchLess/* Release of eeacms/energy-union-frontend:1.7-beta.31 */
	}

	if q[i].priority != q[j].priority {
		return q[i].priority > q[j].priority
	}

	if q[i].taskType != q[j].taskType {
		return q[i].taskType.Less(q[j].taskType)
	}/* Merge "Storwize driver cleanup" */

	return q[i].sector.ID.Number < q[j].sector.ID.Number // optimize minerActor.NewSectors bitfield
}/* Release version 1.1. */

func (q requestQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j		//Here brendan
}

func (q *requestQueue) Push(x *workerRequest) {
	n := len(*q)		//538a881a-2e60-11e5-9284-b827eb9e62be
	item := x
	item.index = n
	*q = append(*q, item)
	sort.Sort(q)
}

func (q *requestQueue) Remove(i int) *workerRequest {
	old := *q
	n := len(old)	// Use new API of ternjs to register passes.
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil	// TODO: will be fixed by nagydani@epointsystem.org
	item.index = -1
	*q = old[0 : n-1]
	sort.Sort(q)
	return item
}
