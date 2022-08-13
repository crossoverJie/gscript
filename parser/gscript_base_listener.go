// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGScriptListener is a complete listener for a parse tree produced by GScriptParser.
type BaseGScriptListener struct{}

var _ GScriptListener = &BaseGScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseGScriptListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseGScriptListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseGScriptListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseGScriptListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseGScriptListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseGScriptListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseGScriptListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseGScriptListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseGScriptListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseGScriptListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseGScriptListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseGScriptListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterTypeTypeOrVoid is called when production typeTypeOrVoid is entered.
func (s *BaseGScriptListener) EnterTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// ExitTypeTypeOrVoid is called when production typeTypeOrVoid is exited.
func (s *BaseGScriptListener) ExitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// EnterQualifiedNameList is called when production qualifiedNameList is entered.
func (s *BaseGScriptListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {}

// ExitQualifiedNameList is called when production qualifiedNameList is exited.
func (s *BaseGScriptListener) ExitQualifiedNameList(ctx *QualifiedNameListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseGScriptListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseGScriptListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseGScriptListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseGScriptListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseGScriptListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseGScriptListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BaseGScriptListener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BaseGScriptListener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BaseGScriptListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BaseGScriptListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseGScriptListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseGScriptListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseGScriptListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseGScriptListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseGScriptListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseGScriptListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseGScriptListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseGScriptListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseGScriptListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseGScriptListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseGScriptListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseGScriptListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseGScriptListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseGScriptListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseGScriptListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseGScriptListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseGScriptListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseGScriptListener) ExitInt(ctx *IntContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseGScriptListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseGScriptListener) ExitFloat(ctx *FloatContext) {}

// EnterString is called when production String is entered.
func (s *BaseGScriptListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseGScriptListener) ExitString(ctx *StringContext) {}

// EnterBool is called when production Bool is entered.
func (s *BaseGScriptListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production Bool is exited.
func (s *BaseGScriptListener) ExitBool(ctx *BoolContext) {}

// EnterNull is called when production Null is entered.
func (s *BaseGScriptListener) EnterNull(ctx *NullContext) {}

// ExitNull is called when production Null is exited.
func (s *BaseGScriptListener) ExitNull(ctx *NullContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseGScriptListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseGScriptListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGScriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGScriptListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStms is called when production BlockStms is entered.
func (s *BaseGScriptListener) EnterBlockStms(ctx *BlockStmsContext) {}

// ExitBlockStms is called when production BlockStms is exited.
func (s *BaseGScriptListener) ExitBlockStms(ctx *BlockStmsContext) {}

// EnterBlockVarDeclar is called when production BlockVarDeclar is entered.
func (s *BaseGScriptListener) EnterBlockVarDeclar(ctx *BlockVarDeclarContext) {}

// ExitBlockVarDeclar is called when production BlockVarDeclar is exited.
func (s *BaseGScriptListener) ExitBlockVarDeclar(ctx *BlockVarDeclarContext) {}

// EnterBlockStm is called when production BlockStm is entered.
func (s *BaseGScriptListener) EnterBlockStm(ctx *BlockStmContext) {}

// ExitBlockStm is called when production BlockStm is exited.
func (s *BaseGScriptListener) ExitBlockStm(ctx *BlockStmContext) {}

// EnterBlockFunc is called when production BlockFunc is entered.
func (s *BaseGScriptListener) EnterBlockFunc(ctx *BlockFuncContext) {}

// ExitBlockFunc is called when production BlockFunc is exited.
func (s *BaseGScriptListener) ExitBlockFunc(ctx *BlockFuncContext) {}

// EnterStmBlockLabel is called when production StmBlockLabel is entered.
func (s *BaseGScriptListener) EnterStmBlockLabel(ctx *StmBlockLabelContext) {}

// ExitStmBlockLabel is called when production StmBlockLabel is exited.
func (s *BaseGScriptListener) ExitStmBlockLabel(ctx *StmBlockLabelContext) {}

// EnterStmIfElse is called when production StmIfElse is entered.
func (s *BaseGScriptListener) EnterStmIfElse(ctx *StmIfElseContext) {}

// ExitStmIfElse is called when production StmIfElse is exited.
func (s *BaseGScriptListener) ExitStmIfElse(ctx *StmIfElseContext) {}

// EnterStmFor is called when production StmFor is entered.
func (s *BaseGScriptListener) EnterStmFor(ctx *StmForContext) {}

// ExitStmFor is called when production StmFor is exited.
func (s *BaseGScriptListener) ExitStmFor(ctx *StmForContext) {}

// EnterStmReturn is called when production StmReturn is entered.
func (s *BaseGScriptListener) EnterStmReturn(ctx *StmReturnContext) {}

// ExitStmReturn is called when production StmReturn is exited.
func (s *BaseGScriptListener) ExitStmReturn(ctx *StmReturnContext) {}

// EnterStmExpr is called when production StmExpr is entered.
func (s *BaseGScriptListener) EnterStmExpr(ctx *StmExprContext) {}

// ExitStmExpr is called when production StmExpr is exited.
func (s *BaseGScriptListener) ExitStmExpr(ctx *StmExprContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseGScriptListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseGScriptListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseGScriptListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseGScriptListener) ExitForInit(ctx *ForInitContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseGScriptListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseGScriptListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseGScriptListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseGScriptListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseGScriptListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseGScriptListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterMultDivExpr is called when production MultDivExpr is entered.
func (s *BaseGScriptListener) EnterMultDivExpr(ctx *MultDivExprContext) {}

// ExitMultDivExpr is called when production MultDivExpr is exited.
func (s *BaseGScriptListener) ExitMultDivExpr(ctx *MultDivExprContext) {}

// EnterPostfixExpr is called when production PostfixExpr is entered.
func (s *BaseGScriptListener) EnterPostfixExpr(ctx *PostfixExprContext) {}

// ExitPostfixExpr is called when production PostfixExpr is exited.
func (s *BaseGScriptListener) ExitPostfixExpr(ctx *PostfixExprContext) {}

// EnterPlusSubExpr is called when production PlusSubExpr is entered.
func (s *BaseGScriptListener) EnterPlusSubExpr(ctx *PlusSubExprContext) {}

// ExitPlusSubExpr is called when production PlusSubExpr is exited.
func (s *BaseGScriptListener) ExitPlusSubExpr(ctx *PlusSubExprContext) {}

// EnterPrimaryExpr is called when production PrimaryExpr is entered.
func (s *BaseGScriptListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production PrimaryExpr is exited.
func (s *BaseGScriptListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseGScriptListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseGScriptListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterModExpr is called when production ModExpr is entered.
func (s *BaseGScriptListener) EnterModExpr(ctx *ModExprContext) {}

// ExitModExpr is called when production ModExpr is exited.
func (s *BaseGScriptListener) ExitModExpr(ctx *ModExprContext) {}

// EnterDotExpr is called when production DotExpr is entered.
func (s *BaseGScriptListener) EnterDotExpr(ctx *DotExprContext) {}

// ExitDotExpr is called when production DotExpr is exited.
func (s *BaseGScriptListener) ExitDotExpr(ctx *DotExprContext) {}

// EnterEqualOrNot is called when production EqualOrNot is entered.
func (s *BaseGScriptListener) EnterEqualOrNot(ctx *EqualOrNotContext) {}

// ExitEqualOrNot is called when production EqualOrNot is exited.
func (s *BaseGScriptListener) ExitEqualOrNot(ctx *EqualOrNotContext) {}

// EnterGLeExpr is called when production GLeExpr is entered.
func (s *BaseGScriptListener) EnterGLeExpr(ctx *GLeExprContext) {}

// ExitGLeExpr is called when production GLeExpr is exited.
func (s *BaseGScriptListener) ExitGLeExpr(ctx *GLeExprContext) {}

// EnterFuncCallExpr is called when production FuncCallExpr is entered.
func (s *BaseGScriptListener) EnterFuncCallExpr(ctx *FuncCallExprContext) {}

// ExitFuncCallExpr is called when production FuncCallExpr is exited.
func (s *BaseGScriptListener) ExitFuncCallExpr(ctx *FuncCallExprContext) {}

// EnterAssignExpr is called when production AssignExpr is entered.
func (s *BaseGScriptListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production AssignExpr is exited.
func (s *BaseGScriptListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterExprPrimary is called when production ExprPrimary is entered.
func (s *BaseGScriptListener) EnterExprPrimary(ctx *ExprPrimaryContext) {}

// ExitExprPrimary is called when production ExprPrimary is exited.
func (s *BaseGScriptListener) ExitExprPrimary(ctx *ExprPrimaryContext) {}

// EnterLiterPrimary is called when production LiterPrimary is entered.
func (s *BaseGScriptListener) EnterLiterPrimary(ctx *LiterPrimaryContext) {}

// ExitLiterPrimary is called when production LiterPrimary is exited.
func (s *BaseGScriptListener) ExitLiterPrimary(ctx *LiterPrimaryContext) {}

// EnterIdentifierPrimary is called when production IdentifierPrimary is entered.
func (s *BaseGScriptListener) EnterIdentifierPrimary(ctx *IdentifierPrimaryContext) {}

// ExitIdentifierPrimary is called when production IdentifierPrimary is exited.
func (s *BaseGScriptListener) ExitIdentifierPrimary(ctx *IdentifierPrimaryContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseGScriptListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseGScriptListener) ExitTypeList(ctx *TypeListContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BaseGScriptListener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BaseGScriptListener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseGScriptListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseGScriptListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseGScriptListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseGScriptListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BaseGScriptListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BaseGScriptListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BaseGScriptListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BaseGScriptListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseGScriptListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseGScriptListener) ExitParse(ctx *ParseContext) {}
