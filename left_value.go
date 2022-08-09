package gscript

import (
	"github.com/crossoverJie/gscript/stack"
	"github.com/crossoverJie/gscript/symbol"
)

type LeftValue struct {
	variable *symbol.Variable
	object   stack.Object
}
