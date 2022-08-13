// Code generated from GScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GScriptListener is a complete listener for a parse tree produced by GScriptParser.
type GScriptListener interface {
	antlr.ParseTreeListener

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterTypeTypeOrVoid is called when entering the typeTypeOrVoid production.
	EnterTypeTypeOrVoid(c *TypeTypeOrVoidContext)

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

	// EnterVariableModifier is called when entering the variableModifier production.
	EnterVariableModifier(c *VariableModifierContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

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

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

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

	// EnterStmBlockLabel is called when entering the StmBlockLabel production.
	EnterStmBlockLabel(c *StmBlockLabelContext)

	// EnterStmIfElse is called when entering the StmIfElse production.
	EnterStmIfElse(c *StmIfElseContext)

	// EnterStmFor is called when entering the StmFor production.
	EnterStmFor(c *StmForContext)

	// EnterStmReturn is called when entering the StmReturn production.
	EnterStmReturn(c *StmReturnContext)

	// EnterStmExpr is called when entering the StmExpr production.
	EnterStmExpr(c *StmExprContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterMultDivExpr is called when entering the MultDivExpr production.
	EnterMultDivExpr(c *MultDivExprContext)

	// EnterPostfixExpr is called when entering the PostfixExpr production.
	EnterPostfixExpr(c *PostfixExprContext)

	// EnterPlusSubExpr is called when entering the PlusSubExpr production.
	EnterPlusSubExpr(c *PlusSubExprContext)

	// EnterPrimaryExpr is called when entering the PrimaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterModExpr is called when entering the ModExpr production.
	EnterModExpr(c *ModExprContext)

	// EnterDotExpr is called when entering the DotExpr production.
	EnterDotExpr(c *DotExprContext)

	// EnterEqualOrNot is called when entering the EqualOrNot production.
	EnterEqualOrNot(c *EqualOrNotContext)

	// EnterGLeExpr is called when entering the GLeExpr production.
	EnterGLeExpr(c *GLeExprContext)

	// EnterFuncCallExpr is called when entering the FuncCallExpr production.
	EnterFuncCallExpr(c *FuncCallExprContext)

	// EnterAssignExpr is called when entering the AssignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterExprPrimary is called when entering the ExprPrimary production.
	EnterExprPrimary(c *ExprPrimaryContext)

	// EnterLiterPrimary is called when entering the LiterPrimary production.
	EnterLiterPrimary(c *LiterPrimaryContext)

	// EnterIdentifierPrimary is called when entering the IdentifierPrimary production.
	EnterIdentifierPrimary(c *IdentifierPrimaryContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeType is called when entering the typeType production.
	EnterTypeType(c *TypeTypeContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitTypeTypeOrVoid is called when exiting the typeTypeOrVoid production.
	ExitTypeTypeOrVoid(c *TypeTypeOrVoidContext)

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

	// ExitVariableModifier is called when exiting the variableModifier production.
	ExitVariableModifier(c *VariableModifierContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

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

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

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

	// ExitStmBlockLabel is called when exiting the StmBlockLabel production.
	ExitStmBlockLabel(c *StmBlockLabelContext)

	// ExitStmIfElse is called when exiting the StmIfElse production.
	ExitStmIfElse(c *StmIfElseContext)

	// ExitStmFor is called when exiting the StmFor production.
	ExitStmFor(c *StmForContext)

	// ExitStmReturn is called when exiting the StmReturn production.
	ExitStmReturn(c *StmReturnContext)

	// ExitStmExpr is called when exiting the StmExpr production.
	ExitStmExpr(c *StmExprContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitMultDivExpr is called when exiting the MultDivExpr production.
	ExitMultDivExpr(c *MultDivExprContext)

	// ExitPostfixExpr is called when exiting the PostfixExpr production.
	ExitPostfixExpr(c *PostfixExprContext)

	// ExitPlusSubExpr is called when exiting the PlusSubExpr production.
	ExitPlusSubExpr(c *PlusSubExprContext)

	// ExitPrimaryExpr is called when exiting the PrimaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitModExpr is called when exiting the ModExpr production.
	ExitModExpr(c *ModExprContext)

	// ExitDotExpr is called when exiting the DotExpr production.
	ExitDotExpr(c *DotExprContext)

	// ExitEqualOrNot is called when exiting the EqualOrNot production.
	ExitEqualOrNot(c *EqualOrNotContext)

	// ExitGLeExpr is called when exiting the GLeExpr production.
	ExitGLeExpr(c *GLeExprContext)

	// ExitFuncCallExpr is called when exiting the FuncCallExpr production.
	ExitFuncCallExpr(c *FuncCallExprContext)

	// ExitAssignExpr is called when exiting the AssignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitExprPrimary is called when exiting the ExprPrimary production.
	ExitExprPrimary(c *ExprPrimaryContext)

	// ExitLiterPrimary is called when exiting the LiterPrimary production.
	ExitLiterPrimary(c *LiterPrimaryContext)

	// ExitIdentifierPrimary is called when exiting the IdentifierPrimary production.
	ExitIdentifierPrimary(c *IdentifierPrimaryContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeType is called when exiting the typeType production.
	ExitTypeType(c *TypeTypeContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)
}
