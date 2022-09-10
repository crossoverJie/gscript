package gscript

import (
	"encoding/json"
	"fmt"
	"github.com/crossoverJie/gscript/parser"
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/xjson"
	"hash/fnv"
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
	if len(paramValues) != 2 {
		// todo crossoverJie 编译器报错
		panic("")
	}
	// todo crossoverJie 参数是个变量，需要取左值
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

func (v *Visitor) JSON(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 1 {
		// todo crossoverJie 运行时报错
		panic("")
	}
	value := paramValues[0]
	switch value.(type) {
	case *stack.ClassObject:
		classObject := value.(*stack.ClassObject)
		data := v.classObject2Map(classObject)
		marshal, err := json.Marshal(data)
		if err != nil {
			// todo crossoverJie 运行时报错
			panic("")
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
				// todo crossoverJie 运行时报错
				panic("")
			}
			return string(marshal)
		}
		// int[] a = {1,2,3}; json = JSON(a)
		marshal, err := json.Marshal(dataList)
		if err != nil {
			// todo crossoverJie 运行时报错
			panic("")
		}
		// {1,2,3}
		return string(marshal)
	default:
		marshal, err := json.Marshal(value)
		if err != nil {
			// todo crossoverJie 运行时报错
			panic("")
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
	if len(paramValues) != 2 {
		// todo crossoverJie 编译器报错
		panic("")
	}
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
			// todo crossoverJie 编译器报错
			panic("")
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
			// todo crossoverJie 编译器报错
			panic("")
		}
	case string:
		v2 = fmt.Sprintf("%s", p1)
	}
	return xjson.Get(v1, v2).Raw()

}

func (v *Visitor) getCurrentTime(ctx *parser.FunctionCallContext) string {
	paramValues := v.buildParamValues(ctx)
	if len(paramValues) != 2 {
		// todo crossoverJie 运行时报错
		panic("")
	}
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

	location, err := time.LoadLocation(tz)
	if err != nil {
		// todo crossoverJie 运行时报错
		panic(err)
	}
	local := time.Now().In(location)
	return local.Format(layout)

}
