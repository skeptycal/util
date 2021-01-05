// package graph is designed to handle nodes
// and edges in simple graphs
//
// Reference: https://flaviocopes.com/golang-data-structure-graph/
package graph

import (
	"fmt"
	"testing"
)

var g ItemGraph

func TestTraverse(t *testing.T) {
	g.Traverse(func(n *Node) {
		fmt.Printf("%v\n", n)
	})
}

func fillGraph() {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}

func TestAdd(t *testing.T) {
	fillGraph()
	g.String()
}
