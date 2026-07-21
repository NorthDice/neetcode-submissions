/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	oldToNew:=make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if v, ok := oldToNew[node];ok {
			return v
		}

		cp := &Node{Val:node.Val}

		oldToNew[node] = cp

		for _, n := range node.Neighbors {
			cp.Neighbors = append(cp.Neighbors,dfs(n))
		}
		return cp
	}
	return dfs(node)
}
