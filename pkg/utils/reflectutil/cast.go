package reflectutil

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"time"
)

func setAddr_(destStruct *reflect.Value, from interface{}) {
	destStruct.Set(reflect.ValueOf(from))
}
func IsPrimitive(dst interface{}) bool {
	if dst == nil {
		return false
	}
	switch dst.(type) {
	case *int, *int8, *int16, *int32, *int64,
		*uint, *uint8, *uint16, *uint32, *uint64,
		*float32, *float64, *string, *bool, *time.Time:
		return true
	default:
		return false
	}
	return false

}

func SetPrimitive(dest interface{}, from interface{}) error {
	destStruct := reflect.Indirect(reflect.ValueOf(dest))
	dtype := destStruct.Kind()
	//log.Println("type=", dtype,destStruct.Interface() )
	switch dtype {
	case reflect.String:
		setAddr_(&destStruct, cast.ToString(from))

	case reflect.Int8:
		setAddr_(&destStruct, cast.ToInt8(from))
	case reflect.Int16:
		setAddr_(&destStruct, cast.ToInt16(from))
	case reflect.Int32:
		setAddr_(&destStruct, cast.ToInt32(from))
	case reflect.Int:
		setAddr_(&destStruct, cast.ToInt(from))
	case reflect.Int64:
		setAddr_(&destStruct, cast.ToInt64(from))
	case reflect.Float32:
		setAddr_(&destStruct, cast.ToFloat32(from))
	case reflect.Float64:
		setAddr_(&destStruct, cast.ToFloat64(from))
	case reflect.Bool:
		setAddr_(&destStruct, cast.ToBool(from))
	case reflect.Uint8:
		setAddr_(&destStruct, cast.ToUint8(from))
	case reflect.Uint:
		setAddr_(&destStruct, cast.ToUint(from))
	case reflect.Uint16:
		setAddr_(&destStruct, cast.ToUint16(from))
	case reflect.Uint32:
		setAddr_(&destStruct, cast.ToUint32(from))
	case reflect.Uint64:
		setAddr_(&destStruct, cast.ToUint64(from))
	default:
		switch dest.(type) {
		case *time.Time:
			setAddr_(&destStruct, cast.ToTime(from))
		case *time.Duration:
			setAddr_(&destStruct, cast.ToDuration(from))
		default:
			return fmt.Errorf("cannot parsed values")
		}
	}
	return nil

}

func SetString(dest *string, from interface{}) {
	*dest = cast.ToString(from)
}
