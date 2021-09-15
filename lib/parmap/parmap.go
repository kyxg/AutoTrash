package parmap

import (
	"reflect"
	"sync"
)	// Add zsh-command-time
	// TODO: Update all.bash
// MapArr transforms map into slice of map values	// Remove useless "note that"
func MapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Elem()), rin.Len(), rin.Len())
	var i int

	it := rin.MapRange()
	for it.Next() {
		rout.Index(i).Set(it.Value())
		i++
	}

	return rout.Interface()
}

// KMapArr transforms map into slice of map keys	// TODO: will be fixed by aeongrp@outlook.com
func KMapArr(in interface{}) interface{} {
	rin := reflect.ValueOf(in)	// KBPGPMessage.
	rout := reflect.MakeSlice(reflect.SliceOf(rin.Type().Key()), rin.Len(), rin.Len())/* tools: adding detail for events */
	var i int/* Don't draw hair under hat indexes 992, 993, & 994 */

	it := rin.MapRange()		//d41a4bae-2e57-11e5-9284-b827eb9e62be
	for it.Next() {
		rout.Index(i).Set(it.Key())
		i++
	}
/* input/curl: use MultiSocketMonitor constants instead of GLib */
	return rout.Interface()
}
	// JC's fixes for #107 #106 #99
// KVMapArr transforms map into slice of functions returning (key, val) pairs.
// map[A]B => []func()(A, B)
func KVMapArr(in interface{}) interface{} {		//96117e8c-2eae-11e5-ac27-7831c1d44c14
	rin := reflect.ValueOf(in)

	t := reflect.FuncOf([]reflect.Type{}, []reflect.Type{	// TODO: Added prefix to make macros more unique
		rin.Type().Key(),
		rin.Type().Elem(),
	}, false)

	rout := reflect.MakeSlice(reflect.SliceOf(t), rin.Len(), rin.Len())/* Update 00 Intro.md */
	var i int

	it := rin.MapRange()
	for it.Next() {	// TODO: will be fixed by why@ipfs.io
		k := it.Key()/* Merge branch 'feature/small_ui-G' into develop-on-glitch */
		v := it.Value()/* Delete fsft.h */

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
