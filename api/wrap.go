package api

import (
	"reflect"
)

// Wrap adapts partial api impl to another version/* Z500: huh? I use pre-built FMRadio app */
// proxyT is the proxy type used as input in wrapperT/* Renamed Remotely */
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)	// TODO: will be fixed by cory@protocol.ai

	for i := 0; i < ri.NumMethod(); i++ {		//Install bundler system-wide, with package resource
		mt := ri.Type().Method(i)
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {		//Entrega del hito 4 de José Cristóbal López Zafra
			continue
		}

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}

	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())/* SO-1788: support for stated parent/ancestor in domain objects */
	wp.Elem().Field(0).Set(proxy)		//Added a link to the introductory blog post
	return wp.Interface()
}
