package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

type Var string
type literal float64
type unary struct {
	op rune
	x  Expr
}
type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}
type list struct {
	fn   string
	args []Expr
}

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	String() string
}

func (v Var) String() string {
	return fmt.Sprintf("%s", string(v))
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) String() string {
	return fmt.Sprintf("%s %s", string(u.op), u.x.String())
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x.String(), string(b.op), b.y.String())
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) String() string {
	strs := []string{}
	for _, str := range c.args {
		strs = append(strs, str.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(strs, ","))
}
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported call operator: %q", c.fn))
}
func (l list) String() string {
	strs := []string{}
	for _, str := range l.args {
		strs = append(strs, str.String())
	}
	return fmt.Sprintf("%s(%s)", l.fn, strings.Join(strs, ","))
}
func (l list) Eval(env Env) (lt float64) {
	switch l.fn {
	case "min":
		lt = l.args[0].Eval(env)
		for _, ls := range l.args {
			lt = math.Min(lt, ls.Eval(env))
		}
		return
	case "max":
		lt = l.args[0].Eval(env)
		for _, ls := range l.args {
			lt = math.Max(lt, ls.Eval(env))
		}
		return
	}
	return
}

var token = map[rune]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'(': true,
	')': true,
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

type lexPanic string

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token))
}

func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func Parse(str string) (expr Expr, err error) {
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(str))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next()
	e := parseExpr(lex)
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

func parseExpr(lex *lexer) Expr { return parseBinary(lex, 1) }

// binary = unary ('+' binary)*
// parseBinary stops when it encounters an
// operator of lower precedence than prec1.
func parseBinary(lex *lexer, prec1 int) Expr {
	lhs := parseUnary(lex)
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		for precedence(lex.token) == prec {
			op := lex.token
			lex.next() // consume operator
			rhs := parseBinary(lex, prec+1)
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// unary = '+' expr | primary
func parseUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next() // consume '+' or '-'
		return unary{op, parseUnary(lex)}
	}
	return parsePrimary(lex)
}
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next()
		//for variable
		if lex.token != '(' {
			return Var(id)
		}
		// for 'math(' function(expr,expr,,)
		lex.next()
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parseExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next()
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %q, want ')'", lex.token)
				panic(lexPanic(msg))
			}
		}
		lex.next()
		if id == "min" || id == "max" {
			return list{id, args}
		}

		return call{id, args}
	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next()
		return literal(f)
	case '(':
		lex.next()
		e := parseExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next()
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	exprStr := scanner.Text()
	envs := map[Var]float64{}

	expr, err := Parse(exprStr)
	if err != nil {
		fmt.Errorf("parse error: %v", err)
	}
	vs := getVars(expr)
	for _, v := range vs {
		fmt.Printf("Please Input Value %s: \n", v)
		scanner.Scan()
		f := scanner.Text()
		envs[Var(v)], err = strconv.ParseFloat(f, 64)
		if err != nil {
			fmt.Errorf("%v", err)
			return
		}
	}
	fmt.Printf("%s = %f, env: %v \n", exprStr, expr.Eval(envs), envs)

}

func getVars(expr Expr) (vars []string) {
	wfs := []Expr{expr}
	for len(wfs) > 0 {
		nextWfs := []Expr{}
		for _, e := range wfs {
			switch reflect.TypeOf(e).Name() {
			case "Var":
				vars = append(vars, e.(Var).String())
			case "unary":
				nextWfs = append(nextWfs, e.(unary).x)
			case "binary":
				nextWfs = append(nextWfs, e.(binary).x, e.(binary).y)
			case "call":
				nextWfs = append(nextWfs, e.(call).args...)
			case "list":
				nextWfs = append(nextWfs, e.(list).args...)
			}
		}
		wfs = nextWfs
	}
	return
}
