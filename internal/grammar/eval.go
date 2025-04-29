package grammar

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/flohansen/antlr-performance-test/generated/go/parser"
)

type VarStore map[string]any

func (vs VarStore) Set(key string, value any) {
	vs[key] = value
}

func (vs VarStore) Get(key string) any {
	return vs[key]
}

type Evaluator struct {
	parser.BaseGrammarParserVisitor
	Vars VarStore
}

func NewEvaluator() *Evaluator {
	return &Evaluator{
		Vars: make(VarStore),
	}
}

func (e *Evaluator) Eval(tree AstNode) any {
	switch node := tree.(type) {
	case *Identifier:
		return e.Vars.Get(node.Name)
	case *BinaryExpression:
		left := e.Eval(node.Left).(float64)
		right := e.Eval(node.Right).(float64)

		switch node.Operator {
		case BinaryOperatorAdd:
			return left + right
		case BinaryOperatorSub:
			return left - right
		case BinaryOperatorMul:
			return left * right
		case BinaryOperatorDiv:
			return left / right
		default:
			return nil
		}
	case *NumberConstant:
		return node.Value
	default:
		return nil
	}
}

func (e *Evaluator) Visit(tree antlr.ParseTree) any {
	return tree.Accept(e)
}

func (e *Evaluator) VisitTerm(ctx *parser.TermContext) any {
	return ctx.Add().Accept(e)
}

func (e *Evaluator) VisitAdd(ctx *parser.AddContext) any {
	left := ctx.Mul(0).Accept(e).(AstNode)

	for i := 1; i < len(ctx.AllMul()); i++ {
		operator := ctx.AddOp(i - 1)
		right := ctx.Mul(i).Accept(e).(AstNode)

		if operator.ADD() != nil {
			left = &BinaryExpression{
				Left:     left,
				Operator: BinaryOperatorAdd,
				Right:    right,
			}
		} else if operator.SUB() != nil {
			left = &BinaryExpression{
				Left:     left,
				Operator: BinaryOperatorSub,
				Right:    right,
			}
		}
	}

	return left
}

func (e *Evaluator) VisitMul(ctx *parser.MulContext) any {
	left := ctx.Primary(0).Accept(e).(AstNode)

	for i := 1; i < len(ctx.AllPrimary()); i++ {
		operator := ctx.MulOp(i - 1)
		right := ctx.Primary(i).Accept(e).(AstNode)

		if operator.MUL() != nil {
			left = &BinaryExpression{
				Left:     left,
				Operator: BinaryOperatorMul,
				Right:    right,
			}
		} else if operator.DIV() != nil {
			left = &BinaryExpression{
				Left:     left,
				Operator: BinaryOperatorDiv,
				Right:    right,
			}
		}
	}

	return left
}

func (e *Evaluator) VisitPrimary(ctx *parser.PrimaryContext) any {
	if ctx.ID() != nil {
		return &Identifier{Name: ctx.ID().GetText()}
	} else if ctx.NUMBER() != nil {
		value, err := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
		if err != nil {
			panic(err)
		}

		return &NumberConstant{Value: value}
	} else if ctx.LPAREN() != nil {
		return ctx.Term().Accept(e)
	}

	return nil
}
