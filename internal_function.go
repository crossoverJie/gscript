package gscript

import (
	"encoding/json"
	"fmt"
	"github.com/crossoverJie/gscript/log"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/xjson"
	"hash/fnv"
	"io/fs"
	"os"
	"os/exec"
	"reflect"
	"time"
)

func (v *Visitor) println(ctx *parser.FunctionCallContext) {
	if ctx.ExpressionList() != nil {
		ret := v.VisitExpressionList(ctx.ExpressionList().(*parser.ExpressionListContext))
		switch ret.(type) {
		case *LeftValue:
			ret = ret.(*LeftValue).GetValue()
			switch ret.(type) {
			case []interface{}:
				// 对象数组 Person[] list = {p1,p2}; println(list);
				i := ret.([]interface{})
				ret = v.printArray(i)
			}
		case *ArrayObject:
			arrayObject := ret.(*ArrayObject)
			ret = arrayObject.GetIndexValue()
		}
		fmt.Println(ret)
		fmt.Print()
	} else {
		fmt.Println("")
	}
}

// 打印数组对象 Person[] p
func (v *Visitor) printArray(value []interface{}) []interface{} {
	if len(value) == 0 {
		return nil
	}
	var retList []interface{}
	for _, val := range value {
		switch val.(type) {
		case *LeftValue:
			leftValue := val.(*LeftValue).GetValue()
			switch leftValue.(type) {
			case *stack.ClassObject:
				classObject := leftValue.(*stack.ClassObject)
				var data []interface{}
				for _, val := range classObject.AllField() {
					data = append(data, val)
				}
				// [[1 a] [1 b]]  Person[] list = {p1,p2};
				retList = append(retList, data)
			}
		}
	}
	if len(retList) > 0 {
		return retList
	}
	return value
}

func (v *Visitor) assertEqual(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	// todo crossoverJie 参数是个变量，需要取左值，也可以是个数组取值 a[0]
	if paramValues[0] != paramValues[1] {
		line := ctx.GetStart().GetLine()
		column := ctx.GetStart().GetColumn()
		panic(fmt.Sprintf("assertEqual fail [%v,%v] in line:%d and column:%d", paramValues[0], paramValues[1], line, column))
	}
}

func (v *Visitor) append(ctx *parser.FunctionCallContext) []interface{} {
	paramValues, left := v.buildAppendParamValuesReturnLeft(ctx)
	switch paramValues[0].(type) {
	case []interface{}:
		array := paramValues[0].([]interface{})
		array = append(array, paramValues[1])
		left.SetValue(array)
		return array
	case []byte:
		array := paramValues[0].([]byte)
		value := paramValues[1].([]byte)
		array = append(array, value...)
		left.SetValue(array)
	default:
		log.RuntimePanic(ctx, fmt.Sprintf("first argument to append must be array"))

	}
	return nil
}

func (v *Visitor) len(ctx *parser.FunctionCallContext) int {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	switch p0.(type) {
	case []interface{}:
		return len(p0.([]interface{}))
	case []string:
		return len(p0.([]string))
	case []float64:
		return len(p0.([]float64))
	case []bool:
		return len(p0.([]bool))
	case []int:
		return len(p0.([]int))
	case []byte:
		return len(p0.([]byte))
	}
	return 0
}
func (v *Visitor) cap(ctx *parser.FunctionCallContext) int {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	switch p0.(type) {
	case []interface{}:
		return cap(p0.([]interface{}))
	case []string:
		return cap(p0.([]string))
	case []float64:
		return cap(p0.([]float64))
	case []bool:
		return cap(p0.([]bool))
	case []int:
		return cap(p0.([]int))
	case []byte:
		return cap(p0.([]byte))
	}
	return 0
}

