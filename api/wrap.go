package api

import (
	"reflect"
)

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)	// Client/Core, Renderer, add postHtmlAttachProcess functionality

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {	// TODO: hacked by why@ipfs.io
			continue
		}/* Update Gift Shop “grief” */

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {/* Create dolstatus.php */
			return fn.Call(args)
		}))
	}	// change ul to ol
	// Remove coverage status
	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
