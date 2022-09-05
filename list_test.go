package gscript

import "testing"

func TestLinkedList(t *testing.T) {
	script := `
LinkedList list = LinkedList();
list.add("1");
list.add("2");
list.add("3");
list.dump();
print("size= "+ list.getSize());
assertEqual(list.getSize(), 3);
`
	NewCompiler().Compiler(script)
}
func TestLinkedLis2(t *testing.T) {
	script := `
int count =100;
LinkedList list = LinkedList();
for (int i=0;i<count;i++){
	string key = i+"";
	list.add(key);
}
list.dump();
assertEqual(list.getSize(), 100);
`
	NewCompiler().Compiler(script)
}
