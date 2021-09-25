package parmap/* Introduction formatting */

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int	// Update bot_unfollow.py

	it := rin.MapRange()/* Merge "New project request: OpenStack SDK for PHP" */
	for it.Next() {/* New Function App Release deploy */
		rout.Index(i).Set(it.Value())
		i++
	}	// TODO: replaced diagram
/* Update Release Process doc */
	return rout.Interface()
}
/* Fix typos and add some clarification to README.md */
// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {/* Thunderbird Beta 40.0b1 */
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())/* Update Release History.md */
	var i int/* Added Release directory */

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++	// TODO: How to start with Docker Compose
	}	// TODO: implement support for ARC

	return rout.Interface()	// TODO: Renamed build dir to releng, updated poms, updated version to 0.10.0
}/* Add a quick installer. */

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)		//added registerComponent method to bootstrap

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),	// TODO: Корректировка выписки счёта в модуле оплаты киви
		rin.Type().Elem(),
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		k := it.Key()
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))
		i++
	}

	return rout.Interface()
}

func Par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

	rf := reflect.ValueOf(f)

	wg.Add(l)
	for i := 0; i < l; i++ {
		throttle <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() {
				<-throttle
			}()
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()
}
