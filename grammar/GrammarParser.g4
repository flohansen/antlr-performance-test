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
