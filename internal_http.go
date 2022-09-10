package gscript

import (
	"fmt"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/gscript/symbol"
	"net/http"
)

type httpH func(http.ResponseWriter, *http.Request)

type httpTool struct {
	w http.ResponseWriter
	r *http.Request
}

var path2HttpTool map[string]*httpTool

func (v *Visitor) httpHandle(ctx *parser.FunctionCallContext) interface{} {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 2 {
		// todo crossoverJie 运行时报错
		panic("")
	}
	p0 := paramValues[0]
	p1 := paramValues[1]
	var path string
	switch p0.(type) {
	case string:
		path = fmt.Sprintf("%s", p0)
	default:
		// todo crossoverJie 运行时报错
		panic("")
	}
	switch p1.(type) {
	case *stack.FuncObject:
		funcObject := p1.(*stack.FuncObject)
		function := funcObject.GetFunction()
		var h = func(w http.ResponseWriter, r *http.Request) {
			// 保存 path 与 HTTPTool 映射关系
			if path2HttpTool == nil {
				path2HttpTool = make(map[string]*httpTool)
			}
			path2HttpTool[path] = &httpTool{w: w, r: r}

			// 注入 classObject 对象:HttpContext
			// todo crossoverJie 基于该能力可以实现反射
			class := symbol.NewClass(ctx, "HttpContext")
			classObject := stack.NewClassObject(class)
			pathVar := v.at.GetHttpPathVariable()
			classObject.SetValue(pathVar, path)
			v.executeFunctionCall(stack.NewFuncObject(function), []interface{}{classObject})
		}
		http.HandleFunc(path, h)
	}

	return nil

}

func (v *Visitor) httpRun(ctx *parser.FunctionCallContext) interface{} {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 1 {
		// todo crossoverJie 运行时报错
		panic("")
	}
	p0 := paramValues[0]
	var addr string
	switch p0.(type) {
	case string:
		addr = fmt.Sprintf("%s", p0)
	}

	if addr != "" {
		fmt.Println(fmt.Sprintf("http host on port%s", addr))
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			// todo crossoverJie 运行时报错
			panic(err)
		}
	} else {
		// todo crossoverJie 运行时报错
	}

	return nil
}

func (v *Visitor) fprintfJSON(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 3 {
		// todo crossoverJie 运行时报错
		panic("")
	}
	p0 := paramValues[0]
	p1 := paramValues[1]
	p2 := paramValues[2]
	var (
		path, json string
		code       int
	)

	switch p0.(type) {
	case int:
		code = p0.(int)
	default:
		// todo crossoverJie 运行时报错
	}
	switch p1.(type) {
	case string:
		path = fmt.Sprintf("%s", p1)
	}
	switch p2.(type) {
	case string:
		json = fmt.Sprintf("%s", p2)
	}

	tool, ok := path2HttpTool[path]
	if !ok {
		// todo crossoverJie 运行时报错
		panic("")
	}
	tool.w.Header().Set("Content-Type", "application/json")
	tool.w.WriteHeader(code)
	fmt.Fprintf(tool.w, json)

}
