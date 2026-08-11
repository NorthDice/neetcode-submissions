/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	
	_, balanced := dfs(root)
	return balanced
	
}

func dfs(root *TreeNode) (int, bool){
	if root == nil {
		return 0, true
	}

	leftNode,leftBalanced := dfs(root.Left)
	rightNode, rightBalanced := dfs(root.Right)
	diff := leftNode - rightNode >= -1 && leftNode - rightNode <= 1
	isBalanced := diff && leftBalanced && rightBalanced
	currentHeight := 1 + max(leftNode,rightNode)
	
	return currentHeight, isBalanced
}