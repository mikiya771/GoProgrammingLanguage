package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}

func (d *Decoder) Decode(out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(d.r)
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	tags := make(map[string]string)
	v := reflect.ValueOf(out).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("sexpr")
		if name == "" {
			name = fieldInfo.Name
		}
		tags[name] = fieldInfo.Name
	}
	read(lex, reflect.ValueOf(out).Elem(), tags)
	return nil
}

func Unmarshal(data []byte, out interface{}) (err error) {
	d := NewDecoder(bytes.NewReader(data))
	return d.Decode(out)
}

func read(lex *lexer, v reflect.Value, tags map[string]string) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, err := strconv.Unquote(lex.text())
		if err != nil {
			panic(fmt.Sprintf("parse error: cannnot interpret %s as a String", s))
		}
		v.SetString(s)
		lex.next()
		return

	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v, tags)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}
func readList(lex *lexer, v reflect.Value, tags map[string]string) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i), tags)
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item, tags)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want filed name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(tags[name]), tags)
			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key, tags)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value, tags)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	default:
		panic(fmt.Sprintf("cannot decode list int %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
