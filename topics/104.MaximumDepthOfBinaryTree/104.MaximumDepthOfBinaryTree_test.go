package leetcode

import (
	"testing"

	"github.com/rosenlo/leetcode/util"
)

func TestMaxDepth(t *testing.T) {
	tests := [][]int{
		{3, 9, 20, util.NULL, util.NULL, 15, 7},
		{1, util.NULL, 2},
		{1},
		{},
		{
			3,
			4, 5,
			-7, -6, util.NULL, util.NULL,
			-7, util.NULL, -5, util.NULL, util.NULL, util.NULL, -4,
		},
		{
			1,
			2, 2,
			3, 3, util.NULL, util.NULL,
			util.NULL, 4,
			util.NULL, util.NULL, 5, 5,
		},
	}
	results := []int{
		3,
		2,
		1,
		0,
		5,
		5,
	}
	for i := 0; i < len(tests); i++ {
		root := util.Ints2TreeNode(tests[i])
		ret := maxDepth(root)
		if ret != results[i] {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret, results[i])
		}
	}
}
