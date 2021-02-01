package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	for i := 1; i < n; i++ {
		node := queue[0]
		queue = queue[1:]

		node.Left = &TreeNode{Val: ints[i]}
		queue = append(queue, node.Left)
		i++

		if i < n {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
	}

	return root
}
