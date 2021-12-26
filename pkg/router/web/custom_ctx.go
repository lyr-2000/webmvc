package web

import (
	"fmt"
	"log"
	"net/http"
	"webmvc/pkg/router"
)

type CtxDefault struct {
	W      http.ResponseWriter
	R      *http.Request
	Params router.Params
}

func (c CtxDefault) Bind(res interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c CtxDefault) BindJSON(res interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c CtxDefault) BindParam(res interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c CtxDefault) BindQuery(res interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c CtxDefault) BindForm(res interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c CtxDefault) Response() http.ResponseWriter {
	//TODO implement me
	return c.W
	//panic("implement me")
}

func (c CtxDefault) Request() *http.Request {
	//TODO implement me
	return c.R
}

func (c CtxDefault) Parameters() router.Params {
	//TODO implement me
	return c.Params
}

type Ctx interface {
	Response() http.ResponseWriter
	Request() *http.Request
	Parameters() router.Params

	Bind(res interface{}) error
	BindJSON(res interface{}) error
	BindParam(res interface{}) error
	BindQuery(res interface{}) error
	BindForm(res interface{}) error
}

type Cmd interface {
}
type CmdDefault struct {
	Data interface{}
	Err  error
}

type Router struct {
	Router            *router.Router
	PostResultHandler []PostHandler
	CtxFactory        func(w http.ResponseWriter, r *http.Request, params router.Params) Ctx
}

func (r *Router) RegisterHandler(method string, path string, proxyHandler func(c Ctx, h Handler) interface{}, h Handler) {
	log.Printf("[%v] -> [%v]\n", method, path)
	r.Router.Handle(method, path, func(writer http.ResponseWriter, request *http.Request, p router.Params) {
		var ctx Ctx
		if r.CtxFactory != nil {
			ctx = r.CtxFactory(writer, request, p)
		}
		if ctx == nil {
			ctx = &CtxDefault{
				W:      writer,
				R:      request,
				Params: p,
			}
		}
		var res interface{}
		//使用代理的方法
		if proxyHandler != nil {
			res = proxyHandler(ctx, h)
		} else {
			//直接调用 本方法
			res = h(ctx)
		}
		if res == nil {
			return
		}
		//如果 有返回值，就调用 后置处理器处理返回值
		for _, postHandler := range r.PostResultHandler {
			var ok = postHandler(ctx, res)
			if ok {
				return
			}
		}
		log.Printf("无法处理执行结果 [%v]\n", request.RequestURI)
		// if  not ok ,use default case

	})
}

type Handler func(c Ctx) interface{}

type PostHandler func(ctx Ctx, res interface{}) bool

//
//func DefaultProxyHandler(c Ctx, h Handler) interface{} {
//	var res = h(c)
//	return res
//}

func DefaultRouter() *Router {
	return &Router{
		Router:            router.New(),
		PostResultHandler: nil,
		CtxFactory:        nil,
	}
}

func (r *Router) GET(path string, proxyHandler func(c Ctx, h Handler) interface{}, f Handler) {
	r.RegisterHandler(http.MethodGet, path, proxyHandler, f)
}

// HEAD is a shortcut for router.Handle(http.MethodHead, path, handle)
func (r *Router) HEAD(path string, proxyHandler func(c Ctx, h Handler) interface{}, f Handler) {
	r.RegisterHandler(http.MethodHead, path, proxyHandler, f)
}

// OPTIONS is a shortcut for router.Handle(http.MethodOptions, path, handle)
func (r *Router) OPTIONS(path string, proxyHandler func(c Ctx, h Handler) interface{}, f Handler) {
	r.RegisterHandler(http.MethodOptions, path, proxyHandler, f)
}

// POST is a shortcut for router.Handle(http.MethodPost, path, handle)
func (r *Router) POST(path string, proxyHandler func(c Ctx, h Handler) interface{}, f Handler) {
	r.RegisterHandler(http.MethodPost, path, proxyHandler, f)
}

// PUT is a shortcut for router.Handle(http.MethodPut, path, handle)
func (r *Router) PUT(path string, proxyHandler func(c Ctx, h Handler) interface{}, f Handler) {
	r.RegisterHandler(http.MethodPut, path, proxyHandler, f)
}

// PATCH is a shortcut for router.Handle(http.MethodPatch, path, handle)
func (r *Router) PATCH(path string, proxyHandler func(c Ctx, h Handler) interface{}, f Handler) {
	r.RegisterHandler(http.MethodPatch, path, proxyHandler, f)
}

// DELETE is a shortcut for router.Handle(http.MethodDelete, path, handle)
func (r *Router) DELETE(path string, proxyHandler func(c Ctx, h Handler) interface{}, f Handler) {
	r.RegisterHandler(http.MethodDelete, path, proxyHandler, f)
}

func (r *Router) Listen(port int) error {
	log.Printf("listen on port = [%d]\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r.Router)
}
