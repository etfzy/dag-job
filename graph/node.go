package graph

import "fmt"

type Node struct {
	name  string
	pres  map[string]*Node
	nexts map[string]*Node
}

func newNode(name string) *Node {
	return &Node{
		name:  name,
		pres:  make(map[string]*Node),
		nexts: make(map[string]*Node),
	}
}

func (n *Node) addPreNode(Node *Node) {
	n.pres[Node.name] = Node
}

func (n *Node) addNextNode(Node *Node) {
	n.nexts[Node.name] = Node
}

func (n *Node) GetName() string {
	return n.name
}

func (n *Node) GetPreNodes() map[string]*Node {
	return n.pres
}
func (n *Node) GetNextNodes() map[string]*Node {
	return n.nexts
}

// 检测整个图是否存在环（从任意节点出发）
func HasCycle(nodes map[string]*Node) error {
	visited := make(map[string]bool)        // 永久访问记录
	recursionStack := make(map[string]bool) // 当前递归路径

	for name := range nodes {
		if !visited[name] {
			if err := dfs(nodes[name], visited, recursionStack); err != nil {
				return err
			}
		}
	}
	return nil
}

func dfs(node *Node, visited, recursionStack map[string]bool) error {
	if recursionStack[node.name] {
		// 发现环：当前节点已在递归路径中
		return fmt.Errorf("cycle node:%s", node.name)
	}
	if visited[node.name] {
		// 已永久访问过，无需处理
		return nil
	}

	// 标记为当前路径访问
	visited[node.name] = true
	recursionStack[node.name] = true

	// 遍历所有后继节点（nexts）
	for _, next := range node.nexts {
		if err := dfs(next, visited, recursionStack); err != nil {
			return err
		}
	}

	// 回溯：移出当前路径
	recursionStack[node.name] = false
	return nil
}
