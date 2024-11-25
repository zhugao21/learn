package code

//type graphNode struct {
//	key      string
//	nodeList []string
//}
//
//func hasCycle(graph []graphNode) bool {
//	var nodeMap = make(map[string]graphNode, len(graph))
//	for _, gr := range graph {
//		nodeMap[gr.key] = gr
//	}
//	var visit = make(map[string]struct{}, len(graph))
//	for _, gr := range graph {
//		if checkNodeCycleByDFS(gr, nodeMap, visit) {
//			return true
//		}
//	}
//	return false
//}
//
//func checkNodeCycleByDFS(node graphNode, nodeMap map[string]graphNode, visit map[string]struct{}) bool {
//	if _, ok := visit[node.key]; ok {
//		return true
//	} else {
//		visit[node.key] = struct{}{}
//	}
//
//	for _, n := range node.nodeList {
//		if checkNodeCycleByDFS(nodeMap[n], nodeMap, visit) {
//			return true
//		}
//		delete(visit, n)
//	}
//	return false
//}

type GraphNode struct {
	Key   string
	Nodes []*GraphNode
}

func HasCycle(nodes []*GraphNode) bool {
	var visit = make(map[string]struct{})
	for _, node := range nodes {
		//if _, ok := visit[node.Key]; ok {
		//	return true
		//}
		if NodeHasCycle(node, visit) {
			return true
		}
	}
	return false
}

func NodeHasCycle(node *GraphNode, visit map[string]struct{}) bool {
	if _, ok := visit[node.Key]; ok {
		return true
	}
	visit[node.Key] = struct{}{}
	for _, nextNode := range node.Nodes {
		if NodeHasCycle(nextNode, visit) {
			return true
		}
	}
	delete(visit, node.Key)
	return false
}
