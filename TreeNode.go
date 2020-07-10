package dailyleetcode

import (
	"fmt"
	"strings"
)

// TreeNode  sample struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(in string) *TreeNode {
	var root *TreeNode
	if in[0] == '[' && in[len(in)-1] == ']' {
		in = in[1 : len(in)-1]
	}
	for _, s := range strings.Split(in, ",") {
		s = strings.Trim(s, " ")
		fmt.Println(s)
	}

	return root
}
