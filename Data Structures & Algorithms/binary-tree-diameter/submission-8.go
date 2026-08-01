/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var dfs func (*TreeNode)int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		res = max(res,l+r)

		return 1 + max(dfs(root.Left), dfs(root.Right))
	}

	dfs(root)
	return res
}
