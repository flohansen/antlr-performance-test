// Code generated from GrammarParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // GrammarParser

import "github.com/antlr4-go/antlr/v4"

// GrammarParserListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarParserListener interface {
	antlr.ParseTreeListener

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterAddOp is called when entering the addOp production.
	EnterAddOp(c *AddOpContext)

	// EnterMul is called when entering the mul production.
	EnterMul(c *MulContext)

	// EnterMulOp is called when entering the mulOp production.
	EnterMulOp(c *MulOpContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitAddOp is called when exiting the addOp production.
	ExitAddOp(c *AddOpContext)

	// ExitMul is called when exiting the mul production.
	ExitMul(c *MulContext)

	// ExitMulOp is called when exiting the mulOp production.
	ExitMulOp(c *MulOpContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)
}
