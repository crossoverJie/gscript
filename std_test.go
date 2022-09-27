package gscript

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExec(t *testing.T) {
	cmd := exec.Command("./gscript", "example/print_triangle.gs")
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Println(string(stdoutStderr), err)
}
