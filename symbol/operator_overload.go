package symbol

const OperName = "operator"

// OpOverload 重载符
type OpOverload struct {
	function  *Func
	tokenType int
}

func NewOpOverload(function *Func, tokenType int) *OpOverload {
	return &OpOverload{
		function:  function,
		tokenType: tokenType,
	}
}

func (o *OpOverload) GetFunc() *Func {
	return o.function
}

func (o *OpOverload) GetTokenType() int {
	return o.tokenType
}
