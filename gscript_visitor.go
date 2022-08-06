package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/crossoverJie/gscript/parser"
	"strconv"
)

type GScriptVisitor struct {
	parser.BaseGScriptVisitor
}

func (v *GScriptVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(ctx)
	case *parser.BlockStmContext:
		return v.VisitBlockStm(ctx)
	case *parser.BlockLabelContext:
		return v.VisitBlockLabel(ctx)
	case *parser.ParseContext:
		return v.VisitParse(ctx)
	case *parser.MultDivExprContext:
		return v.VisitMultDivExpr(ctx)
	case *parser.LiterContext:
		return v.VisitLiter(ctx)
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
	case *parser.NestedExprContext:
		return v.VisitNestedExpr(ctx)
	case *parser.UnaryExprContext:
		return v.VisitUnaryExpr(ctx)
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

func (v *GScriptVisitor) VisitMultDivExpr(ctx *parser.MultDivExprContext) interface{} {
	lhs := v.Visit(ctx.GetLhs())
	rhs := v.Visit(ctx.GetRhs())

	if ctx.GetBop().GetTokenType() == parser.GScriptLexerMULT {
		return lhs.(int64) * rhs.(int64)
	} else {
		return lhs.(int64) / rhs.(int64)
	}
}

func (v *GScriptVisitor) VisitPlusSubExpr(ctx *parser.PlusSubExprContext) interface{} {
	lhs := v.Visit(ctx.GetLhs())
	rhs := v.Visit(ctx.GetRhs())

	if ctx.GetBop().GetTokenType() == parser.GScriptLexerPLUS {
		return lhs.(int64) + rhs.(int64)
	} else {
		return lhs.(int64) - rhs.(int64)
	}
}

func (v *GScriptVisitor) VisitLiter(ctx *parser.LiterContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *GScriptVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	val, _ := strconv.ParseInt(ctx.GetText(), 32, 10)
	return val
}

func (v *GScriptVisitor) VisitFloat(ctx *parser.FloatContext) interface{} {
	val, _ := strconv.ParseFloat(ctx.GetText(), 10)
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

func (v *GScriptVisitor) VisitNestedExpr(ctx *parser.NestedExprContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *GScriptVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	return -v.Visit(ctx.Expr()).(int64)
}

func (v *GScriptVisitor) VisitModExpr(ctx *parser.ModExprContext) interface{} {
	lhs := v.Visit(ctx.GetLhs())
	rhs := v.Visit(ctx.GetRhs())
	return lhs.(int64) % rhs.(int64)
}

func (v *GScriptVisitor) VisitGLe(ctx *parser.GLeContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int64)
	right := v.Visit(ctx.Expr(1)).(int64)
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
	left := v.Visit(ctx.Expr(0)).(int64)
	right := v.Visit(ctx.Expr(1)).(int64)
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
