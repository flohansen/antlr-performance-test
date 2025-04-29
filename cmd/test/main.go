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
		fmt.Println(bench.Run("Handwritten", func(b *bench.B) {
			var tree grammar.AstNode

			b.Run("parsing", func(b *bench.B) {
				lexer := grammar.NewLexer(content)
				parser := grammar.NewParser(lexer)

				var err error
				tree, err = parser.ParseTerm()
				if err != nil {
					panic(err)
				}
			})

			b.Run("evaluation", func(b *bench.B) {
				evaluator.Eval(tree)
			})
		}, 10000))
	}

	fmt.Println()

	{
		fmt.Println(bench.Run("ANTLR4", func(b *bench.B) {
			var ast grammar.AstNode

			b.Run("parsing", func(b *bench.B) {
				input := antlr.NewInputStream(string(content))
				lexer := parser.NewGrammarLexer(input)
				tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
				p := parser.NewGrammarParser(tokenStream)

				tree := p.Term()
				ast = evaluator.Visit(tree).(grammar.AstNode)
			})

			b.Run("evaluation", func(b *bench.B) {
				evaluator.Eval(ast)
			})
		}, 10000))
	}
}
