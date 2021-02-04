package display

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

func Display(name string, x interface{}, w io.Writer) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), w)
}

const MaxDepth = 3

func formatAtom(v reflect.Value, currentDepth int) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Array:
		if currentDepth >= MaxDepth {
			return v.Type().String() + " value"
		}
		a := []string{}
		for i := 0; i < v.Len(); i++ {
			a = append(a, formatAtom(v.Index(i), currentDepth+1))
		}
		return v.Type().String() + "{" + strings.Join(a, ", ") + "}"
	case reflect.Struct:
		if currentDepth >= MaxDepth {
			return v.Type().String() + " value"
		}
		a := []string{}
		for i := 0; i < v.NumField(); i++ {
			a = append(a, fmt.Sprintf("%s: %s", v.Type().Field(i).Name, formatAtom(v.Field(i), currentDepth+1)))
		}
		return v.Type().String() + "{" + strings.Join(a, ", ") + "}"
	default:
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value, w io.Writer) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(w, "%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), w)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), w)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key, 0)), v.MapIndex(key), w)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprintf(w, "%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), w)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(w, "%s = nil\n", path)
		} else {
			fmt.Fprintf(w, "%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), w)
		}
	default:
		fmt.Fprintf(w, "%s = %s\n", path, formatAtom(v, 0))
	}
}
