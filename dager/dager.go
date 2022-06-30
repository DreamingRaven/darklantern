// Multi-Directed Acyclic Graph interfacer
package dager

import (
	"fmt"
	"sync"
)

type Node struct {
	name string
	item any
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.name)
}

// nulti-Directed Acyclic-enforcing Graph intERface
type Dager interface {
	AddNode(n *Node)
	AddEdge(from, to *Node)
	RmNode(n *Node)
	RmEdge(from, to *Node)
}

// Multi-Directed Graph but not itself acyclic
type mdg struct {
	nodes []*Node
	edges map[*Node][]*Node
	lock  sync.RWMutex
}

// AddNode adds a given node to the graph as is
func (g *mdg) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds a single directed edge between two nodes
func (g *mdg) AddEdge(from, to *Node) {
	g.lock.Lock()
	g.edges[from] = append(g.edges[from], to)
	g.lock.Unlock()
}

// RmNode removes a node and its associated edges
func (g *mdg) RmNode(n *Node) {}

// RmEdge removed a directed edge between two nodes
func (g *mdg) RmEdge(from, to *Node) {}

// NewMDGDager initialises a multi-directed graph (MDG) struct and wraps it in the Dager handling interface
func NewMDGDager() Dager {
	return &mdg{edges: make(map[*Node][]*Node)}
}
