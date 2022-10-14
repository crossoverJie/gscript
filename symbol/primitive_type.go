package symbol

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

const (
	PrimitiveInt    = "int"
	PrimitiveString = "string"
	PrimitiveFloat  = "float"
	PrimitiveBool   = "bool"
	PrimitiveNil    = "nil"
	PrimitiveAny    = "any"
	PrimitiveByte   = "byte"
)

var (
	Int    = &PrimitiveType{name: PrimitiveInt}
	String = &PrimitiveType{name: PrimitiveString}
	Float  = &PrimitiveType{name: PrimitiveFloat}
	Bool   = &PrimitiveType{name: PrimitiveBool}
	Byte   = &PrimitiveType{name: PrimitiveByte}
	Void   = &VoidType{}
	Nil    = &PrimitiveType{name: PrimitiveNil}
	Any    = &PrimitiveType{name: PrimitiveAny}
)

type PrimitiveType struct {
	name    string
	isArray bool
}

func (b *PrimitiveType) GetName() string {
	return b.name
}

func (b *PrimitiveType) GetEncloseScope() Scope {
	return nil
}

func (b *PrimitiveType) IsType(t Type) bool {
	if b.name == PrimitiveAny || t.GetName() == PrimitiveAny {
		return true
	}
	return b == t
}

func (b *PrimitiveType) String() string {
	return b.name
}

func (b *PrimitiveType) IsArray() bool {
	// todo crossoverJie 数组校验
	return b.isArray
}
func (b *PrimitiveType) SetArray(isArray bool) {
	b.isArray = isArray
}

type VoidType struct {
}

func (v *VoidType) GetName() string {
	return "void"
}

func (v *VoidType) GetEncloseScope() Scope {
	return nil
}

func (v *VoidType) IsType(t Type) bool {
	if MatchNil(t) {
		return true
	}
	return v == t
}

func (v *VoidType) IsArray() bool {
	return false
}
func (v *VoidType) SetArray(isArray bool) {
}

// GetUpperType 根据两个参数类型，推导返回的类型
func GetUpperType(ctx antlr.ParserRuleContext, t1, t2 Type) Type {
	if t1 == String || t2 == String {
		return String
	} else if t1 == Nil || t2 == Nil {
		return Nil
	} else if t1 == Any || t2 == Any {
		if t1 != Any {
			return t1
		}
		if t2 != Any {
			return t2
		}
		// 都是 any，则需要在运行时判断
		return Any
	} else if t1 == String && t2 == String {
		return String
	} else if t1 == Int && t2 == Int {
		return Int
	} else if t1 == Float && t2 == Float {
		return Float
	} else if t1 == Int && t2 == Float {
		return Float
	} else if t1 == Float && t2 == Int {
		return Float
	} else {
		// todo crossoverJie ?
		return nil
	}
}

// GetUpperTypeWithValue 根据两个参数值，返回推到类型
func GetUpperTypeWithValue(ctx antlr.ParserRuleContext, v1, v2 interface{}) Type {
	var (
		v1Type, v2Type Type
	)
	v1Type = getTypeWithValue(v1)
	v2Type = getTypeWithValue(v2)
	return GetUpperType(ctx, v1Type, v2Type)
}

// 根据值返回类型
func getTypeWithValue(v interface{}) Type {
	var t Type
	switch v.(type) {
	case int:
		t = Int
	case string:
		t = String
	case float64:
		t = Float
	case bool:
		t = Bool
	default:
		t = Nil
	}
	return t
}

func Value2Float(v interface{}) float64 {
	switch v.(type) {
	case float64:
		return v.(float64)
	case int:
		return float64(v.(int))
	default:
		return 0
	}
}

// GetType 运行时根据实际值推算出类型
/**
if( n[0] + n[1] == 3){
	println(3);
}
推断出  n[0] + n[1] 的 type 为 int，因为这个表达式 type 在编译器计算不出来
*/
func GetType(t1, t2 Type, left, right interface{}) (Type, Type) {
	var (
		type1, type2 = t1, t2
	)
	if t1 == nil {
		switch left.(type) {
		case string:
			type1 = String
		case int:
			type1 = Int
		case float64:
			type1 = Float
		case bool:
			type1 = Bool
		}
	}

	if t2 == nil {
		switch right.(type) {
		case string:
			type2 = String
		case int:
			type2 = Int
		case float64:
			type2 = Float
		case bool:
			type2 = Bool
		}
	}
	return type1, type2
}
