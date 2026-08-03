/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return nil
	}
	res := make([][]int,0)
	queue:=[]*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]int,0,levelSize)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue,node.Right)
			}
			
			levelVals = append(levelVals, node.Val)
		}

		res = append(res, levelVals)	
	}

	return res
}
