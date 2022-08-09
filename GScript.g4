grammar GScript;

prog
    : blockStatements
    ;

block
    : '{' blockStatements '}'
    ;

blockStatements
    : blockStatement* #BlockStms
    ;

blockStatement
    : variableDeclarators #BlockVarDeclar
    | statement # BlockStm
   // | localTypeDeclaration
    | functionDeclaration #BlockFunc
    //| classDeclaration
    ;


parse
 : expr_list+=expr+ EOF
 ;

expr
    : primary #PrimaryExpr
    | lhs=expr postfix=('++' | '--') #PostfixExpr
    | prefix=('~'|'!') rhs=expr #NotExpr
    | lhs=expr bop=( MULT | DIV ) rhs=expr #MultDivExpr
    | lhs=expr bop=MOD rhs=expr            #ModExpr
    | lhs=expr bop=( PLUS | SUB ) rhs=expr #PlusSubExpr
    | expr bop=(LE | GE | GT | LT ) expr # GLe
    | expr bop=(EQUAL | NOTEQUAL) expr # EqualOrNot
    ;

primary
    : '(' expr ')' #ExprPrimary
    //| THIS
    //| SUPER
    | literal #LiterPrimary
    | IDENTIFIER #IdentifierPrimay
    // | typeTypeOrVoid '.' CLASS
    ;


statement
    : blockLabel=block #BlockLabel
    | IF parExpression statement (ELSE statement)? #IfElse
    | FOR '(' forControl ')' statement #For
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

functionDeclaration
    : typeTypeOrVoid? IDENTIFIER formalParameters ('[' ']')*
      (THROWS qualifiedNameList)?
      functionBody
    ;

functionBody
    : block
    ;

qualifiedNameList
    : qualifiedName (',' qualifiedName)*
    ;

formalParameters
    : '(' formalParameterList? ')'
    ;

formalParameterList
    : formalParameter (',' formalParameter)* (',' lastFormalParameter)?
    | lastFormalParameter
    ;

formalParameter
    : variableModifier* typeType variableDeclaratorId
    ;

lastFormalParameter
    : variableModifier* typeType '...' variableDeclaratorId
    ;

qualifiedName
    : IDENTIFIER ('.' IDENTIFIER)*
    ;

variableModifier
// todo final require?
    : FINAL
    //| annotation
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
    : arrayInitializer
    | expr
    ;

arrayInitializer
    : '{' (variableInitializer (',' variableInitializer)* (',')? )? '}'
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
    : INT
    | STRING
    | FLOAT
    | BOOLEAN
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

// Keywords
FINAL:              'final';
THROWS:             'throws';
INT:                'int';
STRING:             'string';
FLOAT:              'float';
BOOLEAN:            'bool';


// Separators

LPAREN:             '(';
RPAREN:             ')';
LBRACE:             '{';
RBRACE:             '}';
LBRACK:             '[';
RBRACK:             ']';

ASSIGN:             '=';
GT:                 '>';
LT:                 '<';
BANG:               '!';
TILDE:              '~';
INC:                '++';
DEC:                '--';
MULT : '*';
DIV  : '/';
PLUS : '+';
SUB  : '-';
MOD  : '%';
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