/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if nil == root || nil == root.Left && nil == root.Right {
		return 0
	}

	if nil == root.Left {
		return sumOfLeftLeaves(root.Right)
	}

	ans := 0
	if nil == root.Left.Left && nil == root.Left.Right {
		ans += root.Left.Val
	} else {
		ans += sumOfLeftLeaves(root.Left)
	}

	ans += sumOfLeftLeaves(root.Right)
	return ans
}