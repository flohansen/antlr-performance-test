# Evaluating ANTLR4: A Performance Comparison of Generated Code vs. Handwritten Lexer and Parser

**Motivation:**
When I embarked on this experiment, I was immersed in the development of a
domain-specific language (DSL) that required rapid processing capabilities. The
primary focus of this endeavor was performance, particularly in the evaluation
phase.

As I collaborated with engineers who had limited experience in systems
programming, I began to consider the advantages of utilizing a framework that
would enable my team members to maintain the code more effectively. This led me
to discover ANTLR4, which appeared to be an ideal solution due to its
declarative approach to defining grammar through g4 files.

In the past, I had hesitated to use code generators, but I recognized that, in
this instance, leveraging one could be highly beneficial. Thus, the motivation
for this project emerged: I aim to measure the performance of ANTLR4 against
that of a handwritten parser, providing valuable insights into the trade-offs
between the two approaches.

## Test Results
### Golang
```
Benchmark Results 'Handwritten' (10000 iterations):
    min: 275ns, max: 7.974µs, avg: 362ns

Benchmark Results 'ANTLR4' (10000 iterations):
    min: 4.786µs, max: 614.346µs, avg: 6.233µs
```

## Test Setup
For this performance evaluation, we will use a straightforward mathematical
expression to measure the speed of both the ANTLR4-generated code and the
handwritten lexer and parser.

```
((x+1) + (x+2)) * (x+3) - (x+4) / (x+5)
```

This expression includes addition, multiplication, and division, along with
parentheses to indicate the order of operations. It provides a good mix of
operations for testing.

Next, we will define the grammar that will be used to create both a handwritten
lexer and parser, as well as to generate them using ANTLR4.

```
lexer grammar GrammarLexer;

LPAREN: '(' ;
RPAREN: ')' ;
ADD: '+' ;
SUB: '-' ;
MUL: '*' ;
DIV: '/' ;

ID: [a-zA-Z][a-zA-Z0-9_]*;
NUMBER: [0-9]+('.'[0-9]*)?;
WS: [ \r\t\n\f]+ -> skip;
```

```
parser grammar GrammarParser;
options { tokenVocab = GrammarLexer; }

term: add;

add: mul (addOp mul)*;
addOp: ADD | SUB;

mul: primary (mulOp primary)*;
mulOp: MUL | DIV;

primary
    : ID
    | NUMBER
    | LPAREN term RPAREN
    ;
```

The code for the handwritten lexer and parser can be found in the following
files:

- Lexer: `internal/grammar/lexer.go`
- Parser: `internal/grammar/parser.go`

The generated code can be found here:

- Lexer: `generated/go/parser/grammar_lexer.go`
- Parser: `generated/go/parser/grammar_parser.go`

And finally the evaluator used to evaluate the output (AST) of both the
handwritten and the generated parser:

- Evaluator (Visitor): `internal/grammar/eval.go`

## Test Execution

To run the test you simply just have to run

    go run cmd/test/main.go

## Conclusion
It is evident that the handwritten approach significantly outperforms the
generated one in terms of speed. In this relatively simple experiment, using
ANTLR4 to generate the lexer and parser results in evaluations that are
approximately 20 times slower, which is quite substantial. This performance
difference is particularly critical when speed is a non-functional requirement.
The reason for this performance difference may stem from the abstraction levels
employed by ANTLR4, which are designed to support robust code generation across
a wide range of use cases.

On the other hand, while the handwritten implementation offers better
performance, the development time required to create a lexer and parser from
scratch can be considerably higher. This challenge becomes even more pronounced
as the complexity of the language increases, making the manual approach less
feasible for more intricate projects.

## Contributing

Feel free to contribute to this experiment. E.g. by adding more test scenarios
using different languages or by using input expressions with different
complexity. Thank you for your interest!
