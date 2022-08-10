// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGScriptVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitBlockStms(ctx *BlockStmsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitBlockVarDeclar(ctx *BlockVarDeclarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitBlockStm(ctx *BlockStmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitBlockFunc(ctx *BlockFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitMultDivExpr(ctx *MultDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitPostfixExpr(ctx *PostfixExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitPlusSubExpr(ctx *PlusSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitModExpr(ctx *ModExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitEqualOrNot(ctx *EqualOrNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitGLeExpr(ctx *GLeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitExprPrimary(ctx *ExprPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitLiterPrimary(ctx *LiterPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitIdentifierPrimary(ctx *IdentifierPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitStmBlockLabel(ctx *StmBlockLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitStmIfElse(ctx *StmIfElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitStmFor(ctx *StmForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitStmReturn(ctx *StmReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitStmExpr(ctx *StmExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitForControl(ctx *ForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitVariableModifier(ctx *VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitTypeType(ctx *TypeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGScriptVisitor) VisitParExpression(ctx *ParExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}