func (v *Visitor) copy(ctx *parser.FunctionCallContext) int {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	p1 := paramValues[1]
	return copy(p0.([]interface{}), p1.([]interface{}))
}
func (v *Visitor) hash(ctx *parser.FunctionCallContext) int {
	paramValues := v.buildParamValues(ctx)
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

func (v *Visitor) JSON(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	value := paramValues[0]
	switch value.(type) {
	case *stack.ClassObject:
		classObject := value.(*stack.ClassObject)
		data := v.classObject2Map(classObject)
		marshal, err := json.Marshal(data)
		if err != nil {
			log.RuntimePanic(ctx, fmt.Sprintf("JSON function error occurred,error:%s", err))
		}
		return string(marshal)
	case []interface{}:
		dataList := value.([]interface{})
		var dataClass []interface{}
		for _, data := range dataList {
			switch data.(type) {
			case *LeftValue:
				// Person[] list = {p1,p2}; string j = JSON(list);
				leftValue := data.(*LeftValue)
				getValue := leftValue.GetValue()
				classObject, ok := getValue.(*stack.ClassObject)
				if ok {
					object2Map := v.classObject2Map(classObject)
					dataClass = append(dataClass, object2Map)
				}

			}
		}
		if dataClass != nil {
			marshal, err := json.Marshal(dataClass)
			if err != nil {
				log.RuntimePanic(ctx, fmt.Sprintf("JSON function error occurred,error:%s", err))

			}
			return string(marshal)
		}
		// int[] a = {1,2,3}; json = JSON(a)
		marshal, err := json.Marshal(dataList)
		if err != nil {
			log.RuntimePanic(ctx, fmt.Sprintf("JSON function error occurred,error:%s", err))

		}
		// {1,2,3}
		return string(marshal)
	default:
		marshal, err := json.Marshal(value)
		if err != nil {
			log.RuntimePanic(ctx, fmt.Sprintf("JSON function error occurred,error:%s", err))

		}
		return string(marshal)
	}
	return ""
}

func (v *Visitor) classObject2Map(classObject *stack.ClassObject) map[string]interface{} {
	data := make(map[string]interface{})
	for variable, val := range classObject.AllField() {
		switch val.(type) {
		case *stack.ClassObject:
			// class 包含 class 的情况
			classObject := val.(*stack.ClassObject)
			classObject2Map := v.classObject2Map(classObject)
			data[variable.GetName()] = classObject2Map
		default:
			data[variable.GetName()] = val
		}

	}
	return data
}

func (v *Visitor) JSONGet(ctx *parser.FunctionCallContext) interface{} {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	p1 := paramValues[1]
	var (
		v1, v2 string
	)
	switch p0.(type) {
	case *LeftValue:
		value := p0.(*LeftValue).GetValue()
		switch value.(type) {
		case string:
			v1 = fmt.Sprintf("%s", value)
		default:
			log.RuntimePanic(ctx, fmt.Sprintf("JSONGet JSON parameter are not a string: %v", reflect.TypeOf(value)))
		}
	case string:
		v1 = fmt.Sprintf("%s", p0)
	}
	switch p1.(type) {
	case *LeftValue:
		value := p1.(*LeftValue).GetValue()
		switch value.(type) {
		case string:
			v2 = fmt.Sprintf("%s", value)
		default:
			log.RuntimePanic(ctx, fmt.Sprintf("JSONGet path parameter are not a string: %v", reflect.TypeOf(value)))
		}
	case string:
		v2 = fmt.Sprintf("%s", p1)
	}
	return xjson.Get(v1, v2).Raw()

}

func (v *Visitor) getCurrentTime(ctx *parser.FunctionCallContext) string {
	tz, layout := v.getTzAndLayout(ctx)

	location, err := time.LoadLocation(tz)
	if err != nil {
		log.RuntimePanic(ctx, fmt.Sprintf("getCurrentTime function error occurred,error:%s", err))
	}
	local := time.Now().In(location)
	return local.Format(layout)

}
func (v *Visitor) unix(ctx *parser.FunctionCallContext) int64 {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	var (
		tz string
	)
	switch p0.(type) {
	case string:
		tz = fmt.Sprintf("%s", p0)
	}
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.RuntimePanic(ctx, fmt.Sprintf("unix function error occurred,error:%s", err))
	}
	local := time.Now().In(location)
	return local.Unix()

}

func (v *Visitor) getTzAndLayout(ctx *parser.FunctionCallContext) (string, string) {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	p1 := paramValues[1]
	var (
		tz, layout string
	)
	switch p0.(type) {
	case string:
		tz = fmt.Sprintf("%s", p0)
	}
	switch p1.(type) {
	case string:
		layout = fmt.Sprintf("%s", p1)
	}
	return tz, layout
}

func (v *Visitor) getOSArgs(ctx *parser.FunctionCallContext) []string {
	return Args
}

