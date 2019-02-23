/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}

	if nil == root.Left && nil == root.Right {
		return 1
	} else if nil == root.Left {
		return minDepth(root.Right) + 1
	} else if nil == root.Right {
		return minDepth(root.Left) + 1
	}

	left, right := minDepth(root.Left), minDepth(root.Right)
	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}