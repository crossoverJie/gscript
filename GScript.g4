grammar GScript;

prog
    : blockStatements
    ;

blockStatements
    : blockStatement* #BlockStms
    ;

blockStatement
    : variableDeclarators #BlockVarDeclar
    | statement # BlockStm
   // | localTypeDeclaration
    //| functionDeclaration
    //| classDeclaration
    ;


parse
 : expr_list+=expr+ EOF
 ;

expr
    : '(' expr ')'                        #NestedExpr
    | liter=literal #Liter
    | SUB expr                            #UnaryExpr
    | lhs=expr bop=( MULT | DIV ) rhs=expr #MultDivExpr
    | lhs=expr bop=MOD rhs=expr            #ModExpr
    | lhs=expr bop=( PLUS | SUB ) rhs=expr #PlusSubExpr
    | expr bop=(LE | GE | GT | LT ) expr # GLe
    | expr bop=(EQUAL | NOTEQUAL) expr # EqualOrNot
    ;

block
    : '{' blockStatements '}'
    ;

statement
    : blockLabel=block #blockLabel
    | IF parExpression statement (ELSE statement)? #IfElse
    //| FOR '(' forControl ')' statement #For
    | RETURN expr?  #Return
    | statementExpression=expr #StmExpr
    ;

forControl
    : forInit? ';' expr? ';' forUpdate=expressionList?
    ;

forInit
    : variableDeclarators
    | expressionList
    ;
variableDeclarators
    : typeType variableDeclarator (',' variableDeclarator)*
    ;

variableDeclarator
    : variableDeclaratorId ('=' variableInitializer)?
    ;

variableDeclaratorId
    : IDENTIFIER ('[' ']')*
    ;

variableInitializer
    : expr
    ;

literal
//    : integerLiteral #Int
    : DECIMAL_LITERAL #Int
//    | floatLiteral #Float
    | FLOAT_LITERAL #Float
//    | CHAR_LITERAL
    | STRING_LITERAL #String
    | BOOL_LITERAL # Bool
    | NULL_LITERAL # Null
    ;

integerLiteral
    : DECIMAL_LITERAL
//    | HEX_LITERAL
//    | OCT_LITERAL
//    | BINARY_LITERAL
    ;
floatLiteral
    : FLOAT_LITERAL
//    | HEX_FLOAT_LITERAL
    ;

typeType
    : (classOrInterfaceType| functionType | primitiveType) ('[' ']')*
    ;

primitiveType
    : NUMBER
    | STRING
    ;

functionType
    : FUNCTION typeTypeOrVoid '(' typeList? ')'
    ;

typeList
    : typeType (',' typeType)*
    ;

typeTypeOrVoid
    : typeType
    | VOID
    ;

classOrInterfaceType
    : IDENTIFIER ('.' IDENTIFIER)*
    //: IDENTIFIER
    ;

expressionList
    : expr (',' expr)*
    ;

parExpression
    : '(' expr ')'
    ;

MULT : '*';
DIV  : '/';
PLUS : '+';
SUB  : '-';
MOD  : '%';
ASSIGN:             '=';
GT:                 '>';
LT:                 '<';
EQUAL:              '==';
LE:                 '<=';
GE:                 '>=';
NOTEQUAL:           '!=';

FUNCTION:           'func';
VOID:               'void';
FOR:                'for';
IF:                 'if';
ELSE:               'else';
RETURN:             'return';

BOOL_LITERAL:       'true'
            |       'false'
            ;

//CHAR_LITERAL:       '\'' (~['\\\r\n] | EscapeSequence) '\'';

STRING_LITERAL:     '"' (~["\\\r\n] | EscapeSequence)* '"';

NULL_LITERAL:       'null';

DECIMAL_LITERAL:    ('0' | [1-9] (Digits? | '_'+ Digits)) [lL]?;
FLOAT_LITERAL:      (Digits '.' Digits? | '.' Digits) ExponentPart? [fFdD]?
             |       Digits (ExponentPart [fFdD]? | [fFdD])
             ;
IDENTIFIER:         Letter LetterOrDigit*;
fragment Digits
    : [0-9] ([0-9_]* [0-9])?
    ;

NUMBER
 : ( D* '.' )? D+
 ;

STRING:             'string';

SPACES
 : [ \t\r\n] -> skip
 ;

fragment D : [0-9];

fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;

fragment LetterOrDigit
    : Letter
    | [0-9]
    ;

fragment ExponentPart
    : [eE] [+-]? Digits
    ;

fragment EscapeSequence
    : '\\' [btnfr"'\\]
    | '\\' ([0-3]? [0-7])? [0-7]
    | '\\' 'u'+ HexDigit HexDigit HexDigit HexDigit
    ;

fragment HexDigit
    : [0-9a-fA-F]
    ;