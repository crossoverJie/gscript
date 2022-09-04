package gscript

import (
	"fmt"
	"github.com/crossoverJie/gscript/parser"
	"hash/fnv"
)

func (v *Visitor) print(ctx *parser.FunctionCallContext) {
	if ctx.ExpressionList() != nil {
		ret := v.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext))
		switch ret.(type) {
		case *LeftValue:
			ret = ret.(*LeftValue).GetValue()
		case *ArrayObject:
			arrayObject := ret.(*ArrayObject)
			ret = arrayObject.GetIndexValue()
		}
		fmt.Println(ret)
	} else {
		fmt.Println("")
	}
}

func (v *Visitor) assertEqual(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 2 {
		// todo crossoverJie 编译器报错
		panic("")
	}
	if paramValues[0] != paramValues[1] {
		line := ctx.GetStart().GetLine()
		column := ctx.GetStart().GetColumn()
		panic(fmt.Sprintf("assertEqual fail [%v,%v] in line:%d and column:%d", paramValues[0], paramValues[1], line, column))
	}
}

func (v *Visitor) append(ctx *parser.FunctionCallContext) []interface{} {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 2 {
		// todo crossoverJie 运行时报错
		panic("")
	}
	switch paramValues[0].(type) {
	case []interface{}:
		array := paramValues[0].([]interface{})
		array = append(array, paramValues[1])
		return array
	default:
		// todo crossoverJie 运行时报错
	}
	return nil
}

func (v *Visitor) len(ctx *parser.FunctionCallContext) int {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 1 {
		// todo crossoverJie 运行时报错
		panic("")
	}
	switch paramValues[0].(type) {
	case []interface{}:
		return len(paramValues[0].([]interface{}))
	}
	return 0
}
func (v *Visitor) hash(ctx *parser.FunctionCallContext) int {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 1 {
		// todo crossoverJie 运行时报错
		panic("")
	}
	return hash(paramValues[0])
}

func hash(v interface{}) int {
	//hash := sha256.New()
	//hash.Write([]byte(fmt.Sprintf("%v", v)))

	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", v)))
	return int(h.Sum32())

	//return fmt.Sprintf("%x", hash.Sum(nil))
}
