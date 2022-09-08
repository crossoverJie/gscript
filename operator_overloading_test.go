package gscript

import "testing"

func TestOperator(t *testing.T) {
	script := `
class Person{
	int age;
	Person(int a){
		age = a;
	}
}

Person operator + (Person p1, Person p2){
	Person pp = Person(p1.age+p2.age);
	return pp;
}
Person operator - (Person p1, Person p2){
	Person pp = Person(p1.age-p2.age);
	return pp;
}
Person p1 = Person(10);
Person p2 = Person(20);
//Person p3 =  operator(p1,p2);
Person p3 = p1+p2;
println("p3.age="+p3.age);

Person p4 = p1-p2;
println("p4.age="+p4.age);
`
	NewCompiler().CompilerWithoutNative(script)
}
