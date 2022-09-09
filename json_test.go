package gscript

import (
	"encoding/json"
	"fmt"
	"github.com/crossoverJie/xjson"
	"testing"
)

func TestJSON(t *testing.T) {
	script := `
class Person{
	int age;
	string name;
	float weight;
	bool man;
	Person(string n, int a, float w, bool m){
		name = n;
		age = a;
		weight = w;
		man =m;
	}
}
Person p1 = Person("abc",10,99.99,true);
Person p2 = Person("a",11,999.99,false);
string json = JSON(p1);
println(json);

int[] a = {1,2,3};
json = JSON(a);
println(json);

int b =10;
json = JSON(b);
println(json);

Person[] list = {p1,p2};
println(list);
Person temp = list[0];
println(temp.name);
json = JSON(list);
//[{"age":10,"man":true,"name":"abc","weight":99.99},{"age":11,"man":false,"name":"a","weight":999.99}]
println(json);

`
	NewCompiler().CompilerWithoutNative(script)
}
func TestJSONGet(t *testing.T) {
	script := `
class Person{
	int age;
	string name;
	float weight;
	bool man;
	Person(string n, int a, float w, bool m){
		name = n;
		age = a;
		weight = w;
		man =m;
	}
}
Person p1 = Person("abc",10,99.99,true);
string json = JSON(p1);
println(json);

int age = JSONGet(json, "age");
println(age);
assertEqual(age,10);

string man = JSONGet(json, "man");
println(man);
assertEqual(man,true);

string name = JSONGet(json, "name");
println(name);
assertEqual(name,"abc");

float weight = JSONGet(json, "weight");
println(weight);
assertEqual(weight,99.99);

string j=^{"age":10, "abc":{"def":"def"},"list":[1,2,3]}^;
String def = JSONGet(j, "abc.def");
println(def);
assertEqual(def,"def");
int l1 = JSONGet(j, "list[0]");
println(l1);
assertEqual(l1,1);


string str=^
{
    "name": "bob",
    "age": 20,
    "skill": {
        "lang": [
            {
                "go": {
                    "feature": [
                        "goroutine",
                        "channel",
                        "simple",
                        true
                    ]
                }
            }
        ]
    }
}
^;
String g = JSONGet(str, "skill.lang[0].go.feature[0]");
println(g);
assertEqual(g,"goroutine");
`
	NewCompiler().Compiler(script)
}

func TestJSON1(t *testing.T) {
	a := 10
	marshal, err := json.Marshal(a)
	fmt.Println(string(marshal), err)

	s1 := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{Name: "a", Age: 100}
	s2 := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{Name: "a", Age: 10}

	list := []interface{}{s1, s2}
	bytes, err := json.Marshal(list)
	fmt.Println(string(bytes), err)
	fmt.Println(list)
}

func TestVisitor_JSONGet(t *testing.T) {
	str := "{\"age\":10, \"abc\":{\"def\":\"def\"}}"
	get := xjson.Get(str, "abc.def")
	fmt.Println(get.String())
}
func TestVisitor_JSONGet2(t *testing.T) {
	script := `
string s="1";
println(s);
string s1=^123^;
println(s1);
`
	NewCompiler().CompilerWithoutNative(script)
}
func TestVisitor_JSONGet3(t *testing.T) {
	script := `
class P{
	string name;
	P(string n){
		name = n;
	}
}
class Object{
	P p;
	int x;
	Object(P pp, int xx){
		p = pp;
		x = xx;
	}
}
P p1 = P("abc");
Object o1 = Object(p1, 100);
string json = JSON(o1);
println(json);

int x = JSONGet(json,"x");
println(x);
assertEqual(x,100);

string name = JSONGet(json,"p.name");
println(name);
assertEqual(name,"abc");

`
	NewCompiler().Compiler(script)
}
