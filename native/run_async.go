package native

import (
	"fmt"
	"reflect"

	"github.com/samber/lo"
)

type Callback func(retVal ...reflect.Value)

func RunAsync(fn any, cb Callback, args ...any) error {
	v := reflect.ValueOf(fn)

	if v.Kind() != reflect.Func {
		return fmt.Errorf("argument must be a function")
	}

	argVals := make([]reflect.Value, 0)
	argVals = lo.Map(args, func(arg any, i int) reflect.Value {
		return reflect.ValueOf(arg)
	})

	go func() {
		retVal := v.Call(argVals)
		cb(retVal...)
	}()

	return nil
}
