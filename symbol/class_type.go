package symbol

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Class struct {
	// todo crossoverJie 父类
	*scope
	defaultConstructorFunc *DefaultConstructorFunc
}

func NewClass(ctx antlr.ParserRuleContext, name string) *Class {
	return &Class{
		scope: &scope{
			symbol: &symbol{
				name: name,
				ctx:  ctx,
			},
		},
	}
}

func (c *Class) IsType(t Type) bool {
	// todo crossoverJie 父类判断
	if MatchNil(t) {
		return true
	}
	return c == t
}

func (c *Class) String() string {
	return "Class -> " + c.GetName()
}

// GetVariable 在 class 中查找变量
func (c *Class) GetVariable(name string) *Variable {
	// todo crossoverJie 查找父类
	return c.scope.GetVariable(name)
}

// GetFunction 查找类中的函数
func (c *Class) GetFunction(name string, paramTypes []Type) *Func {
	// todo crossoverJie 查找父类的函数
	return c.scope.GetFunction(name, paramTypes)
}

// GetClassFunctionVariable 查找类中的函数变量
func (c *Class) GetClassFunctionVariable(name string, paramTypes []Type) *Variable {
	// todo crossoverJie 查找父类的函数
	return c.scope.GetFunctionVariable(name, paramTypes)
}

func (c *Class) ContainsSymbol(symbol Symbol) bool {
	// todo crossoverJie 查找父类
	return c.scope.ContainsSymbol(symbol)
}

// GetConstructorFunc 查找显式的构造函数，函数名称就是类名称
func (c *Class) GetConstructorFunc(paramTypes []Type) *Func {
	return c.GetFunction(c.name, paramTypes)
}

func (c *Class) GetDefaultConstructorFunc() *DefaultConstructorFunc {
	if c.defaultConstructorFunc == nil {
		c.defaultConstructorFunc = NewDefaultConstructorFunc(c.name, c)
	}
	return c.defaultConstructorFunc
}

// AddSymbol 要实现改函数，不然走的是 scope 的函数
func (c *Class) AddSymbol(symbol Symbol) {
	c.symbols = append(c.symbols, symbol)
	symbol.SetEncloseScope(c)
}

type DefaultConstructorFunc struct {
	*scope
}

func NewDefaultConstructorFunc(name string, class *Class) *DefaultConstructorFunc {
	return &DefaultConstructorFunc{
		scope: &scope{
			symbol: &symbol{
				name:         name,
				encloseScope: class,
			},
		},
	}
}

// GetClass 获取构造函数对应的 class
func (d *DefaultConstructorFunc) GetClass() *Class {
	return d.GetEncloseScope().(*Class)
}
func (d *DefaultConstructorFunc) String() string {
	return "defaultConstructorFunc - >" + d.GetName()
}
