// Code generated from GrammarParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GrammarParser

import "github.com/antlr4-go/antlr/v4"

// BaseGrammarParserListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarParserListener struct{}

var _ GrammarParserListener = &BaseGrammarParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseGrammarParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseGrammarParserListener) ExitTerm(ctx *TermContext) {}

// EnterAdd is called when production add is entered.
func (s *BaseGrammarParserListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BaseGrammarParserListener) ExitAdd(ctx *AddContext) {}

// EnterAddOp is called when production addOp is entered.
func (s *BaseGrammarParserListener) EnterAddOp(ctx *AddOpContext) {}

// ExitAddOp is called when production addOp is exited.
func (s *BaseGrammarParserListener) ExitAddOp(ctx *AddOpContext) {}

// EnterMul is called when production mul is entered.
func (s *BaseGrammarParserListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production mul is exited.
func (s *BaseGrammarParserListener) ExitMul(ctx *MulContext) {}

// EnterMulOp is called when production mulOp is entered.
func (s *BaseGrammarParserListener) EnterMulOp(ctx *MulOpContext) {}

// ExitMulOp is called when production mulOp is exited.
func (s *BaseGrammarParserListener) ExitMulOp(ctx *MulOpContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseGrammarParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseGrammarParserListener) ExitPrimary(ctx *PrimaryContext) {}
