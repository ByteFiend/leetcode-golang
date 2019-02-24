/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(root *TreeNode, min, max int, minEnabled, maxEnabled bool) bool {
	if nil == root {
		return true
	}

	if minEnabled && root.Val <= min || maxEnabled && root.Val >= max {
		return false
	}

	return helper(root.Left, min, root.Val, minEnabled, true) && helper(root.Right, root.Val, max, true, maxEnabled)
}

func isValidBST(root *TreeNode) bool {
	return helper(root, 0, 0, false, false)
}