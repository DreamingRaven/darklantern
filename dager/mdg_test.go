package dager

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/deepcypher/darklantern/node"
)

func TestRemoveLastNode(t *testing.T) {
	g := NewMDGDager()
	nA := node.Node{Name: "A", Item: "someitem"}
	nB := node.Node{Name: "B", Item: "someitem"}
	err := g.AddNode(&nA)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddNode(&nB)
	if err != nil {
		t.Fatal(err)
	}
	err = g.RmNode(&nB)
	if err != nil {
		t.Fatal(err)
	}
	err = g.RmNode(&nA)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewMDGDager(t *testing.T) {
	NewMDGDager()
}

func TestNodeManipulation(t *testing.T) {
	g := NewMDGDager()
	nA := node.Node{Name: "A", Item: "someitem"}
	nB := node.Node{Name: "B", Item: "someitem"}

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
	nA := node.Node{Name: "A", Item: "in the graph"}
	nB := node.Node{Name: "B", Item: "in the graph"}
	nC := node.Node{Name: "C", Item: "in the graph"}
	nD := node.Node{Name: "D", Item: "not in the graph"}

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

	nA := node.Node{Name: "A", Item: "someitem"}
	nB := node.Node{Name: "B", Item: "someitem"}
	nC := node.Node{Name: "C", Item: "someitem"}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nB, &nC)

	before := g.List()
	err := g.RmNode(&nB)
	if err != nil {
		t.Fatal(err)
	}
	if g.IsNode(&nA) == false {
		fmt.Printf("before:\n%s\n", before)
		fmt.Printf("after:\n%s\n", g.List())
		t.Fatal("Dager has murdered a wrong node A")
	}
	if g.IsNode(&nB) == true {
		fmt.Println(before)
		fmt.Println(g.List())
		t.Fatal("Dager has failed to eliminate the target B")
	}
	if g.IsNode(&nC) == false {
		fmt.Println(before)
		fmt.Println(g.List())
		t.Fatal("Dager has murdered a wrong node C")
	}

	err = g.RmNode(&nB)
	if err == nil {
		t.Fatal("Dager failed to catch removal of non-existant item")
	}
}

func TestList(t *testing.T) {

	g := NewMDGDager()

	nA := node.Node{Name: "A", Item: "someitem"}
	nB := node.Node{Name: "B", Item: "someitem"}
	nC := node.Node{Name: "C", Item: "someitem"}

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
