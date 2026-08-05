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

	leftNode, leftBalanced := dfs(root.Left)
	rightNode, rightBalanced := dfs(root.Right)

	diff := leftNode - rightNode 
	isCurrentlyBalanced := diff <= 1 && diff >= -1
	currentHeight := 1 + max(leftNode, rightNode)
	isBalanced := isCurrentlyBalanced && rightBalanced && leftBalanced
	
	return currentHeight, isBalanced
}