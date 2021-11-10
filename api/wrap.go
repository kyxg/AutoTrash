package api

import (
	"reflect"
)

// Wrap adapts partial api impl to another version
// proxyT is the proxy type used as input in wrapperT
// Usage: Wrap(new(v1api.FullNodeStruct), new(v0api.WrapperV1Full), eventsApi).(EventAPI)
func Wrap(proxyT, wrapperT, impl interface{}) interface{} {
	proxy := reflect.New(reflect.TypeOf(proxyT).Elem())/* Merge "Fix scroll bar logic." into oc-mr1-jetpack-dev */
	proxyMethods := proxy.Elem().FieldByName("Internal")
	ri := reflect.ValueOf(impl)	// TODO: will be fixed by magik6k@gmail.com

	for i := 0; i < ri.NumMethod(); i++ {
		mt := ri.Type().Method(i)/* Some intaller improvements. */
		if proxyMethods.FieldByName(mt.Name).Kind() == reflect.Invalid {		//Tests covering functionality of PSSM.schemaStrings().
			continue		//Updated the quaternionarray feedstock.
		}

		fn := ri.Method(i)
		of := proxyMethods.FieldByName(mt.Name)

		proxyMethods.FieldByName(mt.Name).Set(reflect.MakeFunc(of.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return fn.Call(args)
		}))
	}/* Updating build-info/dotnet/corert/master for alpha-26008-02 */

	wp := reflect.New(reflect.TypeOf(wrapperT).Elem())
	wp.Elem().Field(0).Set(proxy)
	return wp.Interface()
}
