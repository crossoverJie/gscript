// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GScriptListener is a complete listener for a parse tree produced by GScriptParser.
type GScriptListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStms is called when entering the BlockStms production.
	EnterBlockStms(c *BlockStmsContext)

	// EnterBlockVarDeclar is called when entering the BlockVarDeclar production.
	EnterBlockVarDeclar(c *BlockVarDeclarContext)

	// EnterBlockStm is called when entering the BlockStm production.
	EnterBlockStm(c *BlockStmContext)

	// EnterBlockFunc is called when entering the BlockFunc production.
	EnterBlockFunc(c *BlockFuncContext)

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterMultDivExpr is called when entering the MultDivExpr production.
	EnterMultDivExpr(c *MultDivExprContext)

	// EnterPostfixExpr is called when entering the PostfixExpr production.
	EnterPostfixExpr(c *PostfixExprContext)

	// EnterGLe is called when entering the GLe production.
	EnterGLe(c *GLeContext)

	// EnterPlusSubExpr is called when entering the PlusSubExpr production.
	EnterPlusSubExpr(c *PlusSubExprContext)

	// EnterPrimaryExpr is called when entering the PrimaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterModExpr is called when entering the ModExpr production.
	EnterModExpr(c *ModExprContext)

	// EnterEqualOrNot is called when entering the EqualOrNot production.
	EnterEqualOrNot(c *EqualOrNotContext)

	// EnterExprPrimary is called when entering the ExprPrimary production.
	EnterExprPrimary(c *ExprPrimaryContext)

	// EnterLiterPrimary is called when entering the LiterPrimary production.
	EnterLiterPrimary(c *LiterPrimaryContext)

	// EnterIdentifierPrimay is called when entering the IdentifierPrimay production.
	EnterIdentifierPrimay(c *IdentifierPrimayContext)

	// EnterBlockLabel is called when entering the BlockLabel production.
	EnterBlockLabel(c *BlockLabelContext)

	// EnterIfElse is called when entering the IfElse production.
	EnterIfElse(c *IfElseContext)

	// EnterFor is called when entering the For production.
	EnterFor(c *ForContext)

	// EnterReturn is called when entering the Return production.
	EnterReturn(c *ReturnContext)

	// EnterStmExpr is called when entering the StmExpr production.
	EnterStmExpr(c *StmExprContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterQualifiedNameList is called when entering the qualifiedNameList production.
	EnterQualifiedNameList(c *QualifiedNameListContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterLastFormalParameter is called when entering the lastFormalParameter production.
	EnterLastFormalParameter(c *LastFormalParameterContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterVariableModifier is called when entering the variableModifier production.
	EnterVariableModifier(c *VariableModifierContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterBool is called when entering the Bool production.
	EnterBool(c *BoolContext)

	// EnterNull is called when entering the Null production.
	EnterNull(c *NullContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterTypeType is called when entering the typeType production.
	EnterTypeType(c *TypeTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeTypeOrVoid is called when entering the typeTypeOrVoid production.
	EnterTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStms is called when exiting the BlockStms production.
	ExitBlockStms(c *BlockStmsContext)

	// ExitBlockVarDeclar is called when exiting the BlockVarDeclar production.
	ExitBlockVarDeclar(c *BlockVarDeclarContext)

	// ExitBlockStm is called when exiting the BlockStm production.
	ExitBlockStm(c *BlockStmContext)

	// ExitBlockFunc is called when exiting the BlockFunc production.
	ExitBlockFunc(c *BlockFuncContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitMultDivExpr is called when exiting the MultDivExpr production.
	ExitMultDivExpr(c *MultDivExprContext)

	// ExitPostfixExpr is called when exiting the PostfixExpr production.
	ExitPostfixExpr(c *PostfixExprContext)

	// ExitGLe is called when exiting the GLe production.
	ExitGLe(c *GLeContext)

	// ExitPlusSubExpr is called when exiting the PlusSubExpr production.
	ExitPlusSubExpr(c *PlusSubExprContext)

	// ExitPrimaryExpr is called when exiting the PrimaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitModExpr is called when exiting the ModExpr production.
	ExitModExpr(c *ModExprContext)

	// ExitEqualOrNot is called when exiting the EqualOrNot production.
	ExitEqualOrNot(c *EqualOrNotContext)

	// ExitExprPrimary is called when exiting the ExprPrimary production.
	ExitExprPrimary(c *ExprPrimaryContext)

	// ExitLiterPrimary is called when exiting the LiterPrimary production.
	ExitLiterPrimary(c *LiterPrimaryContext)

	// ExitIdentifierPrimay is called when exiting the IdentifierPrimay production.
	ExitIdentifierPrimay(c *IdentifierPrimayContext)

	// ExitBlockLabel is called when exiting the BlockLabel production.
	ExitBlockLabel(c *BlockLabelContext)

	// ExitIfElse is called when exiting the IfElse production.
	ExitIfElse(c *IfElseContext)

	// ExitFor is called when exiting the For production.
	ExitFor(c *ForContext)

	// ExitReturn is called when exiting the Return production.
	ExitReturn(c *ReturnContext)

	// ExitStmExpr is called when exiting the StmExpr production.
	ExitStmExpr(c *StmExprContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitQualifiedNameList is called when exiting the qualifiedNameList production.
	ExitQualifiedNameList(c *QualifiedNameListContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitLastFormalParameter is called when exiting the lastFormalParameter production.
	ExitLastFormalParameter(c *LastFormalParameterContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitVariableModifier is called when exiting the variableModifier production.
	ExitVariableModifier(c *VariableModifierContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitBool is called when exiting the Bool production.
	ExitBool(c *BoolContext)

	// ExitNull is called when exiting the Null production.
	ExitNull(c *NullContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitTypeType is called when exiting the typeType production.
	ExitTypeType(c *TypeTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeTypeOrVoid is called when exiting the typeTypeOrVoid production.
	ExitTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)
}
