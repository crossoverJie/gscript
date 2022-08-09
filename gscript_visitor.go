package gscript

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"strconv"
)

type GScriptVisitor struct {
	parser.BaseGScriptVisitor
}

func ArithmeticOperators(script string) interface{} {
	input := antlr.NewInputStream(script)
	lexer := parser.NewGScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	tree := parser.NewGScriptParser(stream).Prog()
	visitor := GScriptVisitor{}
	return visitor.Visit(tree)
}

func (v *GScriptVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(ctx)
	case *parser.BlockStmContext:
		return v.VisitBlockStm(ctx)
	case *parser.BlockLabelContext:
		return v.VisitBlockLabel(ctx)
	case *parser.BlockVarDeclarContext:
		return v.VisitBlockVarDeclar(ctx)
	case *parser.VariableDeclaratorsContext:
		return v.VisitVariableDeclarators(ctx)
	case *parser.ParseContext:
		return v.VisitParse(ctx)
	case *parser.PostfixExprContext:
		return v.VisitPostfixExpr(ctx)
	case *parser.NotExprContext:
		return v.VisitNotExpr(ctx)
	case *parser.MultDivExprContext:
		return v.VisitMultDivExpr(ctx)
	case *parser.PrimaryExprContext:
		return v.VisitPrimaryExpr(ctx)
	case *parser.LiterPrimaryContext:
		return v.VisitLiterPrimary(ctx)
	case *parser.IntContext:
		return v.VisitInt(ctx)
	case *parser.FloatContext:
		return v.VisitFloat(ctx)
	case *parser.StringContext:
		return v.VisitString(ctx)
	case *parser.BoolContext:
		return v.VisitBool(ctx)
	case *parser.NullContext:
		return v.VisitNull(ctx)
	case *parser.PlusSubExprContext:
		return v.VisitPlusSubExpr(ctx)
	case *parser.ExprPrimaryContext:
		return v.VisitExprPrimary(ctx)
	//case *parser.UnaryExprContext:
	//	return v.VisitUnaryExpr(ctx)
	case *parser.ModExprContext:
		return v.VisitModExpr(ctx)
	case *parser.GLeContext:
		return v.VisitGLe(ctx)
	case *parser.EqualOrNotContext:
		return v.VisitEqualOrNot(ctx)
	case *parser.IfElseContext:
		return v.VisitIfElse(ctx)
	case *parser.ReturnContext:
		return v.VisitReturn(ctx)
	case *parser.StmExprContext:
		return v.VisitStmExpr(ctx)
	default:
		panic("Unknown context")
	}
}

func (v *GScriptVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	return v.VisitBlockStms(ctx.BlockStatements().(*parser.BlockStmsContext))
}

func (v *GScriptVisitor) VisitBlockStms(ctx *parser.BlockStmsContext) interface{} {
	var ret interface{}
	for _, context := range ctx.AllBlockStatement() {
		ret = v.Visit(context)
	}
	return ret
}

func (v *GScriptVisitor) VisitBlockVarDeclar(ctx *parser.BlockVarDeclarContext) interface{} {
	return v.Visit(ctx.VariableDeclarators())
}

func (v *GScriptVisitor) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	var ret interface{}
	for _, context := range ctx.AllVariableDeclarator() {
		ret = v.VisitVariableDeclarator(context.(*parser.VariableDeclaratorContext))
	}
	return ret
}

func (v *GScriptVisitor) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	v.VisitVariableDeclaratorId(ctx.VariableDeclaratorId().(*parser.VariableDeclaratorIdContext))
	if ctx.VariableInitializer() != nil {
		v.VisitVariableInitializer(ctx.VariableInitializer().(*parser.VariableInitializerContext))
	}
	return v.VisitChildren(ctx)
}

func (v *GScriptVisitor) VisitVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) interface{} {
	return ""
}
func (v *GScriptVisitor) VisitVariableInitializer(ctx *parser.VariableInitializerContext) interface{} {
	// todo array init
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	return nil
}

func (v *GScriptVisitor) VisitBlockStm(ctx *parser.BlockStmContext) interface{} {
	return v.Visit(ctx.Statement())
}

func (v *GScriptVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitBlockStms(ctx.BlockStatements().(*parser.BlockStmsContext))
}

func (v *GScriptVisitor) VisitBlockLabel(ctx *parser.BlockLabelContext) interface{} {
	return v.VisitBlock(ctx.GetBlockLabel().(*parser.BlockContext))
}

func (v *GScriptVisitor) VisitParse(ctx *parser.ParseContext) interface{} {
	for _, expr := range ctx.GetExpr_list() {
		return v.Visit(expr)
	}
	return nil
}

func (v *GScriptVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	return v.Visit(ctx.Primary())
}

func (v *GScriptVisitor) VisitPostfixExpr(ctx *parser.PostfixExprContext) interface{} {
	lhs := ctx.GetLhs()
	value := v.Visit(lhs)
	switch ctx.GetPostfix().GetTokenType() {
	case parser.GScriptParserINC:
		switch value.(type) {
		case int:
			value = value.(int) + 1
			return value
		}
	case parser.GScriptParserDEC:
		switch value.(type) {
		case int:
			value = value.(int) - 1
			return value
		}
	}
	panic("invalid postfix")
}

func (v *GScriptVisitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	rhs := ctx.GetRhs()
	value := v.Visit(rhs)
	switch value.(type) {
	case bool:
		return !value.(bool)
	}
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	panic(fmt.Sprintf("invalid ! symbol in line:%d and column:%d", line, column))
}

