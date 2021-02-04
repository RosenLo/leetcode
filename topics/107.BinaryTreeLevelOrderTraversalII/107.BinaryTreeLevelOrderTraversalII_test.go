package leetcode

import (
	"reflect"
	"testing"

	"github.com/rosenlo/leetcode/util"
)

func TestLevelOrderBottom(t *testing.T) {
	tests := [][]int{
		{3, 9, 20, util.NULL, util.NULL, 15, 7},
		{1},
		{},
	}
	results := [][][]int{
		{
			{15, 7}, {9, 20}, {3},
		},
		{
			{1},
		},
		{},
	}
	for i := 0; i < len(tests); i++ {
		root := util.Ints2TreeNode(tests[i])
		ret := levelOrderBottom(root)
		if !reflect.DeepEqual(results[i], ret) {
			t.Fatalf("Wrong Answer, testcase: %v, actual: %v expected: %v", tests[i], ret, results[i])
		}
	}
}
