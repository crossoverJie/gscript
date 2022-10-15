package gscript

import (
	"fmt"
	"strings"
	"testing"
)

func TestByte(t *testing.T) {
	script := `
string v5="9898";
byte b=1;
println(b);
`
	//os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestByte2(t *testing.T) {
	script := `
int[] a={1,2};
println(a);
`
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
	NewCompiler().Compiler(script)
}
func TestStringBuilder3(t *testing.T) {
	script := `
Strings s = Strings();
string[] elems = {"name=xxx","age=xx"};
string ret = s.join(elems, "&");
println(ret);
assertEqual(ret, "name=xxx&age=xx");
`
	//os.Setenv(RuntimeError, "true")
	NewCompiler().Compiler(script)
}
func TestStringBuilder4(t *testing.T) {
	script := `
int[] a = [0]{1,2};
println(len(a));
println(cap(a));
append(a,3);
println(cap(a));
assertEqual(cap(a),4);
`
	NewCompiler().Compiler(script)
}
func TestStringBuilder5(t *testing.T) {
	script := `
StringBuilder sb = StringBuilder();
sb.grow(15);
`
	NewCompiler().Compiler(script)
}
func TestStringBuilder6(t *testing.T) {
	script := `
Strings s = Strings();
bool b = s.hasPrefix("http://www.xx.com", "http");
println(b);
assertEqual(b,true);
b = s.hasPrefix("http://www.xx.com", "https");
println(b);
assertEqual(b,false);
`
	NewCompiler().Compiler(script)
}
func TestArraySlice(t *testing.T) {
	script := `
int[] a = {1,2,3};
int s=1;
int[] b = a[s:len(a)];
println(b);
`
	NewCompiler().Compiler(script)
}
func TestArraySlice2(t *testing.T) {
	script := `
string[] a = {"1","2","3"};
string s="";
string[] b = a[1:len(a)];
println(b);
`
	NewCompiler().Compiler(script)
}
func TestArraySlice3(t *testing.T) {
	script := `
string S = "abcdefx";
string T = "def";
int myIndex(string S,string T,int pos) {
  int i = pos;
  int j = 0;
  byte[] s = toByteArray(S);
  byte[] t = toByteArray(T);

  int slen = len(s);
  int tlen = len(t);
  for(i <= slen && j <= tlen) {
	byte si = s[i];
	byte sj = s[j];
    if (si == sj) {
      i = i+1;
      j = j+1;
    }else{
      i = i - j + 1;
      j = 0;
    }
  }
  if (j == tlen) {
    return i - tlen;
  } else {
    return slen-i;
  }
}

int res = myIndex(S,T,0);
println(res);
`
	NewCompiler().Compiler(script)
}
func TestStringBuilder2(t *testing.T) {

	var x strings.Builder
	x.WriteString("10")
	x.WriteString("20")
	l, _ := x.WriteString("30")
	fmt.Println(l)
	//fmt.Println(strings.Join([]string{"name=xxx", "age=xx"}, "&"))

	a := []int{1, 2}
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 3)
	fmt.Println(cap(a))
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
