package code

type graphNode struct {
	key      string
	nodeList []string
}

func hasCycle(graph []graphNode) bool {
	var nodeMap = make(map[string]graphNode, len(graph))
	for _, gr := range graph {
		nodeMap[gr.key] = gr
	}
	var visit = make(map[string]struct{}, len(graph))
	for _, gr := range graph {
		if checkNodeCycleByDFS(gr, nodeMap, visit) {
			return true
		}
	}
	return false
}

func checkNodeCycleByDFS(node graphNode, nodeMap map[string]graphNode, visit map[string]struct{}) bool {
	if _, ok := visit[node.key]; ok {
		return true
	} else {
		visit[node.key] = struct{}{}
	}

	for _, n := range node.nodeList {
		if checkNodeCycleByDFS(nodeMap[n], nodeMap, visit) {
			return true
		}
		delete(visit, n)
	}
	return false
}
