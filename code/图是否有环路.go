package code

//type GraphNode struct {
//	Key   string
//	Nodes []*GraphNode
//}
//
//func HasCycle(nodes []*GraphNode) bool {
//	var visit = make(map[string]bool, len(nodes))
//
//	for _, node := range nodes {
//		if !visit[node.Key] && nodeHasCycle(node, visit) {
//			return true
//		}
//	}
//
//	return false
//}
//
//func nodeHasCycle(node *GraphNode, visit map[string]bool) bool {
//	if visit[node.Key] {
//		return true
//	}
//	visit[node.Key] = true
//	for _, n := range node.Nodes {
//		if nodeHasCycle(n, visit) {
//			return true
//		}
//	}
//	delete(visit, node.Key)
//	return false
//}

type GraphNode struct {
	Key   string
	Nodes []*GraphNode
}

func HasCycle(nodes []*GraphNode) bool {
	var visit = make(map[string]bool, len(nodes))
	for _, node := range nodes {
		if !visit[node.Key] && nodeHasCycle(node, visit) {
			return true
		}
	}
	return false
}

func nodeHasCycle(node *GraphNode, visit map[string]bool) bool {
	if visit[node.Key] {
		return true
	}
	visit[node.Key] = true
	for _, n := range node.Nodes {
		if nodeHasCycle(n, visit) {
			return true
		}
	}
	visit[node.Key] = false
	return false
}
