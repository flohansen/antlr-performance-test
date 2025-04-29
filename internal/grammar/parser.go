package grammar

import (
	"fmt"
	"strconv"
)

type AstNode interface {
	String() string
}

type Identifier struct {
	Name string
}

func (id Identifier) String() string {
	return id.Name
}

type NumberConstant struct {
	Value float64
}

func (nc NumberConstant) String() string {
	return fmt.Sprintf("%.2f", nc.Value)
}

type BinaryExpression struct {
	Left     AstNode
	Operator BinaryOperator
	Right    AstNode
}

func (be BinaryExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", be.Left, be.Operator, be.Right)
}

type BinaryOperator int

const (
	BinaryOperatorAdd BinaryOperator = iota
	BinaryOperatorSub
	BinaryOperatorMul
	BinaryOperatorDiv
)

func (bo BinaryOperator) String() string {
	switch bo {
	case BinaryOperatorAdd:
		return "+"
	case BinaryOperatorSub:
		return "-"
	case BinaryOperatorMul:
		return "*"
	case BinaryOperatorDiv:
		return "/"
	default:
		panic(fmt.Sprintf("unknown binary operator %#v", bo))
	}
}

func addOperatorFrom(token Token) (BinaryOperator, bool) {
	switch token.Type {
	case TokenTypeAdd:
		return BinaryOperatorAdd, true
	case TokenTypeSub:
		return BinaryOperatorSub, true
	default:
		return BinaryOperatorAdd, false
	}
}

func mulOperatorFrom(token Token) (BinaryOperator, bool) {
	switch token.Type {
	case TokenTypeMul:
		return BinaryOperatorMul, true
	case TokenTypeDiv:
		return BinaryOperatorDiv, true
	default:
		return BinaryOperatorMul, false
	}
}

type Parser struct {
	lexer *Lexer
}

func NewParser(lexer *Lexer) *Parser {
	p := &Parser{
		lexer: lexer,
	}

	p.consume()
	return p
}

func (p *Parser) ParseTerm() (AstNode, error) {
	return p.ParseAdd()
}

func (p *Parser) ParseAdd() (AstNode, error) {
	left, err := p.ParseMul()
	if err != nil {
		return nil, err
	}

	for {
		operator, ok := addOperatorFrom(p.peek())
		if !ok {
			break
		}
		p.consume()

		right, err := p.ParseMul()
		if err != nil {
			return nil, err
		}

		left = &BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left, nil
}

func (p *Parser) ParseMul() (AstNode, error) {
	left, err := p.ParsePrimary()
	if err != nil {
		return nil, err
	}

	for {
		operator, ok := mulOperatorFrom(p.peek())
		if !ok {
			break
		}
		p.consume()

		right, err := p.ParsePrimary()
		if err != nil {
			return nil, err
		}

		left = &BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left, nil
}

func (p *Parser) ParsePrimary() (AstNode, error) {
	token := p.consume()

	switch token.Type {
	case TokenTypeID:
		return &Identifier{Name: token.Value}, nil
	case TokenTypeNumber:
		value, err := strconv.ParseFloat(token.Value, 64)
		if err != nil {
			return nil, p.newError("could not parse token value as float: %s", err)
		}

		return &NumberConstant{Value: value}, nil
	case TokenTypeLParen:
		term, err := p.ParseTerm()
		if err != nil {
			return nil, err
		}

		if _, err := p.expect(TokenTypeRParen); err != nil {
			return nil, err
		}

		return term, nil
	default:
		return nil, p.newError("unexpected token %s", token.Type)
	}
}

func (p *Parser) peek() Token {
	return p.lexer.Token
}

func (p *Parser) consume() Token {
	token := p.peek()
	p.lexer.Next()
	return token
}

func (p *Parser) expect(expected TokenType) (Token, error) {
	if p.lexer.Token.Type != expected {
		return Token{}, p.newError("expected %s, got %s", expected, p.lexer.Token.Type)
	}

	token := p.consume()
	return token, nil
}

func (p *Parser) newError(format string, a ...any) error {
	msg := fmt.Sprintf(format, a...)
	return fmt.Errorf("%s (line %d:%d)", msg, p.lexer.Line, p.lexer.Column)
}
