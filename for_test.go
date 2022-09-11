package gscript

import "testing"

func TestFor(t *testing.T) {
	script := `
int[] nums = {1,2,3};
for (int j = len(nums) -1 ;j > 0 ;j--){
	int r = nums[j];
	println(r);
	//println(j);
}
`
	NewCompiler().Compiler(script)
}
