package reflection

import (
	"reflect"
)

func Walk(x any, fn func(string)) {
	value := getValue(x)

	var numberOfValues int
	var getValue func(int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
		return
	case reflect.Slice, reflect.Array:
		numberOfValues = value.Len()
		getValue = value.Index
	case reflect.Struct:
		numberOfValues = value.NumField()
		getValue = value.Field
	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			Walk(iter.Key().Interface(), fn)
			Walk(iter.Value().Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		Walk(getValue(i).Interface(), fn)
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
