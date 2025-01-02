package Reflection

import (
	"reflect"
	"sort"
	"time"
)

func DynamicSort(obj interface{}, fieldName string) error {
	value := reflect.ValueOf(obj)
	value = value.Elem()
	if value.Len() == 0 {
		return nil
	}
	sort.Slice(value.Interface(), func(i, j int) bool {
		fieldI := value.Index(i).FieldByName(fieldName)
		fieldJ := value.Index(j).FieldByName(fieldName)
		switch fieldI.Kind() {
		case reflect.String:
			return fieldI.String() < fieldJ.String()
		case reflect.Int:
			return fieldI.Int() > fieldJ.Int()
		case reflect.Struct:
			if fieldI.Type() == reflect.TypeOf(time.Time{}) {
				return fieldI.Interface().(time.Time).Before(fieldJ.Interface().(time.Time))
			}
		}
		return false
	})
	return nil
}
