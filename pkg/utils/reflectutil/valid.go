package reflectutil

import "reflect"

func IsAnyZeroValue(res interface{}) bool {
	value := reflect.Indirect(reflect.ValueOf(res))

	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).IsZero() {
			return true
		}
	}
	return false
}