func (v *GScriptVisitor) VisitMultDivExpr(ctx *parser.MultDivExprContext) interface{} {
	val1 := v.Visit(ctx.GetLhs())
	val2 := v.Visit(ctx.GetRhs())

	var (
		v1i int
		v1f float64
		v2i int
		v2f float64
	)
	switch val1.(type) {
	case int:
		v1i = val1.(int)
	case float64:
		v1f = val1.(float64)
	}

	switch val2.(type) {
	case int:
		v2i = val2.(int)
	case float64:
		v2f = val2.(float64)
	}

	if ctx.GetBop().GetTokenType() == parser.GScriptParserMULT {
		if v1i != 0 && v2i != 0 {
			return v1i * v2i
		}
		if v1i != 0 && v2f != 0 {
			return float64(v1i) * v2f
		}
		if v1f != 0 && v2i != 0 {
			return v1f * float64(v2i)
		}
		if v1f != 0 && v2f != 0 {
			return v1f * v2f
		}
	} else {
		if v1i != 0 && v2i != 0 {
			return v1i / v2i
		}
		if v1i != 0 && v2f != 0 {
			return float64(v1i) / v2f
		}
		if v1f != 0 && v2i != 0 {
			return v1f / float64(v2i)
		}
		if v1f != 0 && v2f != 0 {
			return v1f / v2f
		}
	}
	return 0
}

func (v *GScriptVisitor) VisitPlusSubExpr(ctx *parser.PlusSubExprContext) interface{} {
	val1 := v.Visit(ctx.GetLhs())
	val2 := v.Visit(ctx.GetRhs())

	var (
		v1i int
		v1f float64
		v2i int
		v2f float64
	)
	switch val1.(type) {
	case int:
		v1i = val1.(int)
	case float64:
		v1f = val1.(float64)
	}

	switch val2.(type) {
	case int:
		v2i = val2.(int)
	case float64:
		v2f = val2.(float64)
	}

	if ctx.GetBop().GetTokenType() == parser.GScriptLexerPLUS {
		if v1i != 0 && v2i != 0 {
			return v1i + v2i
		}
		if v1i != 0 && v2f != 0 {
			return float64(v1i) + v2f
		}
		if v1f != 0 && v2i != 0 {
			return v1f + float64(v2i)
		}
		if v1f != 0 && v2f != 0 {
			return v1f + v2f
		}
	} else {
		if v1i != 0 && v2i != 0 {
			return v1i - v2i
		}
		if v1i != 0 && v2f != 0 {
			return float64(v1i) - v2f
		}
		if v1f != 0 && v2i != 0 {
			return v1f - float64(v2i)
		}
		if v1f != 0 && v2f != 0 {
			return v1f - v2f
		}
	}
	return 0
}

func (v *GScriptVisitor) VisitLiterPrimary(ctx *parser.LiterPrimaryContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *GScriptVisitor) VisitIdentifierPrimary(ctx *parser.IdentifierPrimayContext) interface{} {
	return ctx.IDENTIFIER().GetText()
}

func (v *GScriptVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	val, _ := strconv.Atoi(ctx.GetText())
	return val
}

func (v *GScriptVisitor) VisitFloat(ctx *parser.FloatContext) interface{} {
	val, _ := strconv.ParseFloat(ctx.GetText(), 0)
	return val
}

func (v *GScriptVisitor) VisitString(ctx *parser.StringContext) interface{} {
	return ctx.GetText()[1 : len(ctx.GetText())-1]
}

func (v *GScriptVisitor) VisitBool(ctx *parser.BoolContext) interface{} {
	val, _ := strconv.ParseBool(ctx.GetText())
	return val
}

func (v *GScriptVisitor) VisitNull(ctx *parser.NullContext) interface{} {
	return nil
}

func (v *GScriptVisitor) VisitExprPrimary(ctx *parser.ExprPrimaryContext) interface{} {
	return v.Visit(ctx.Expr())
}

//func (v *GScriptVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
//	return -v.Visit(ctx.Expr()).(int)
//}

func (v *GScriptVisitor) VisitModExpr(ctx *parser.ModExprContext) interface{} {
	lhs := v.Visit(ctx.GetLhs())
	rhs := v.Visit(ctx.GetRhs())
	return lhs.(int) % rhs.(int)
}

func (v *GScriptVisitor) VisitGLe(ctx *parser.GLeContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)
	switch ctx.GetBop().GetTokenType() {
	case parser.GScriptParserGT:
		return left > right
	case parser.GScriptParserLT:
		return left < right
	case parser.GScriptParserGE:
		return left >= right
	case parser.GScriptParserLE:
		return left <= right
	default:
		panic("invalid symbol")
	}
}

func (v *GScriptVisitor) VisitEqualOrNot(ctx *parser.EqualOrNotContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)
	switch ctx.GetBop().GetTokenType() {
	case parser.GScriptParserEQUAL:
		return left == right
	case parser.GScriptParserNOTEQUAL:
		return left != right
	default:
		panic("invalid symbol")
	}
}

func (v *GScriptVisitor) VisitIfElse(ctx *parser.IfElseContext) interface{} {
	condition := v.VisitParExpression(ctx.ParExpression().(*parser.ParExpressionContext)).(bool)
	if condition {
		return v.Visit(ctx.Statement(0))
	} else if ctx.ELSE() != nil {
		return v.Visit(ctx.Statement(1))
	}
	return nil
}

func (v *GScriptVisitor) VisitReturn(ctx *parser.ReturnContext) interface{} {
	if ctx.Expr() != nil {
		return v.Visit(ctx.Expr())
	}
	return nil
}

func (v *GScriptVisitor) VisitStmExpr(ctx *parser.StmExprContext) interface{} {
	return v.Visit(ctx.GetStatementExpression())
}

func (v *GScriptVisitor) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return v.Visit(ctx.Expr())
}
