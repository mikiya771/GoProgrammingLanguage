package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "smith",
		Subtitle: "i want an alian",
		Year:     2008,
		Color:    false,
		Actor: map[string]string{
			"kyon":    "sugita",
			"koizumi": "ono",
		},
		Oscars: []string{
			"be",
			"do",
		},
	}
	b, err := Marshal(strangelove)
	if err != nil {
		fmt.Printf("Error:%v \n", err)
		return
	}

	fmt.Printf("%s\n", b)
	tmp := map[string]interface{}{
		"hoge": []int{1, 2, 3},
	}
	b, err = Marshal(tmp)
	if err != nil {
		fmt.Printf("Error:%v \n", err)
		return
	}

	fmt.Printf("%s", b)
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func encode(buf *bytes.Buffer, v reflect.Value, tab int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), tab)

	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i), tab); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		createOpenBrace(buf, tab)
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				createOpenBrace(buf, tab+1)
			} else {
				buf.WriteByte('(')
			}
			fmt.Fprintf(buf, "%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), tab+1+1+len(v.Type().Field(i).Name)); err != nil {
				return err
			}
			buf.WriteByte(')')
			for i := 0; i < tab; i++ {
				buf.WriteByte(' ')
			}
		}
		buf.WriteByte(')')
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte('\n')
				createOpenBrace(buf, tab+2)
			} else {
				buf.WriteByte('(')
			}
			if err := encode(buf, key, tab); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), tab); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(c), imag(c))
	case reflect.Interface:
		buf.WriteByte('(')
		t := v.Type()
		if t.Name() == "" {
			fmt.Fprintf(buf, "%q ", v.Elem().Type().String())
		} else {
			fmt.Fprintf(buf, `"%s.%s" `, t.PkgPath(), t.Name())
		}
		if err := encode(buf, v.Elem(), tab); err != nil {
			return err
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
func createOpenBrace(buf *bytes.Buffer, tab int) {
	for i := 0; i < tab; i++ {
		buf.WriteByte(' ')
	}
	buf.WriteByte('(')
}
