package gscript

import (
	"testing"
)

func Test_map(t *testing.T) {
	script := `
class Map{
	string key, value;
	put(string k, string v){
		key=k;
		value=v;
	}
	string str(){
		return key+value;
	}
	string getValue(){
		return value;
	}
}
Map m=Map();
m.put("1","1");
println(m.str());
println(m.getValue());

Map[] list = {};
list = append(list,m);
Map m2=Map();
m2.put("2","2");
list = append(list,m2);
println(list);
println(len(list));
`
	NewCompiler().CompilerWithoutNative(script)
}
func Test_map2(t *testing.T) {
	script := `
println(hash("123"));
println(hash(1));
println(hash(1.1));
println(hash(true));
println(hash(true));
`
	NewCompiler().Compiler(script)
}
func Test_map3(t *testing.T) {
	script := `
MapString m1 = MapString();
m1.put("1","10");
m1.put("2","20");
m1.put("3","30");
println("get 1=" + m1.get("1"));
println("get 2=" + m1.get("2"));
println("get 3=" + m1.get("3"));
m1.put("3","300");
println("get 3=" + m1.get("3"));
println("size=" + m1.getSize());
assertEqual(m1.getSize(),3);
assertEqual(m1.get("1"),"10");
assertEqual(m1.get("2"),"20");
assertEqual(m1.get("3"),"300");
`
	NewCompiler().Compiler(script)
}
func Test_map4(t *testing.T) {
	script := `
class List{
	List next;
}
List list = List();
println(list.next);
if(list.next == nil){
	println("nil");
}
List next = List();
list.next = next;
if(list.next != nil){
	println("!nil");
}
`
	NewCompiler().Compiler(script)
}

func Test_map5(t *testing.T) {
	script := `
int count =100;
MapString m1 = MapString();
for (int i=0;i<count;i++){
	string key = i+"";
	string value = key;
	m1.put(key,value);
}
println(m1.getSize());
assertEqual(m1.getSize(),count);

for (int i=0;i<count;i++){
	string key = i+"";
	string value = m1.get(key);
	println("key="+key+ ":"+ value);
	assertEqual(key,value);
}
`
	NewCompiler().Compiler(script)
}
func Test_map6(t *testing.T) {
	script := `
int count =100;
MapString m1 = MapString();
m1.put("","");
println(m1.getSize());
assertEqual(m1.getSize(),0);
`
	NewCompiler().Compiler(script)
}
