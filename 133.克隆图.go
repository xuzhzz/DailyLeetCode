package dailyleetcode

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return node
	}
	if n, ok := visited[node]; ok {
		return n
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i, neighbor := range node.Neighbors {
		newNode.Neighbors[i] = clone(neighbor, visited)
	}
	return newNode
}

// @lc code=end
