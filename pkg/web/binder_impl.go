package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"webmvc/pkg/utils/reflectutil"
)

func formOrJsonName_(tag reflect.StructTag) string {
	var form = tag.Get("form")
	if form != "" {
		return form
	}
	var jsonName = tag.Get("json")
	return jsonName
}
func getSimpleName(name string) string {
	//返回小驼峰
	str := strings.Builder{}
	for i, b := range name {
		if i == 0 {
			if b == '_' {
				continue
			}
			if b >= 'A' && b <= 'Z' {
				b ^= 32
			}
		}
		str.WriteRune(b)
	}

	return str.String()
}

func (c *CtxDefault) Bind(res interface{}) error {
	var req = c.Request()
	htype := req.Header.Get("Content-Type")
	var err error
	switch htype {
	case "application/json":
		err = c.BindJSON(res)
		if err != nil {
			return err
		}
		err = c.BindQuery(res)
	case "application/x-www-form-urlencoded":
		err = c.BindForm(res)
		if err != nil {
			return nil
		}
		err = c.BindQuery(res)
	default:
		err = c.BindForm(res)
		if err != nil {
			return err
		}
		err = c.BindQuery(res)
	}

	return err
}

func (c *CtxDefault) BindJSON(res interface{}) error {
	var req = c.Request()
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	//log.Printf("%s",string(bytes))
	defer req.Body.Close()
	return json.Unmarshal(bytes, res)

}
func (c CtxDefault) BindParam(res interface{}) error {
	//TODO implement me
	params := c.Params
	structValue := reflect.Indirect(reflect.ValueOf(res))
	//parse form or  json
	n := structValue.NumField()
	structType := structValue.Type()
	var err error
	for i := 0; i < n; i++ {
		if structValue.Field(i).IsZero() == false {
			//log.Println("value = ")
			continue
		}
		key := formOrJsonName_(structType.Field(i).Tag)
		if key == "" {
			key = getSimpleName(structType.Field(i).Name)
		}
		value := params.ByName(key)
		if value == "" {
			continue
		}
		var field = structValue.Field(i)
		e := reflectutil.SetPrimitive(field.Addr().Interface(), value)

		if err == nil && e != nil {
			err = e
		}
	}
	return err
}

func (c CtxDefault) BindQuery(res interface{}) error {
	query := c.R.URL.Query()
	structValue := reflect.Indirect(reflect.ValueOf(res))
	//parse form or  json
	n := structValue.NumField()
	structType := structValue.Type()
	var err error
	for i := 0; i < n; i++ {
		if structValue.Field(i).IsZero() == false {
			//log.Println("value = ")
			continue
		}
		key := formOrJsonName_(structType.Field(i).Tag)
		if key == "" {
			key = getSimpleName(structType.Field(i).Name)
		}
		value := query.Get(key)
		if value == "" {
			continue
		}
		var field = structValue.Field(i)
		e := reflectutil.SetPrimitive(field.Addr().Interface(), value)

		if err == nil && e != nil {
			err = e
		}
	}

	return err
}

func (c CtxDefault) BindForm(res interface{}) error {

	req := c.Request()
	if req.Method == http.MethodGet {
		log.Printf("get请求 无法解析请求体参数")

		return c.BindQuery(res)
	}
	htype := req.Header.Get("Content-Type")
	if htype == "application/json" {
		return fmt.Errorf("无法解析 json格式")
	} else if htype != "application/x-www-form-urlencoded" {
		//解析 multipart/form-data
		_ = req.ParseMultipartForm(64)
	} else {
		_ = req.ParseForm()
	}
	//log.Println(req.Header.Get("Content-Type"))
	f := req.PostForm
	structValue := reflect.Indirect(reflect.ValueOf(res))
	//parse form or  json
	n := structValue.NumField()
	structType := structValue.Type()
	var err error
	for i := 0; i < n; i++ {
		if structValue.Field(i).IsZero() == false {
			//log.Println("value = ")
			continue
		}
		key := formOrJsonName_(structType.Field(i).Tag)
		if key == "" {
			key = getSimpleName(structType.Field(i).Name)
		}
		value := f.Get(key)
		if value == "" {
			continue
		}
		var field = structValue.Field(i)
		e := reflectutil.SetPrimitive(field.Addr().Interface(), value)

		if err == nil && e != nil {
			err = e
		}
	}
	return err
}