// 执行 command
func (v *Visitor) command(ctx *parser.FunctionCallContext) string {
	command, variableParams := v.getPrintfParams(ctx)
	var args []string
	if len(variableParams) > 0 {
		params := variableParams[0]
		strings, ok := params.([]interface{})
		if ok {
			for _, s := range strings {
				args = append(args, s.(string))
			}
		}
	}

	cmd := exec.Command(command, args...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.RuntimePanic(ctx, fmt.Sprintf("system.command function error occurred,error:%s, msg:%s", err, stdoutStderr))
	}
	return string(stdoutStderr)
}

func (v *Visitor) writeFile(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	p1 := paramValues[1]
	p2 := paramValues[2]
	fileName := p0.(string)
	value := p1.(string)
	perm := p2.(int)
	err := os.WriteFile(fileName, []byte(value), fs.FileMode(perm))
	if err != nil {
		log.RuntimePanic(ctx, fmt.Sprintf("system.writeFile function error occurred,error:%s", err))
	}
}
func (v *Visitor) remove(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	fileName := p0.(string)
	err := os.Remove(fileName)
	if err != nil {
		log.RuntimePanic(ctx, fmt.Sprintf("system.remove function error occurred,error:%s", err))
	}
}

func (v *Visitor) buildAppendParamValuesReturnLeft(ctx *parser.FunctionCallContext) ([]interface{}, *LeftValue) {
	ret := make([]interface{}, 0)
	if ctx.ExpressionList() == nil {
		return ret, nil
	}
	var left *LeftValue
	for _, context := range ctx.ExpressionList().(*parser.ExpressionListContext).AllExpr() {
		value := v.Visit(context)
		switch value.(type) {
		case *LeftValue:
			leftValue := value.(*LeftValue)
			if left == nil {
				// 只需要赋值一次左值即可，避免 append(buf, temp); temp 也是变量时，第二次遍历时会把 temp 赋值给 left
				left = leftValue
			}
			ret = append(ret, leftValue.GetValue())
		default:
			ret = append(ret, value)
		}
	}
	return ret, left
}

func (v *Visitor) printf(ctx *parser.FunctionCallContext) {
	format, variableParams := v.getPrintfParams(ctx)

	fmt.Printf(format, variableParams...)
}

// 获取格式化字符串参数
func (v *Visitor) getPrintfParams(ctx *parser.FunctionCallContext) (string, []interface{}) {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	var format string
	switch p0.(type) {
	case string:
		format = p0.(string)
	case *LeftValue:
		value := p0.(*LeftValue).GetValue()
		switch value.(type) {
		case string:
			format = value.(string)
		}
	}

	variableParams := paramValues[1:]
	return format, variableParams
}

func (v *Visitor) sprintf(ctx *parser.FunctionCallContext) string {
	format, variableParams := v.getPrintfParams(ctx)
	return fmt.Sprintf(format, variableParams...)
}
func (v *Visitor) print(ctx *parser.FunctionCallContext) {
	paramValues := v.buildParamValues(ctx)
	fmt.Print(paramValues...)
}
func (v *Visitor) dumpAST(ctx *parser.FunctionCallContext) string {
	code := v.getCode(ctx)
	return NewCompiler().GetCompileInfo(code, true)
}
func (v *Visitor) dumpSymbol(ctx *parser.FunctionCallContext) string {
	code := v.getCode(ctx)
	return NewCompiler().GetCompileInfo(code, false)
}

func (v *Visitor) getCode(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	var format string
	switch p0.(type) {
	case string:
		format = p0.(string)
	}
	return format
}

func (v *Visitor) toByteArray(ctx *parser.FunctionCallContext) []byte {
	code := v.getCode(ctx)
	return []byte(code)
}

func (v *Visitor) toString(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	p0 := paramValues[0]
	switch p0.(type) {
	case []byte:
		b := p0.([]byte)
		return string(b)
	case []interface{}:
		b, ok := p0.([]interface{})
		if !ok {
			return ""
		}
		var byteArray []byte
		for _, a := range b {
			byteArray = append(byteArray, a.([]byte)...)
		}
		return string(byteArray)
	}
	return ""
}

func (v *Visitor) getWd(ctx *parser.FunctionCallContext) string {
	str, err := os.Getwd()
	if err != nil {
		log.RuntimePanic(ctx, fmt.Sprintf("system.getwd function error occurred,error:%s", err))
	}
	return str
}
