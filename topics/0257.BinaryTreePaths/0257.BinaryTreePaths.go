package leetcode

import (
	"fmt"

	"github.com/rosenlo/leetcode/structure/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = tree.TreeNode

func travel(node *TreeNode, paths []string, path string) []string {
	if node == nil {
		return paths
	}

	if path != "" {
		path = fmt.Sprintf("%s->%d", path, node.Val)
	} else {
		path = fmt.Sprintf("%d", node.Val)
	}

	if node.Left == nil && node.Right == nil {
		paths = append(paths, path)
	}

	paths = travel(node.Left, paths, path)

	paths = travel(node.Right, paths, path)

	return paths

}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	return travel(root, []string{}, "")

}
