package code

import "testing"

//func Test_graphHasCycle(t *testing.T) {
//	node1 := graphNode{key: "a", nodeList: []string{"b", "c", "d"}}
//	node2 := graphNode{key: "b", nodeList: []string{}}
//	node3 := graphNode{key: "c", nodeList: []string{}}
//	node4 := graphNode{key: "d", nodeList: []string{"c"}}
//
//	t.Log(hasCycle([]graphNode{node1, node2, node3, node4}))
//}

func Test_GraphHasCycle(t *testing.T) {
	a := &GraphNode{Key: "A"}
	b := &GraphNode{Key: "B"}
	c := &GraphNode{Key: "C"}
	d := &GraphNode{Key: "D"}

	a.Nodes = []*GraphNode{b, d}
	d.Nodes = []*GraphNode{c}
	//c.Nodes = []*GraphNode{a}

	t.Log(HasCycle([]*GraphNode{a, b, c, d}))
}
