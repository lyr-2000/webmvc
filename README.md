## 介绍

这是一个简单的 mvc 框架，封装了非常少的方法，可以自己diy

[使用的路由框架](https://github.com/julienschmidt/httprouter)

[反射工具类](https://github.com/spf13/cast)




## 使用方法



```go
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
```



## 自定义 modelAndView ,实现 mvc



```go
package base

import (
	"encoding/json"
	"godoc/pkg/web"
	"log"
	"text/template"
)

type (
	Resp struct {
		Code int8        `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data" form:"data"`
		ErrMsg string  `json:"err_msg,omitempty"`
	}

	ModelAndView interface {
		GetData() interface{}
		GetView() string
	}
	ModelAndViewObj struct {
		Data interface{}
		View string
	}
)

func (a *ModelAndViewObj) GetData() interface{} {
	return a.Data
}
func (a *ModelAndViewObj) GetView() string {
	return a.View
}

//自定义 function
var viewFuncMap = map[string]interface{}{}

func HandleModelAndView(c web.Ctx, res interface{}) bool {
	mv, isMv := res.(ModelAndView)
	if !isMv {
		return false
	}
	//var view = "./static/" + mv.GetView()
	var view = mv.GetView()
	var globalTpl = template.Must(template.ParseGlob("static/*.html"))
	t := globalTpl.Lookup(view)
	if t == nil {
		log.Printf("resolve template error [%v] \n", view)
		return true
	}
	t = t.Funcs(viewFuncMap)
	err := t.ExecuteTemplate(c.Response(), view, mv.GetData())
	//err := t.Execute(c.Response(), mv.GetData())
	if err != nil {
		log.Printf("render error [%+v]\n", err)
		return true
	}
	return true
}

func HandleOk(ctx web.Ctx, res interface{}) bool {
	//log.Printf("%v\n", res)
	if res != nil {
		_, isErr := res.(error)
		if isErr {
			return false
		}

	}

	var result Resp
	result.Data = &res
	result.Msg = "OK"
	body, err := json.Marshal(result)
	if err != nil {
		log.Printf("json marshal error [%+v]", err)
		return false
	}
	var w = ctx.Response()
	w.Header().Set("content-type", "text/json")
	w.Write(body)
	return true

}
func HandleError(ctx web.Ctx, res interface{}) bool {
	err, isErr := res.(error)
	if !isErr {
		return false
	}
	log.Printf("handle error info [%+v]", err)
	var result Resp
	result.Code = 1
	result.Data = err
	result.ErrMsg = err.Error()
	//result
	body, err := json.Marshal(result)
	var w = ctx.Response()
	w.Header().Set("content-type", "text/json")
	w.Write(body)
	return true
}

```





```go
func InitRouter() {
	webRouter := web.DefaultRouter()
	//后置 处理器，用于 处理 返回的结果
	webRouter.PostResultHandler = []web.PostHandler{
		base.HandleOk,
		base.HandleError,
		base.HandleModelAndView,
	}
    webRouter.GET("/test0/:name",nil, func(c Ctx) interface{} {
		var Params struct{
			Name string
			Value string
		}
		c.BindParam(&Params)
		return Params
	})
    
	//注册路由 ----
	//sqlclient.RegisterSqlClient(webRouter)
	//echo.RegisterEchoClient(webRouter)
	webRouter.Listen(8080)
}

```

