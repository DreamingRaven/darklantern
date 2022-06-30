package dager

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewMDGDager(t *testing.T) {
	NewMDGDager()
}

func TestNodeManipulation(t *testing.T) {
	dager := NewMDGDager()
	nA := Node{name: "A", item: "someitem"}
	nB := Node{name: "B", item: "someitem"}

	dager.AddNode(&nA)
	err := dager.AddNode(&nA)
	if err == nil {
		t.Fatal("Dager has failed to stop the insertion of the same node twice")
	}
	dager.RmNode(&nA)
	dager.AddNode(&nB)
	dager.RmNode(&nB)
}

func TestEdgeManipulation(t *testing.T) {
	g := NewMDGDager()
	nA := Node{name: "A", item: "in the graph"}
	nB := Node{name: "B", item: "in the graph"}
	nC := Node{name: "C", item: "in the graph"}
	nD := Node{name: "D", item: "not in the graph"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nB, &nC)

	// TODO add cycle detection
	g.AddEdge(&nC, &nA)

	err := g.AddEdge(&nC, &nD)
	if err == nil {
		t.Fatal("Dager has failed to prevent edge between an internal and external node")
	}

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

func TestList(t *testing.T) {

	g := NewMDGDager()

	nA := Node{name: "A", item: "someitem"}
	nB := Node{name: "B", item: "someitem"}
	nC := Node{name: "C", item: "someitem"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nB, &nC)

	o := g.List()

	if len(strings.Split(strings.TrimRight(o, "\n"), "\n")) != 3 {
		t.Fatal("Dager has not generated the right number of outputs for the number of nodes")
	}

	fmt.Print(o)
}
