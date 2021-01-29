package leetcode

import "github.com/rosenlo/leetcode/util"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = util.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if p.Left == nil && q.Left == nil {
		return true
	}

	if p.Right == nil && q.Right == nil {
		return true
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
