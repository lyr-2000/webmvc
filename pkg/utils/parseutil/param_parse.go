package parseutil

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"webmvc/pkg/utils/sqlutil"
	"webmvc/pkg/utils/strutil"
)

const TagName = "alias"

func Split(value string) []string {
	return strings.FieldsFunc(value, func(c rune) bool {
		return c == ';' || c == ',' || c == '|' || c == '/' || c == '\\'
	})

}

// func ParseValueString(exp string) string {
// 	strings.Split(exp,"")
// }

func ChangeType(t reflect.Kind, obj interface{}) interface{} {

	switch t {
	case reflect.Uint:
		if obj == nil {
			return 0
		}
		switch obj.(type) {
		case int:
			return uint(obj.(int))
		case float64:
			return uint(obj.(float64))
		case float32:
			return uint(obj.(float32))
		case int16:
			return uint(obj.(int16))
		case int32:
			return uint(obj.(int32))
		case int64:
			return uint(obj.(int64))
		case string:
			return uint(strutil.ParseFloat64(obj.(string)))
		}

	case reflect.Int:
		if obj == nil {
			return 0
		}
		switch obj.(type) {
		case int:
			return obj
		case float64:
			return int(obj.(float64))
		case float32:
			return int(obj.(float32))
		case int16:
			return int(obj.(int16))
		case int32:
			return int(obj.(int32))
		case int64:
			return int(obj.(int64))
		case string:
			return int(strutil.ParseFloat64(obj.(string)))
		}

	case reflect.Int64:
		switch obj.(type) {
		case int:
			return int64(obj.(int))
		case int8:
			return int64(obj.(int8))
		case int16:
			return int64(obj.(int16))
		case int32:
			return int64(obj.(int32))
		case float32:
			return int64(obj.(float32))
		case float64:
			return int64(obj.(float64))
		case int64:
			return obj
		case string:
			s := strutil.ParseFloat64(obj.(string))
			return float64(s)
		}
	case reflect.Float64:
		switch obj.(type) {
		case int:
			return float64(obj.(int))
		case float32:
			return float64(obj.(float32))
		case float64:
			return obj
		case string:
			s, _ := strconv.ParseFloat(obj.(string), 64)
			return float64(s)
		}

	case reflect.String:
		if obj == nil {
			return ""
		}
		switch obj.(type) {
		case string:
			return obj

		}
		return fmt.Sprintf("%v", obj)

	default:

	}
	return obj
}

func CopyFromMap(dest interface{}, mp map[string]interface{}) {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("error: [error recover := %+v ]\n", err)
		}
	}()
	destStruct := reflect.Indirect(reflect.ValueOf(dest))
	//注意，要从这里获取
	destType := destStruct.Type()

	for i := 0; i < destType.NumField(); i++ {
		var tg = destType.Field(i)
		s := tg.Tag.Get(TagName)
		if s == "" {
			continue
		}
		//获取别名
		var arr = Split(s)
		for _, name := range arr {
			var val, ok = mp[name]
			if ok {
				fieldRef := destStruct.Field(i)
				//将两个对象的匹配类型 匹配，并且转化
				f := ChangeType(fieldRef.Kind(), val)

				if f != nil {
					fieldRef.Set(reflect.ValueOf(f))
					break
				}

			}
		}
	}

}

func CopyOutMap(from interface{}) map[string]interface{} {
	destStruct := reflect.Indirect(reflect.ValueOf(from))
	destType := destStruct.Type()
	var mp = make(map[string]interface{}, destType.NumField())
	for i := 0; i < destStruct.NumField(); i++ {
		var field = destType.Field(i)
		tagvalue := field.Tag.Get(TagName)
		var aliasName = Split(tagvalue)
		if len(aliasName) == 0 {
			aliasName = append(aliasName, sqlutil.AsCamelName(field.Name))
		}
		mp[aliasName[0]] = destStruct.Field(i).Interface()

	}
	return mp
}
