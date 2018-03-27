package array

import "reflect"

// ToInterfaces converts an array/slice into an interface slice
// MySQL select ... in (...)
func ToInterfaces(arr interface{}) (items []interface{}) {
	arrValue := reflect.ValueOf(arr)
	if arrValue.Kind() == reflect.Ptr {
		arrValue = arrValue.Elem()
	}
	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil
	}

	arrSize := arrValue.Len()
	for i := 0; i < arrSize; i++ {
		items = append(items, arrValue.Index(i).Interface())
	}

	return
}
