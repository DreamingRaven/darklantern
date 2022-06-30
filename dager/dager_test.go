package dager

import (
	"testing"
)

func TestNewMDGDager(t *testing.T) {
	NewMDGDager()
}

func TestNodeManipulation(t *testing.T) {
	dager := NewMDGDager()
	nA := Node{name: "A", item: "someitem"}
	nB := Node{name: "B", item: "someitem"}

	// TODO add checking that node does not already exist in graph
	dager.AddNode(&nA)
	dager.AddNode(&nA)
	dager.RmNode(&nA)
	dager.AddNode(&nB)
	dager.RmNode(&nB)
}

func TestEdgeManipulation(t *testing.T) {
	g := NewMDGDager()
	nA := Node{name: "A", item: "someitem"}
	nB := Node{name: "B", item: "someitem"}
	nC := Node{name: "C", item: "someitem"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nB, &nC)
	// TODO add cycle detection
	g.AddEdge(&nC, &nA)

	g.RmEdge(&nA, &nB)
	g.RmEdge(&nB, &nC)
	g.RmEdge(&nA, &nB)
	g.RmEdge(&nA, &nB)
}

func TestAdjacentManipulation(t *testing.T) {

	g := NewMDGDager()

	nA := Node{name: "A", item: "someitem"}
	nB := Node{name: "B", item: "someitem"}
	nC := Node{name: "C", item: "someitem"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nB, &nC)

	g.RmNode(&nB)
	// TODO add manipulation of neighbors when connected node is removed
}
