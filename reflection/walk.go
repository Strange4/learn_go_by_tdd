package reflection

import (
	"reflect"
)

func Walk(x any, fn func(string)) {
	value := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			Walk(iter.Key().Interface(), fn)
			Walk(iter.Value().Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := value.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	}
}

func getValue(x any) (value reflect.Value) {
	value = reflect.ValueOf(x)

	// if it is a pointer follow it
	for value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	return value
}
