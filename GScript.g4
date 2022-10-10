grammar GScript;


classDeclaration
    : CLASS IDENTIFIER
      (EXTENDS typeType)?
      (IMPLEMENTS typeList)?
      classBody
    ;

classBody
    : '{' classBodyDeclaration* '}'
    ;

// interfaceBody
//     : '{' interfaceBodyDeclaration* '}'
//     ;

classBodyDeclaration
    : ';'
    //| STATIC? block
    | memberDeclaration
    ;

memberDeclaration
    : functionDeclaration
//    | genericFunctionDeclaration
    | fieldDeclaration
    // | constructorDeclaration
    // | genericConstructorDeclaration
//     | interfaceDeclaration
    // | annotationTypeDeclaration
     //| classDeclaration
    // | enumDeclaration
    ;
functionDeclaration
    : typeTypeOrVoid? IDENTIFIER operatorOverloading? formalParameters ('[' ']')*
      (THROWS qualifiedNameList)?
      functionBody
    ;

operatorOverloading
    : PLUS
    | SUB
    | MULT
    | DIV
    | EQUAL
    | NOTEQUAL
    | LE
    | GE
    | GT
    | LT
    ;

functionBody
    : block
    | ';'
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
    : ('[' DECIMAL_LITERAL ']')? '{' (variableInitializer (',' variableInitializer)* (',')? )? '}'
    ;


classOrInterfaceType
    : IDENTIFIER ('.' IDENTIFIER)*
    //: IDENTIFIER
    ;


literal
//    : integerLiteral #Int
    : DECIMAL_LITERAL
//    | floatLiteral #Float
    | FLOAT_LITERAL
//    | CHAR_LITERAL
    | string_
    | BOOL_LITERAL
    | NULL_LITERAL
    | Nil
    ;

string_
    :RAW_STRING_LIT
    |STRING_LITERAL
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
    : variableDeclarators ';' #BlockVarDeclar
    | statement # BlockStm
   // | localTypeDeclaration
    | functionDeclaration #BlockFunc
    | classDeclaration #BlockClassDeclar
    ;

statement
    : blockLabel=block #StmBlockLabel
    | IF parExpression statement (ELSE statement)? #StmIfElse
    | FOR '(' forControl ')' statement #StmFor
    | FOR parExpression statement #StmWhile
    | RETURN expr?  ';'#StmReturn
    | BREAK IDENTIFIER? ';' #StmBreak
    | CONTINUE IDENTIFIER? ';' # StmContinue
    | statementExpression=expr ';'#StmExpr
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
    : primary
    | expr bop='.'
      ( IDENTIFIER
      | functionCall
      )
    | array=expr '[' index=expr ']'
    | functionCall
    | lhs=expr postfix=('++' | '--')
    | prefix=('-'|'+') rhs=expr
    | prefix=('~'|'!') rhs=expr
    | lhs=expr bop=( MULT | DIV ) rhs=expr
    | lhs=expr bop=MOD rhs=expr
    | lhs=expr bop=( PLUS | SUB ) rhs=expr
    | lhs=expr bop=(LE | GE | GT | LT ) rhs=expr
    | lhs=expr bop=(EQUAL | NOTEQUAL) rhs=expr
    | lhs=expr bop='&&' rhs=expr
    | lhs=expr bop='||' rhs=expr
    // 表明结合性是右结合的，内部原理使用循环代替递归。
    | <assoc=right> lhs=expr
      bop=('=' | '+=' | '-=' | '*=')
      rhs=expr
    ;

primary
    : '(' expr ')'
    //| THIS
    //| SUPER
    | literal
    | IDENTIFIER
    // | typeTypeOrVoid '.' CLASS
    ;






typeList
    : typeType (',' typeType)*
    ;



typeType
    : (classOrInterfaceType| functionType | primitiveType) ('[' ']')*
    ;


functionType
    : FUNCTION typeTypeOrVoid? '(' typeList? ')'
    ;

primitiveType
    : INT
    | STRING
    | FLOAT
    | BOOLEAN
    | BYTE
    | ANY
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
CLASS:              'class';
EXTENDS:            'extends';
IMPLEMENTS:         'implements';
FINAL:              'final';
THROWS:             'throws';
INT:                'int';
STRING:             'string';
FLOAT:              'float';
BOOLEAN:            'bool';
BYTE:               'byte';
SUPER:              'super';
SWITCH:             'switch';
THIS:               'this';
Nil:                'nil';
ANY:                'any';


// Separators

LPAREN:             '(';
RPAREN:             ')';
LBRACE:             '{';
RBRACE:             '}';
LBRACK:             '[';
RBRACK:             ']';
DOT:                '.';

ASSIGN:             '=';
GT:                 '>';
LT:                 '<';
BANG:               '!';
TILDE:              '~';
AND:                '&&';
OR:                 '||';
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
BREAK:              'break';
CONTINUE:           'continue';

BOOL_LITERAL:       'true'
            |       'false'
            ;



//CHAR_LITERAL:       '\'' (~['\\\r\n] | EscapeSequence) '\'';

STRING_LITERAL:     '"' (~["\\\r\n] | EscapeSequence)* '"';
RAW_STRING_LIT         : '^' ~'^'*                      '^';


NULL_LITERAL:       'null';

DECIMAL_LITERAL:    ('0' | [1-9] (Digits? | '_'+ Digits)) [lL]?;
FLOAT_LITERAL:      (Digits '.' Digits? | '.' Digits) ExponentPart? [fFdD]?
             |       Digits (ExponentPart [fFdD]? | [fFdD])
             ;

ELLIPSIS:           '...';

// Whitespace and comments

WS:                 [ \t\r\n\u000C]+ -> channel(HIDDEN);
COMMENT:            '/*' .*? '*/'    -> channel(HIDDEN);
LINE_COMMENT:       '//' ~[\r\n]*    -> channel(HIDDEN);
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