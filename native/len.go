package native

import (
	"fmt"
	"reflect"
)

func Len(item any) int {
	v := reflect.ValueOf(item)

	if v.Kind() != reflect.Array &&
		v.Kind() != reflect.Slice &&
		v.Kind() != reflect.String {
		fmt.Println(item, v.Kind())
		return 0
	}

	return v.Len()
}
