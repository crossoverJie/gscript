package symbol

var (
	Int    = &PrimitiveType{name: "int"}
	String = &PrimitiveType{name: "string"}
	Float  = &PrimitiveType{name: "float"}
	Bool   = &PrimitiveType{name: "bool"}
	Void   = &VoidType{}
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
	return v == t
}

func GetUpperType(t1, t2 Type) Type {
	if t1 == String || t2 == String {
		return String
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
