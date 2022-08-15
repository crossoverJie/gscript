package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/crossoverJie/gscript"
	"github.com/crossoverJie/gscript/syntax"
	"os"
)

func main() {
	s := flag.String("x", "run", "debug or run, the debug mode will print the AST tree.")
	flag.Parse()
	if *s == "debug" {
		syntax.InitRuntime(true)
	} else {
		syntax.InitRuntime(false)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(syntax.Logo)
	var script bytes.Buffer

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		script.WriteString(text)
		ret := gscript.NewCompiler().Compiler(script.String())
		fmt.Println(ret)

	}
}
