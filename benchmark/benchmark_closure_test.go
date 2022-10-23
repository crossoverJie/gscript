package benchmark

import (
	"github.com/crossoverJie/gscript"
	"testing"
)

func BenchmarkClosure(b *testing.B) {
	script := `
class Test{
	int value=0;
	Test(int v){
		value=v;
	}

	int map(func int(int) f){
		return f(value);
	}
}
int square(int v){
	return v*v; 
}
int add(int v){
	return v++; 
}
int add2(int v){
	v=v+2;
	return v; 
}
Test t =Test(100);
func int(int) s=square;
func int(int) a=add;
func int(int) a2=add2;
//println(t.map(s));
assertEqual(t.map(s),10000);

//println(t.map(a));
assertEqual(t.map(a),101);

//println(t.map(a2));
assertEqual(t.map(a2),102);
`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		gscript.NewCompiler().Compiler(script)
	}
}
