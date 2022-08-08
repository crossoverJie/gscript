// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GScriptParser.
type GScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GScriptParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockStms.
	VisitBlockStms(ctx *BlockStmsContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockVarDeclar.
	VisitBlockVarDeclar(ctx *BlockVarDeclarContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockStm.
	VisitBlockStm(ctx *BlockStmContext) interface{}

	// Visit a parse tree produced by GScriptParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by GScriptParser#MultDivExpr.
	VisitMultDivExpr(ctx *MultDivExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#PostfixExpr.
	VisitPostfixExpr(ctx *PostfixExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#GLe.
	VisitGLe(ctx *GLeContext) interface{}

	// Visit a parse tree produced by GScriptParser#Liter.
	VisitLiter(ctx *LiterContext) interface{}

	// Visit a parse tree produced by GScriptParser#PlusSubExpr.
	VisitPlusSubExpr(ctx *PlusSubExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#NestedExpr.
	VisitNestedExpr(ctx *NestedExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#ModExpr.
	VisitModExpr(ctx *ModExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#EqualOrNot.
	VisitEqualOrNot(ctx *EqualOrNotContext) interface{}

	// Visit a parse tree produced by GScriptParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#PrefixExpr.
	VisitPrefixExpr(ctx *PrefixExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GScriptParser#BlockLabel.
	VisitBlockLabel(ctx *BlockLabelContext) interface{}

	// Visit a parse tree produced by GScriptParser#IfElse.
	VisitIfElse(ctx *IfElseContext) interface{}

	// Visit a parse tree produced by GScriptParser#For.
	VisitFor(ctx *ForContext) interface{}

	// Visit a parse tree produced by GScriptParser#Return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by GScriptParser#StmExpr.
	VisitStmExpr(ctx *StmExprContext) interface{}

	// Visit a parse tree produced by GScriptParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by GScriptParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by GScriptParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by GScriptParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by GScriptParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by GScriptParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by GScriptParser#Bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by GScriptParser#Null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by GScriptParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by GScriptParser#floatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by GScriptParser#typeType.
	VisitTypeType(ctx *TypeTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#functionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by GScriptParser#typeTypeOrVoid.
	VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{}

	// Visit a parse tree produced by GScriptParser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by GScriptParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by GScriptParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}
}
