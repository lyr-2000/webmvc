package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"webmvc/pkg/router"
)

func Test_run_server(t *testing.T) {
	r := DefaultRouter()
	//自定义 context工厂，自己可以修改 里面的实现方法
	r.CtxFactory = func(w http.ResponseWriter, r *http.Request, p router.Params) Ctx {
		return &CtxDefault{
			w, r, p,
		}
	}
	//后置处理器，用于处理返回的结果
	r.PostResultHandler = []PostHandler{
		func(ctx Ctx, res interface{}) bool {
			if res != nil {
				log.Println("用来打印日志......")
				return true
			}
			return true
		},
		func(c Ctx, res interface{}) bool {
			if res == nil {
				return true
			}
			bs, _ := json.Marshal(res)
			c.Response().Header().Set("content-type", "text/json")
			c.Response().Write(bs)
			return true
		},

	}

	//404 找不到页面
	r.Router.OnNotFound(func(w http.ResponseWriter, R *http.Request) {
		fmt.Fprintf(w, "找不到页面了 !!!!")
	})
	r.Router.OnMethodNotAllowed(func(w http.ResponseWriter, R *http.Request) {

		fmt.Fprintf(w, "请求方法不正确 !!!!,当前方法 [%v]",R.Method)
	})
	r.GET("/echo", nil, func(c Ctx) interface{} {
		fmt.Fprintf(c.Response(), "这是个 返回！！！")
		return nil
	})
	r.GET("/test0/:name",nil, func(c Ctx) interface{} {
		var Params struct{
			Name string
			Value string
		}
		c.BindParam(&Params)
		return Params
	})
	r.POST("/test", nil, func(c Ctx) interface{} {

		fmt.Fprintf(c.Response(), "这是个 返回！！！")
		return nil
	})

	r.Listen(8080)
}
