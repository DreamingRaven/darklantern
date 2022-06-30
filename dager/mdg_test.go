package dager

import (
	"strings"
	"testing"
)

func TestNewMDGDager(t *testing.T) {
	NewMDGDager()
}

func TestNodeManipulation(t *testing.T) {
	g := NewMDGDager()
	nA := Node{name: "A", item: "someitem"}
	nB := Node{name: "B", item: "someitem"}

	// adding the same node twice
	g.AddNode(&nA)
	err := g.AddNode(&nA)
	if err == nil {
		t.Fatal("Dager has failed to stop the insertion of the same node twice")
	}

	// removing what should be the only node
	err = g.RmNode(&nA)
	if err != nil {
		t.Fatal(err)
	}

	// adding a different node
	err = g.AddNode(&nB)
	if err != nil {
		t.Fatal(err)
	}

	// removing the different node twice
	g.RmNode(&nB)
	err = g.RmNode(&nB)
	if err == nil {
		t.Fatal("Dager allowed the removal of a node that was not present")
	}

	// re-adding what should be already removed nodes
	err = g.AddNode(&nA)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddNode(&nB)
	if err != nil {
		t.Fatal(err)
	}
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

	// add some basic edges
	err := g.AddEdge(&nA, &nB)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(&nB, &nC)
	if err != nil {
		t.Fatal(err)
	}

	// TODO add cycle detection
	err = g.AddEdge(&nC, &nA)
	if err != nil {
		t.Fatal(err)
	}

	err = g.AddEdge(&nC, &nD)
	if err == nil {
		t.Fatal("Dager has failed to prevent edge between an internal and external node")
	}

	// removing all edges then trying to remove more
	err = g.RmEdge(&nA, &nB)
	if err != nil {
		t.Fatal(err)
	}
	err = g.RmEdge(&nB, &nC)
	if err != nil {
		t.Fatal(err)
	}
	err = g.RmEdge(&nA, &nB)
	if err == nil {
		t.Fatal("Dager failed to error when removing a non existant edge")
	}

	// fmt.Println(g.List())
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

	// TODO add manipulation of neighbors when connected node is removed
	g.RmNode(&nB)
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

	// fmt.Print(o)
	// fmt.Println(g.List())
}
