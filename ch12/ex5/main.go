package main

import (
	"bytes"
	"encoding/json"
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
	b, err := JsonMarshal(strangelove)
	if err != nil {
		fmt.Printf("Error:%v \n", err)
		return
	}
	fmt.Printf("%s\n", b)

	c, err := json.Marshal(strangelove)
	if err != nil {
		fmt.Printf("Error:%v \n", err)
		return
	}

	fmt.Printf("%s\n", c)
	var cMovie Movie
	var bMovie Movie
	json.Unmarshal(c, &cMovie)
	json.Unmarshal(b, &bMovie)
	fmt.Printf("%v\n", cMovie)
	fmt.Printf("%v\n", bMovie)
}

func JsonMarshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func encode(buf *bytes.Buffer, v reflect.Value, tab int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

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
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(',')
			}
			buf.WriteByte('\n')
			for i := 0; i < (tab+1)*2; i++ {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i), tab+1); err != nil {
				return err
			}
		}
		buf.WriteByte('\n')
		for i := 0; i < tab*2; i++ {
			buf.WriteByte(' ')
		}
		buf.WriteByte(']')
	case reflect.Struct:
		createOpenBrace(buf, tab*2)
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(',')
				buf.WriteByte('\n')
			}
			for i := 0; i < (tab+1)*2; i++ {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "\"%s\": ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), tab+1); err != nil {
				return err
			}
			for i := 0; i < tab; i++ {
				buf.WriteByte(' ')
			}
		}
		createCloseBrace(buf, tab*2)
	case reflect.Map:
		createOpenBrace(buf, 0)
		first := true
		for _, key := range v.MapKeys() {
			if !first {
				buf.WriteByte(',')
				buf.WriteByte('\n')
			}
			for i := 0; i < (tab+1)*2; i++ {
				buf.WriteByte(' ')
			}
			if key.Kind() != reflect.String {
				return fmt.Errorf("unsupported map key: %s", key.Kind())
			}
			if err := encode(buf, key, tab+1); err != nil {
				return err
			}
			buf.WriteByte(':')
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), tab+1); err != nil {
				return err
			}
			first = false
		}
		createCloseBrace(buf, 2*tab)
	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
func createOpenBrace(buf *bytes.Buffer, tab int) {
	for i := 0; i < tab; i++ {
		buf.WriteByte(' ')
	}
	buf.WriteByte('{')
	buf.WriteByte('\n')
}
func createCloseBrace(buf *bytes.Buffer, tab int) {
	buf.WriteByte('\n')
	for i := 0; i < tab; i++ {
		buf.WriteByte(' ')
	}
	buf.WriteByte('}')
}
