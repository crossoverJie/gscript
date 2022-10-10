package gscript

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestByte(t *testing.T) {
	script := `
string v5="9898";
`
	//os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestByte2(t *testing.T) {
	script := `
int[] a={1,2};
println(a);
`
	os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestByte3(t *testing.T) {
	script := `
string s = "10";
byte[] a= toByteArray(s);
printf("a=%v ",a);
string s1 = toString(a);
printf("s1=%s",s1);
assertEqual(s1,s);
`
	os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
	s := "10"
	b := []byte(s)

	s1 := string(b)
	assert.Equal(t, s, s1)
}
