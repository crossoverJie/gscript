grammar GScript;


functionDeclaration
    : typeTypeOrVoid? IDENTIFIER formalParameters ('[' ']')*
      (THROWS qualifiedNameList)?
      functionBody
    ;

functionBody
    : block
    ;

typeTypeOrVoid
    : typeType
    | VOID
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

variableModifier
// todo final require?
    : FINAL
    //| annotation
    ;

qualifiedName
    : IDENTIFIER ('.' IDENTIFIER)*
    ;


fieldDeclaration
    //: typeType variableDeclarators ';'
    : variableDeclarators ';'
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


classOrInterfaceType
    : IDENTIFIER ('.' IDENTIFIER)*
    //: IDENTIFIER
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

statement
    : blockLabel=block #StmBlockLabel
    | IF parExpression statement (ELSE statement)? #StmIfElse
    | FOR '(' forControl ')' statement #StmFor
    | RETURN expr?  #StmReturn
    | statementExpression=expr #StmExpr
    ;

forControl
    : forInit? ';' expr? ';' forUpdate=expressionList?
    ;

forInit
    : variableDeclarators
    | expressionList
    ;

parExpression
    : '(' expr ')'
    ;


expressionList
    : expr (',' expr)*
    ;



functionCall
    : IDENTIFIER '(' expressionList? ')'
    | THIS '(' expressionList? ')'
    | SUPER '(' expressionList? ')'
    ;


expr
    : primary #PrimaryExpr
    | expr bop='.'
      ( IDENTIFIER
      | functionCall
      | expr
      ) #DotExpr
    | functionCall #FuncCallExpr
    | lhs=expr postfix=('++' | '--') #PostfixExpr
    | prefix=('~'|'!') rhs=expr #NotExpr
    | lhs=expr bop=( MULT | DIV ) rhs=expr #MultDivExpr
    | lhs=expr bop=MOD rhs=expr            #ModExpr
    | lhs=expr bop=( PLUS | SUB ) rhs=expr #PlusSubExpr
    | expr bop=(LE | GE | GT | LT ) expr # GLeExpr
    | expr bop=(EQUAL | NOTEQUAL) expr # EqualOrNot
    // 表明结合性是右结合的，内部原理使用循环代替递归。
    | <assoc=right> lhs=expr
      bop=('=' | '+=' | '-=' | '*=')
      rhs=expr  #AssignExpr
    ;

primary
    : '(' expr ')' #ExprPrimary
    //| THIS
    //| SUPER
    | literal #LiterPrimary
    | IDENTIFIER #IdentifierPrimary
    // | typeTypeOrVoid '.' CLASS
    ;






typeList
    : typeType (',' typeType)*
    ;



typeType
    : (classOrInterfaceType| functionType | primitiveType) ('[' ']')*
    ;


functionType
    : FUNCTION typeTypeOrVoid '(' typeList? ')'
    ;

primitiveType
    : INT
    | STRING
    | FLOAT
    | BOOLEAN
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






parse
 : expr_list+=expr+ EOF
 ;

// Keywords
FINAL:              'final';
THROWS:             'throws';
INT:                'int';
STRING:             'string';
FLOAT:              'float';
BOOLEAN:            'bool';
SUPER:              'super';
SWITCH:             'switch';
SYNCHRONIZED:       'synchronized';
THIS:               'this';


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

ADD_ASSIGN:         '+=';
SUB_ASSIGN:         '-=';
MUL_ASSIGN:         '*=';
DIV_ASSIGN:         '/=';
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