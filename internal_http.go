package gscript

import (
	"fmt"
	"github.com/crossoverJie/gscript/log"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/gscript/symbol"
	"io"
	"net"
	"net/http"
	"strings"
)

type httpTool struct {
	w http.ResponseWriter
	r *http.Request
}

// path 与 http 工具的映射关系
var path2HttpTool map[string]*httpTool

func (v *Visitor) httpHandle(ctx *parser.FunctionCallContext) interface{} {
	paramValues := v.buildParamValues(ctx)
	p := paramValues[0]
	p0 := paramValues[1]
	p1 := paramValues[2]
	var (
		method, path string
	)
	switch p.(type) {
	case string:
		method = fmt.Sprintf("%s", p)
	}

	switch p0.(type) {
	case string:
		path = fmt.Sprintf("%s", p0)
	}
	switch p1.(type) {
	case *stack.FuncObject:
		funcObject := p1.(*stack.FuncObject)
		function := funcObject.GetFunction()
		var h = func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					switch x := r.(type) {
					case *log.Log:
						http.Error(w, x.String(), http.StatusInternalServerError)
					default:
						panic(x)
					}
				}
			}()
			if r.Method != strings.ToUpper(method) {
				// todo crossoverJie http panic 单独 recovery
				log.RuntimePanic(ctx, fmt.Sprintf("http method %s not correct", method))
			}
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
	p0 := paramValues[0]
	var addr string
	switch p0.(type) {
	case string:
		addr = fmt.Sprintf("%s", p0)
	}

	if addr != "" {
		ip := externalIP()
		fmt.Println(fmt.Sprintf("http host %s on port%s", ip, addr))
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.RuntimePanic(ctx, fmt.Sprintf("httpRun function error occurred,error:%s", err))
		}
	} else {
		log.RuntimePanic(ctx, fmt.Sprintf("addr cannot be empty"))
	}

	return nil
}

func (v *Visitor) fprintfJSON(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	code, path, json := v.getCodePathValue(paramValues)
	tool, ok := path2HttpTool[path]
	if !ok {
		log.RuntimePanic(ctx, fmt.Sprintf("http handle path not fount"))
	}
	tool.w.Header().Set("Content-Type", "application/json")
	if code != http.StatusOK {
		tool.w.WriteHeader(code)
	}
	fmt.Fprintf(tool.w, json)

}

func (v *Visitor) fprintfHTML(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	code, path, html := v.getCodePathValue(paramValues)
	tool, ok := path2HttpTool[path]
	if !ok {
		log.RuntimePanic(ctx, fmt.Sprintf("http handle path not fount"))
	}
	tool.w.Header().Set("Content-Type", "text/html")
	if code != http.StatusOK {
		tool.w.WriteHeader(code)
	}
	fmt.Fprintf(tool.w, html)
}

func (v *Visitor) requestBody(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	var path string
	switch p0.(type) {
	case string:
		path = fmt.Sprintf("%s", p0)
	}

	tool, ok := path2HttpTool[path]
	if !ok {
		log.RuntimePanic(ctx, fmt.Sprintf("http handle path not fount"))
	}
	data, err := io.ReadAll(tool.r.Body)
	if err != nil {
		log.RuntimePanic(ctx, fmt.Sprintf("read body error:%s", err))
	}

	return string(data)

}

func (v *Visitor) getCodePathValue(paramValues []interface{}) (int, string, string) {
	p0 := paramValues[0]
	p1 := paramValues[1]
	p2 := paramValues[2]
	var (
		path, value string
		code        int
	)

	switch p0.(type) {
	case int:
		code = p0.(int)
	}
	switch p1.(type) {
	case string:
		path = fmt.Sprintf("%s", p1)
	}
	switch p2.(type) {
	case string:
		value = fmt.Sprintf("%s", p2)
	}

	return code, path, value
}

func (v *Visitor) queryPath(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	var path string
	switch p0.(type) {
	case string:
		path = fmt.Sprintf("%s", p0)
	}

	tool, ok := path2HttpTool[path]
	if !ok {
		log.RuntimePanic(ctx, fmt.Sprintf("http handle path not fount"))
	}
	return tool.r.URL.Path

}

func (v *Visitor) formValue(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	p1 := paramValues[1]
	var (
		path, key string
	)

	switch p0.(type) {
	case string:
		path = fmt.Sprintf("%s", p0)
	}

	switch p1.(type) {
	case string:
		key = fmt.Sprintf("%s", p1)
	}

	tool, ok := path2HttpTool[path]
	if !ok {
		log.RuntimePanic(ctx, fmt.Sprintf("http handle path not fount"))
	}

	return tool.r.FormValue(key)

}
func (v *Visitor) postFormValue(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	p1 := paramValues[1]
	var (
		path, key string
	)

	switch p0.(type) {
	case string:
		path = fmt.Sprintf("%s", p0)
	}

	switch p1.(type) {
	case string:
		key = fmt.Sprintf("%s", p1)
	}

	tool, ok := path2HttpTool[path]
	if !ok {
		log.RuntimePanic(ctx, fmt.Sprintf("http handle path not fount"))
	}

	return tool.r.PostFormValue(key)

}

func externalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String()
		}
	}
	return ""
}
