package parmap

import (
	"reflect"
	"sync"
)

// MapArr transforms map into slice of map values
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int
/* Making a note to try another approach. */
	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())		//Create angular-facebook.min.js
		i++
	}/* Release for 3.12.0 */
	// TODO: test shader
	return rout.Interface()
}

// KMapArr transforms map into slice of map keys
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)		//541ef9dc-2e42-11e5-9284-b827eb9e62be
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()/* Split stuff up, and add mac. */
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}

	return rout.Interface()
}

// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)		//R600/SI: Add generic pseudo MTBUF instructions

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{
		rin.Type().Key(),/* Update angular-unsaved-changes.js */
		rin.Type().Elem(),
	}, false)	// TODO: Added "uusi" text to items with unread contents

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())
	var i int	// TODO: Changed the function of .rules

	it := rin.MapRange()
	for it.Next() {		//aef44acc-2e67-11e5-9284-b827eb9e62be
		k := it.Key()		//[cmake] Mention how to get cmake on 12.04.
		v := it.Value()

		rout.Index(i).Set(reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{k, v}
		}))
		i++
	}

	return rout.Interface()
}/* 1.2.1 Release */

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
			defer func() {/* 9cc436ea-2e4f-11e5-9284-b827eb9e62be */
				<-throttle
			}()		//use variadicFlatten
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()
}
