package main

import (
	"bufio"
	"fmt"
	"github.com/crossoverJie/gscript/syntax"
	"os"
)

func main() {
	syntax.InitRuntime(false)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("gscript")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		root, err := syntax.Parse(text)
		if err != nil {
			fmt.Println(err)
		} else {
			syntax.EvaluateWithRuntime(root, "")
		}

	}
}
