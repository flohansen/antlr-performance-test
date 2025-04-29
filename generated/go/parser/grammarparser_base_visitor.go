// Code generated from GrammarParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GrammarParser

import "github.com/antlr4-go/antlr/v4"

type BaseGrammarParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGrammarParserVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarParserVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarParserVisitor) VisitAddOp(ctx *AddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarParserVisitor) VisitMul(ctx *MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarParserVisitor) VisitMulOp(ctx *MulOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGrammarParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}
