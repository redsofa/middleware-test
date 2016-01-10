package handlers

import (
	"net/http"
)

/*
  The source for the Adapter code is :

  https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81#.889urq1mh
*/

//The Adapter type is a function that takes in http.Handler and returns an http.Handler
type Adapter func(http.Handler) http.Handler

//The Execute function takes :
//
//	1) - a http.Handler
//	2) - an arbitrary number of adapters
// and
//  returns an http.Handler function
func Execute(h http.Handler, adapters ...Adapter) http.Handler {
	//Note that we're ranging in reverse order.
	//The last piece of middleware will be the first
	//to execute.
	for i := range adapters {
		adapter := adapters[len(adapters)-1-i]
		h = adapter(h)
	}
	return h
}
