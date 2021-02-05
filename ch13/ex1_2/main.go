package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
type cyclicType struct {
	c *cyclicType
}

func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}
func main() {
	fmt.Printf("require false, actual %t \n", Equal(1.9, 2.0))
	fmt.Printf("require true, actual %t \n", Equal(2.0, 2.0))
	fmt.Printf("require false, actual %t \n", Equal(2.0, 2+1e-9))
	fmt.Printf("require true, actual %t \n", Equal(2.0, 2+1e-11))
	fmt.Printf("require true, actual %t \n", Equal(complex(1, 1), complex(1, 1)))
	fmt.Printf("require true, actual %t \n", Equal(complex(1, 1+1e-11), complex(1, 1)))
	fmt.Printf("require false, actual %t \n", Equal(complex(1, 1+1e-8), complex(1, 1)))
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true
	case reflect.Float32, reflect.Float64:
		return equalFloats(x.Float(), y.Float())
	case reflect.Complex64, reflect.Complex128:
		return equalComplexes(x.Complex(), y.Complex())
	}

	panic("unreachable")
}

func equalFloats(x, y float64) bool {
	if x >= y {
		return x-y <= 1.0e-10
	}
	return y-x <= 1.0e-10
}
func equalComplexes(x, y complex128) bool {
	r := real(x) - real(y)
	i := imag(x) - imag(y)
	return r*r+i*i <= 1.0e-20
}
func cyclic(v reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !v.IsValid() {
		return false
	}
	if v.CanAddr() && v.Kind() != reflect.Struct && v.Kind() != reflect.Array {
		ptr := unsafe.Pointer(v.UnsafeAddr())
		if seen[ptr] {
			return true
		}
		seen[ptr] = true
	}
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if cyclic(v.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if cyclic(v.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if cyclic(v.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	return false
}
