package code

import "testing"

func Test_graphHasCycle(t *testing.T) {
	node1 := graphNode{key: "a", nodeList: []string{"b", "c"}}
	node2 := graphNode{key: "b", nodeList: []string{"c"}}
	node3 := graphNode{key: "c", nodeList: []string{"d"}}
	node4 := graphNode{key: "d", nodeList: []string{}}

	t.Log(hasCycle([]graphNode{node1, node2, node3, node4}))
}
