// Code generated from GrammarParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GrammarParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by GrammarParser#add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by GrammarParser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by GrammarParser#mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by GrammarParser#mulOp.
	VisitMulOp(ctx *MulOpContext) interface{}

	// Visit a parse tree produced by GrammarParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}
}
