package router

import (
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, R *http.Request)

//type HandleController struct {
//
//}

func (r *Router) registerHandler(method string, path string, handler Handler) {
	report_(method, path)
	r.Handle(method, path, func(writer http.ResponseWriter, request *http.Request, _ Params) {
		handler(writer, request)
	})
}
func report_(method string, path string) {
	log.Printf("router.log  [%-6v] - [%v]  \n", method, path)
}

// func (r *Router) Get(path string, handler Handler) {
// 	r.registerHandler(http.MethodGet, path, handler)
// }

// func (r *Router) Post(path string, handler Handler) {
// 	r.registerHandler(http.MethodPost, path, handler)
// }

// func (r *Router) Put(path string, handler Handler) {
// 	r.registerHandler(http.MethodPut, path, handler)
// }
// func (r *Router) Patch(path string, handler Handler) {
// 	r.registerHandler(http.MethodPatch, path, handler)
// }

// func (r *Router) Delete(path string, handler Handler) {
// 	r.registerHandler(http.MethodDelete, path, handler)

// }

type HttpController struct {
	InnerHandler Handler
}

func (hc *HttpController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hc.InnerHandler != nil {
		hc.InnerHandler(w, r)
	}
}

func (r *Router) OnMethodNotAllowed(h Handler) {
	var inner HttpController
	inner.InnerHandler = h
	r.MethodNotAllowed = &inner
}

func (r *Router) OnPanic(h func(http.ResponseWriter, *http.Request, interface{})) {
	r.PanicHandler = h
}
func (r *Router) OnNotFound(h Handler) {
	var inner HttpController
	inner.InnerHandler = h
	r.NotFound = &inner
}

//func (r *Router) GET(path string, handle Handle) {
//	r.Handle(http.MethodGet, path, handle)
//}
