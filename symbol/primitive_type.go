package symbol

const (
	PrimitiveInt    = "int"
	PrimitiveString = "string"
	PrimitiveFloat  = "float"
	PrimitiveBool   = "bool"
	PrimitiveNil    = "nil"
	PrimitiveAny    = "any"
)

var (
	Int    = &PrimitiveType{name: PrimitiveInt}
	String = &PrimitiveType{name: PrimitiveString}
	Float  = &PrimitiveType{name: PrimitiveFloat}
	Bool   = &PrimitiveType{name: PrimitiveBool}
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
	if b.name == PrimitiveAny {
		return true
	}
	return b == t
}

func (b *PrimitiveType) String() string {
	return b.name
}

func (b *PrimitiveType) IsArray() bool {
	// todo crossoverJie 数组校验
	return false
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

// GetUpperType 根据两个参数类型，推导返回的类型
func GetUpperType(t1, t2 Type) Type {
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
		// todo crossoverJie 记录异常
		return nil
	}
}

// GetUpperTypeWithValue 根据两个参数值，返回推到类型
func GetUpperTypeWithValue(v1, v2 interface{}) Type {
	var (
		v1Type, v2Type Type
	)
	v1Type = getTypeWithValue(v1)
	v2Type = getTypeWithValue(v2)
	return GetUpperType(v1Type, v2Type)
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
