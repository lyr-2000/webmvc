package web

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_run_server(t *testing.T) {
	router := DefaultRouter()
	//自定义 context工厂，自己可以修改 里面的实现方法
	router.CtxFactory = func(w http.ResponseWriter, r *http.Request, p router.Params) Ctx {
		return &CtxDefault{
			w, r, p,
		}
	}
	router.Router.ServeFiles("./static", http.Dir("./"))
	//404 找不到页面
	router.Router.OnNotFound(func(w http.ResponseWriter, R *http.Request) {
		fmt.Fprintf(w, "找不到页面了 !!!!")
	})
	router.Router.OnMethodNotAllowed(func(w http.ResponseWriter, R *http.Request) {

		fmt.Fprintf(w, "请求方法不正确 !!!!")
	})
	router.GET("/echo", nil, func(c Ctx) interface{} {
		fmt.Fprintf(c.Response(), "这是个 返回！！！")
		return nil
	})
	router.POST("/test", nil, func(c Ctx) interface{} {

		fmt.Fprintf(c.Response(), "这是个 返回！！！")
		return nil
	})

	router.Listen(8080)
}
