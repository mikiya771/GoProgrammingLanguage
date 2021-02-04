package main

import (
	"bytes"
	"fmt"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol struct {
	Name string
}
type String struct {
	Name string
}
type Int struct {
	Num int
}
type StartList struct{}
type EndList struct{}

type Decoder struct {
	lex *lexer
}

func NewDecoder(data []byte) *Decoder {
	var d Decoder
	d.lex = &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	d.lex.scan.Init(bytes.NewReader(data))
	d.lex.next() //lex.token 初期化
	return &d
}

func (d *Decoder) Token() (Token, error) {
	switch d.lex.token {
	case scanner.Ident:
		name := d.lex.text()
		d.lex.next()
		return Symbol{name}, nil
	case scanner.String:
		v, err := strconv.Unquote(d.lex.text())
		d.lex.next()
		if err != nil {
			return nil, err
		}

		return String{v}, nil
	case scanner.Int:
		v, err := strconv.Atoi(d.lex.text())
		d.lex.next()
		if err != nil {
			return nil, err
		}
		return Int{v}, nil
	case '(':
		d.lex.next()
		return StartList{}, nil
	case ')':
		d.lex.next()
		return EndList{}, nil
	}
	panic(fmt.Sprintf("%d is unknown rune", d.lex.token))
}
