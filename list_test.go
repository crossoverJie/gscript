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
