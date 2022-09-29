package gscript

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestExec(t *testing.T) {
	cmd := exec.Command("./gscript", "example/print_triangle.gs")
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Println(string(stdoutStderr), err)

	fileName := "temp.gs"
	script := "println(123);"
	err = os.WriteFile(fileName, []byte(script), 0666)
	fmt.Println(err)
	defer os.Remove(fileName)
}

func Test_Command(t *testing.T) {
	Args = os.Args
	script := `
System s = System();
string v = s.command("ls","-l","-h");
println(v);

v = s.command("./gscript", "example/print_triangle.gs");
println(v);
`
	NewCompiler().Compiler(script)
}
func Test_Write(t *testing.T) {
	Args = os.Args
	script := `
System s = System();
string fileName = "test.gs";
s.writeFile(fileName,"println(^hello^);",438);
string v = s.command("./gscript",fileName);
println(v);
s.remove(fileName);
`
	NewCompiler().Compiler(script)
}
