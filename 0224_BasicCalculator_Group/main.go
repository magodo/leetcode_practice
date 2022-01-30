package foo

import (
	"math"
	"strconv"
)

type token string

const EOF token = "EOF"

var precedenceLookup = map[token]int{
	"+":      1,
	"-":      1,
	"*":      2,
	"/":      2,
	"**":     3,
	"prefix": 4,
}

var funcLookup = map[token]func(int, int) int{
	"+":  func(i, j int) int { return i + j },
	"-":  func(i, j int) int { return i - j },
	"*":  func(i, j int) int { return i * j },
	"/":  func(i, j int) int { return i / j },
	"**": func(i, j int) int { return int(math.Pow(float64(i), float64(j))) },
}

type lexer struct {
	idx  int
	toks []token
}

func newLexer(s string) *lexer {
	toks := []token{}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+', '-', '/', '(', ')':
			toks = append(toks, token(s[i]))
			continue
		case '*':
			if i < len(s)-1 && s[i+1] == '*' {
				i++
				toks = append(toks, token("**"))
				continue
			}
			toks = append(toks, token(s[i]))
			continue
		case ' ':
			continue
		}
		sidx := i
		for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		}
		toks = append(toks, token(s[sidx:i]))
		i -= 1
	}
	toks = append(toks, EOF)

	return &lexer{
		idx:  0,
		toks: toks,
	}
}

func (l *lexer) nextToken() {
	if l.idx < len(l.toks)-1 {
		l.idx++
	}
}

func (l lexer) curToken() token {
	return l.toks[l.idx]
}
func (l lexer) peekToken() token {
	if l.curToken() != EOF {
		return l.toks[l.idx+1]
	}
	return EOF
}

func calculate(s string) int {
	l := newLexer(s)
	return l.calcExpr(0)
}

func (l *lexer) calcExpr(precedence int) int {
	var left int
	switch l.curToken() {
	case "(":
		l.nextToken()
		left = l.calcExpr(0)
		l.nextToken() // ")"
	case "-":
		l.nextToken()
		left = -l.calcExpr(precedenceLookup["prefix"])
	default:
		left = toNum(l.curToken())
	}

	for l.peekToken() != EOF && precedence < precedenceLookup[l.peekToken()] {
		l.nextToken()
		op := l.curToken()
		opPrecedence, opFunc := precedenceLookup[op], funcLookup[op]

		l.nextToken()
		switch l.curToken() {
		case "(":
			l.nextToken()
			left = opFunc(left, l.calcExpr(0))
			l.nextToken() // ")"
		default:
			left = opFunc(left, l.calcExpr(opPrecedence))
		}
	}
	return left
}

func toNum(tok token) int {
	num, _ := strconv.Atoi(string(tok))
	return num
}
