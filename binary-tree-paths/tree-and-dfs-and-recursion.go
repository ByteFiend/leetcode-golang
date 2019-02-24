/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func helper(root *TreeNode, path string, ans []string) []string {
	if nil == root {
		return ans
	}

	path += "->" + strconv.Itoa(root.Val)
	if nil == root.Left && nil == root.Right {
		ans = append(ans, path[2 : len(path)])
		return ans
	}

	ans = helper(root.Left, path, ans)
	ans = helper(root.Right, path, ans)
	return ans
}

func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	if nil == root {
		return ans
	}

	path := ""
	ans = helper(root, path, ans)
	return ans
}