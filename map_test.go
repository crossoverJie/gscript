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
print(m.str());
print(m.getValue());

Map[] list = {};
list = append(list,m);
Map m2=Map();
m2.put("2","2");
list = append(list,m2);
print(list);
print(len(list));
`
	NewCompiler().Compiler(script)
}
func Test_map2(t *testing.T) {
	script := `
print(hash("123"));
print(hash(1));
print(hash(1.1));
print(hash(true));
print(hash(true));
`
	NewCompiler().Compiler(script)
}
func Test_map3(t *testing.T) {
	script := `
MapString m1 = MapString();
m1.put("1","10");
m1.put("2","20");
print("get 1=" + m1.get("1"));
print("get 2=" + m1.get("2"));
print("size=" + m1.getSize());
`
	NewCompiler().Compiler(script)
}
func Test_map4(t *testing.T) {
	script := `
class List{
	List next;
}
List list = List();
print(list.next);
if(list.next == nil){
	print("nil");
}
List next = List();
list.next = next;
if(list.next != nil){
	print("!nil");
}
`
	NewCompiler().Compiler(script)
}
