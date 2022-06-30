package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArithmeticOperators(t *testing.T) {
	i := ArithmeticOperators("1.1+2+3+4\n")
	assert.Equal(t, i, 10.1)
	operators := ArithmeticOperators("1+2+3+4\n")
	assert.Equal(t, operators, 10)

	assert.Equal(t, ArithmeticOperators("1+2+4*2\n"), 11)

	assert.Equal(t, ArithmeticOperators("1+(2+4)*2\n"), 13)
	assert.Equal(t, ArithmeticOperators("1+(2+4)*1.5\n"), 10.0)
	assert.Equal(t, ArithmeticOperators("(1+(2+4)*1.5)/10\n"), 1.0)

}
