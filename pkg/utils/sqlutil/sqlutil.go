package sqlutil

import (
	"reflect"

	"strings"
)

const TagName = "col"

func AsUnderLine(val string) string {
	var col = strings.Builder{}
	for index, b := range val {
		// userName
		if index == 0 {
			if b >= 'A' && b <= 'Z' {
				b ^= 32
			}
			col.WriteByte(byte(b))
		} else if index > 0 && b >= 'A' && b <= 'Z' {
			col.WriteByte('_')
			col.WriteByte(byte(b ^ 32))
		} else {
			col.WriteRune(b)
		}
	}
	return col.String()
}
func AsCamelName(name string) string {
	var col = strings.Builder{}
	var pre_ bool
	for index,b := range name {
		if b == '_' {
			if index==0 {
				continue
			}
			pre_ = true
			continue
		}
		if index==0 &&(b>='A' && b<='Z') {
			b^=32
		} else if pre_ {
			if b>='a' && b<='z' {
				b^=32
			}
		}

		pre_ = false
		col.WriteRune(b)

	}
	return col.String()
}
func AsBigCamelName(name string) string {
	var col = strings.Builder{}
	var pre_ bool
	for index,b := range name {
		if b == '_' {
			if index==0 {
				continue
			}
			pre_ = true
			continue
		}
		if index==0 &&(b>='a' && b<='z') {
			b^=32

		} else if pre_ {
			if b>='a' && b<='z' {
				b^=32
			}
		}

		pre_ = false
		col.WriteRune(b)

	}
	return col.String()
}
func AsColMap(obj interface{}) map[string]interface{} {
	structValue := reflect.Indirect(reflect.ValueOf(obj))
	var kind = structValue.Kind()
	if kind != reflect.Struct {

		return nil
	}
	var mp = make(map[string]interface{}, structValue.NumField())
	type_ := structValue.Type()
	for i := 0; i < structValue.NumField(); i++ {
		field := type_.Field(i)
		tagName := field.Tag.Get(TagName)
		if tagName == "" {
			tagName = AsUnderLine(field.Name)
		}
		mp[tagName] = structValue.Field(i).Interface()
	}
	return mp
}


