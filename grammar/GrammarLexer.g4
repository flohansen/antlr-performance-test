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
