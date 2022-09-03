// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GScriptParser.
type GScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GScriptParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by GScriptParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by GScriptParser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by GScriptParser#memberDeclaration.
	VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{}

	// Visit a parse tree produced by GScriptParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by GScriptParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by GScriptParser#typeTypeOrVoid.
	VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{}

	// Visit a parse tree produced by GScriptParser#qualifiedNameList.
	VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{}

	// Visit a parse tree produced by GScriptParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by GScriptParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by GScriptParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by GScriptParser#lastFormalParameter.
	VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableModifier.
	VisitVariableModifier(ctx *VariableModifierContext) interface{}

	// Visit a parse tree produced by GScriptParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by GScriptParser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by GScriptParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by GScriptParser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by GScriptParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by GScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockStms.
	VisitBlockStms(ctx *BlockStmsContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockVarDeclar.
	VisitBlockVarDeclar(ctx *BlockVarDeclarContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockStm.
	VisitBlockStm(ctx *BlockStmContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockFunc.
	VisitBlockFunc(ctx *BlockFuncContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockClassDeclar.
	VisitBlockClassDeclar(ctx *BlockClassDeclarContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmBlockLabel.
	VisitStmBlockLabel(ctx *StmBlockLabelContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmIfElse.
	VisitStmIfElse(ctx *StmIfElseContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmFor.
	VisitStmFor(ctx *StmForContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmWhile.
	VisitStmWhile(ctx *StmWhileContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmReturn.
	VisitStmReturn(ctx *StmReturnContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmBreak.
	VisitStmBreak(ctx *StmBreakContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmContinue.
	VisitStmContinue(ctx *StmContinueContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmExpr.
	VisitStmExpr(ctx *StmExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by GScriptParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by GScriptParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}

	// Visit a parse tree produced by GScriptParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by GScriptParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by GScriptParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by GScriptParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by GScriptParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by GScriptParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by GScriptParser#parse.
	VisitParse(ctx *ParseContext) interface{}
}
