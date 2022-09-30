package utils

import "reflect"

func StructToMap(obj any) map[string]any {
	var result = make(map[string]any)
	var value = reflect.ValueOf(obj)
	var typeOfS = value.Type()

	for i := 0; i < value.NumField(); i++ {
		result[typeOfS.Field(i).Name] = value.Field(i).Interface()
	}

	return result
}
