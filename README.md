# golox

syntax
```
program    -> declaration* EOF ;
declaration -> classDecl | varDecl | funDecl | stmt ;
varDecl    -> "var" IDENTIFIER ( "=" expression )? ";" ;
funDecl    -> "fun" function ;
classDecl  -> "class" IDENTIFIER  ( "<" IDENTIFIER )? "{" (function|property)* "}" ;
function   -> IDENTIFIER "(" parameters? ")" block ;
property   -> IDENTIFIER block ;

stmt       -> exprStmt | ifStmt | printStmt | returnStmt | whileStmt | forStmt | breakStmt | continueStmt | block ;
breakStmt  -> "break" ";" ;
continueStmt -> "continue" ";" ;
returnStmt -> "return" expression? ";" ;
ifStmt     -> "if" "(" expression ")" statement ( "else " statement )? ;
whileStmt  -> "while" "(" expression ")" statement ;
forStmt    -> "for" "(" ( varDecl | exprStmt | ";" ) expression? ";" expression? ")" statement ;
block      -> "{" declaration* "}"
exprStmt   -> expression ";" ;
printStmt  -> "print" expression ";" ;

expression -> comma ;
comma      -> assignment ( "," assignment ) * ;
assignment -> (call "." )? IDENTIFIER "=" assignment | logic_or ;
logic_or   -> logic_and ( "or" logic_and )* ;
logic_and  -> ternary ( "and" ternary ) * ;
ternary    -> equality "?"  expression ":" expression ;
equality   -> comparison ( ( "!=" | "==") comparison )* ;
comparison -> addition ( ( ">" | ">=" | "<" | "<=") addition )*;
addition   -> multiplication ( ( "+" | "-" ) multiplication )*;
multiplication -> unary ( ( "/" | "*" ) unary )*;
unary      -> ( "!" | "-" ) unary | power ;
power      -> call ( "**" unary ) *
call       -> primary ( "(" arguments? ")" | "." IDENTIFIER )* ;
arguments  -> expression ( "," expression )* ;
primary    -> NUMBER | STRING | "false" | "true" | "nil" | "this" | "super" | "(" expression ")" | IDENTIFIER ;
```
