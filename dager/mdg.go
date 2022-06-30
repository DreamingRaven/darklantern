package dager

import (
	"errors"
	"fmt"
	"sync"
)

// Multi-Directed Graph but not itself acyclic
type mdg struct {
	nodes []*Node
	edges map[*Node][]*Node
	lock  sync.RWMutex
}

// AddNode adds a given node to the graph as is
// returns error if node is already present by address
func (g *mdg) AddNode(n *Node) error {
	exists := g.IsNode(n) // exists has its own lock
	if exists {
		return errors.New(fmt.Sprintf("node: (%v, <%p>) already exists", n.name, n))
	}
	defer g.lock.Unlock()
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	return nil
}

// AddEdge adds a single directed edge between two nodes
// returns error if either node does not already exist
func (g *mdg) AddEdge(from, to *Node) error {
	if !g.IsNode(from) || !g.IsNode(to) {
		return errors.New(fmt.Sprintf("one of (%v, <%p>) (%v, <%p>) does not exist",
			from.name, from, to.name, to))
	}

	defer g.lock.Unlock()
	g.lock.Lock()
	g.edges[from] = append(g.edges[from], to)
	return nil
}

// RmNode removes a node and its associated edges
func (g *mdg) RmNode(n *Node) error { return nil }

// RmEdge removed a directed edge between two nodes
func (g *mdg) RmEdge(from, to *Node) error { return nil }

// List all nodes and their forward edges line-by-line
func (g *mdg) List() string {
	defer g.lock.RUnlock()
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += fmt.Sprintf("node: (%v, <%p>), f-edges: [", g.nodes[i].name, g.nodes[i])
		for j := 0; j < len(g.edges[g.nodes[i]]); j++ {
			s += fmt.Sprintf("(%v, <%p>)", g.edges[g.nodes[i]][j], g.edges[g.nodes[i]][j])
		}
		s += "]\n"
	}
	return s
}

// IsNode compare node to list of nodes to confirm if it is or is not already being tracked
func (g *mdg) IsNode(n *Node) bool {
	defer g.lock.RUnlock()
	g.lock.RLock()
	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i] == n {
			return true
		}
	}
	return false
}

// NewMDGDager initialises a multi-directed graph (MDG) struct and wraps it in the Dager handling interface
func NewMDGDager() Dager {
	return &mdg{edges: make(map[*Node][]*Node)}
}
