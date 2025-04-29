package grammar

import (
	"fmt"
)

type TokenType int

const (
	TokenTypeEOF TokenType = iota
	TokenTypeID
	TokenTypeNumber
	TokenTypeLParen
	TokenTypeRParen
	TokenTypeAdd
	TokenTypeSub
	TokenTypeMul
	TokenTypeDiv
)

func (tt TokenType) String() string {
	switch tt {
	case TokenTypeEOF:
		return "EOF"
	case TokenTypeID:
		return "ID"
	case TokenTypeNumber:
		return "Number"
	case TokenTypeLParen:
		return "("
	case TokenTypeRParen:
		return ")"
	case TokenTypeAdd:
		return "+"
	case TokenTypeSub:
		return "-"
	case TokenTypeMul:
		return "*"
	case TokenTypeDiv:
		return "/"
	default:
		panic(fmt.Sprintf("unexpected token type: %#v", tt))
	}
}

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input  []byte
	cursor int
	Token  Token
	Line   int
	Column int
}

func NewLexer(input []byte) *Lexer {
	return &Lexer{
		input:  input,
		cursor: 0,
		Token:  Token{Type: TokenTypeEOF},
		Line:   1,
		Column: 0,
	}
}

func (l *Lexer) Next() bool {
	c, err := l.next()
	if err != nil {
		l.Token = Token{Type: TokenTypeEOF}
		return false
	}

	switch true {
	case isWhitespace(c):
		return l.Next()
	case isAlpha(c):
		l.tokenizeID(c)
	case isNum(c):
		l.tokenizeNumber(c)
	case c == '(':
		l.Token = Token{Type: TokenTypeLParen}
	case c == ')':
		l.Token = Token{Type: TokenTypeRParen}
	case c == '+':
		l.Token = Token{Type: TokenTypeAdd}
	case c == '-':
		l.Token = Token{Type: TokenTypeSub}
	case c == '*':
		l.Token = Token{Type: TokenTypeMul}
	case c == '/':
		l.Token = Token{Type: TokenTypeDiv}
	}

	return true
}

func (l *Lexer) tokenizeID(c byte) {
	l.Token = Token{Type: TokenTypeID, Value: string(c)}

	for {
		c, err := l.peek()
		if err != nil {
			break
		}
		if !isAlphaNum(c) && c != '_' {
			break
		}

		l.next()
		l.Token.Value += string(c)
	}
}

func (l *Lexer) tokenizeNumber(c byte) {
	l.Token = Token{Type: TokenTypeNumber, Value: string(c)}
	l.readNumber()

	if c, err := l.peek(); err == nil && c == '.' {
		l.next()
		l.Token.Value += "."
		l.readNumber()
	}
}

func (l *Lexer) readNumber() {
	for {
		c, err := l.peek()
		if err != nil {
			break
		}
		if !isNum(c) {
			break
		}

		l.next()
		l.Token.Value += string(c)
	}
}

func (l *Lexer) peek() (byte, error) {
	if l.cursor >= len(l.input) {
		return 0, fmt.Errorf("end of file (line %d:%d)", l.Line, l.Column)
	}

	return l.input[l.cursor], nil
}

func (l *Lexer) next() (byte, error) {
	c, err := l.peek()
	if err != nil {
		return 0, err
	}

	if c == '\n' {
		l.Line++
		l.Column = 1
	} else {
		l.Column++
	}

	l.cursor++
	return c, nil
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == 'r' || c == '\f'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlphaNum(c byte) bool {
	return isAlpha(c) || isNum(c)
}
