package gscript

import (
	"fmt"
	"os"
	"strings"
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
println("");
`
	os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestStringBuilder(t *testing.T) {
	script := `
StringBuilder b = StringBuilder();
b.writeString("10");
b.writeString("20");
int l = b.writeString("30");
string s = b.String();
printf("s:%s, len=%d ",s,l);
assertEqual(s,"102030");
byte[] b2 = toByteArray("40");
b.WriteBytes(b2);
s = b.String();
assertEqual(s,"10203040");
println(s);
`
	os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestStringBuilder2(t *testing.T) {

	var x strings.Builder
	x.WriteString("10")
	x.WriteString("20")
	l, _ := x.WriteString("30")
	fmt.Println(l)
}

func TestNil2(t *testing.T) {
	var ii I
	m := make(map[string]I)
	m["0"] = ii
	var i I
	i = m["0"]
	if i == nil {
		fmt.Println("i==nil,return")
		return
	}
	fmt.Println("i != nil")
}

type I interface {
	I() string
}
type M struct {
	name string
}

func (m *M) I() string {
	return "123"
}
