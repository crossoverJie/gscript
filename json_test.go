package gscript

import (
	"encoding/json"
	"fmt"
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
