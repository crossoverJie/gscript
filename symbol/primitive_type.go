package symbol

var (
	Int    = &PrimitiveType{name: "int"}
	String = &PrimitiveType{name: "string"}
	Float  = &PrimitiveType{name: "float"}
	Bool   = &PrimitiveType{name: "bool"}
	Void   = &VoidType{}
)

type PrimitiveType struct {
	name string
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
