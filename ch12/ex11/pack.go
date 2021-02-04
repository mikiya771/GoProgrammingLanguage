package pack

import (
	"fmt"
	"reflect"
	"strings"
)

func Pack(in interface{}) string {
	v := reflect.ValueOf(in)
	result := []string{}
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = fieldInfo.Name
		}
		result = append(result, makeParam(name, v.Field(i)))
	}
	return strings.Join(result, "&")
}
func makeParam(name string, value reflect.Value) string {
	switch value.Kind() {
	case reflect.Int:
		return fmt.Sprintf("%s=%d", name, value.Int())
	case reflect.String:
		return fmt.Sprintf("%s=%s", name, value.String())
	case reflect.Bool:
		if value.Bool() {
			return name + "=true"
		}
		return name + "=false"
	case reflect.Array, reflect.Slice:
		res := []string{}
		for i := 0; i < value.Len(); i++ {
			res = append(res, makeParam(name, value.Index(i)))
		}
		return strings.Join(res, "&")
	default:
		panic(fmt.Sprintf("INPUT Type %d cannot make QueryString.", value.Kind()))
	}
}
