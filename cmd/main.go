package main

import (
	"bufio"
	"flag"
	"fmt"
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

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		root, err := syntax.Parse(text)
		if err != nil {
			fmt.Println(err)
		} else {
			if syntax.GetRuntime().Verbose() {
				root.Print("├")
			}
			syntax.EvaluateWithRuntime(root, "")
		}

	}
}
