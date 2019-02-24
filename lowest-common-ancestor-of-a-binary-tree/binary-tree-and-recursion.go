/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 * 题目假设 p q 一定在 root 这棵树上，所以没有做相关的判断
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if nil == root || nil == p || nil == q {
		return nil
	}
	if p == q {
		return p
	}
	if root == p || root == q {
		return root
	}

	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	if nil == left {
		return right
	} else if nil == right {
		return left
	} else {
		return root
	}
}