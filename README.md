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
Benchmark Result 'Handwritten':
  (total): min: 1.93µs, max: 192.278µs, avg: 2.245µs:
  (parsing): min: 1.365µs, max: 191.271µs, avg: 1.637µs
  (evaluation): min: 315ns, max: 114.191µs, avg: 377ns

Benchmark Result 'ANTLR4':
  (total): min: 25.758µs, max: 501.241µs, avg: 33.168µs:
  (parsing): min: 25.136µs, max: 498.293µs, avg: 32.452µs
  (evaluation): min: 324ns, max: 18.754µs, avg: 425ns
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
ANTLR4 to generate the lexer and parser results in parsing which is
approximately 20 times slower, which is quite substantial. This performance
difference is particularly critical when speed is a non-functional requirement.
The reason for this performance difference may stem from the abstraction levels
employed by ANTLR4, which are designed to support robust code generation across
a wide range of use cases. The evaluation speed is approximately the same.

On the other hand, while the handwritten implementation offers better
performance, the development time required to create a lexer and parser from
scratch can be considerably higher. This challenge becomes even more pronounced
as the complexity of the language increases, making the manual approach less
feasible for more intricate projects.

For applications where the parsing speed does not really matter, using ANTLR4
seems to be a good alternative due to the better maintainability and
development speed.

## Contributing

Feel free to contribute to this experiment. E.g. by adding more test scenarios
using different languages or by using input expressions with different
complexity. Thank you for your interest!
