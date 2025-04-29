package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/flohansen/antlr-performance-test/generated/go/parser"
	"github.com/flohansen/antlr-performance-test/internal/bench"
	"github.com/flohansen/antlr-performance-test/internal/grammar"
)

func main() {
	content := []byte("((x+1)+(x+2))*(x+3)-(x+4)/(x+5)")
	evaluator := grammar.NewEvaluator()
	evaluator.Vars.Set("x", 1.0)

	{
		lexer := grammar.NewLexer(content)
		parser := grammar.NewParser(lexer)

		tree, err := parser.ParseTerm()
		if err != nil {
			panic(err)
		}

		fmt.Println(bench.Run("Handwritten", func() {
			evaluator.Eval(tree)
		}, 10000))
	}

	fmt.Println("")

	{
		input := antlr.NewInputStream(string(content))
		lexer := parser.NewGrammarLexer(input)
		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewGrammarParser(tokenStream)

		tree := p.Term()

		fmt.Println(bench.Run("ANTLR4", func() {
			evaluator.Visit(tree)
		}, 10000))
	}
}
