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
func (g *mdg) AddNode(n *Node) error {
	exists := g.IsNode(n)
	defer g.lock.Unlock()
	g.lock.Lock()
	var out error
	if exists {
		// out := errors.New(fmt.Sprintf("node: (%v, <%p>) already exists", n.name, n))
		out = errors.New("Node already exists")
	} else {
		g.nodes = append(g.nodes, n)
	}
	// g.lock.Unlock()
	return out
}

// AddEdge adds a single directed edge between two nodes
func (g *mdg) AddEdge(from, to *Node) error {
	g.lock.Lock()
	g.edges[from] = append(g.edges[from], to)
	g.lock.Unlock()
	return nil
}

// RmNode removes a node and its associated edges
func (g *mdg) RmNode(n *Node) {}

// RmEdge removed a directed edge between two nodes
func (g *mdg) RmEdge(from, to *Node) {}

func (g *mdg) List() string {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += fmt.Sprintf("node: (%v, <%p>), f-edges: [", g.nodes[i].name, g.nodes[i])
		for j := 0; j < len(g.edges[g.nodes[i]]); j++ {
			s += fmt.Sprintf("(%v, <%p>)", g.edges[g.nodes[i]][j], g.edges[g.nodes[i]][j])
		}
		s += "]\n"
	}
	g.lock.RUnlock()
	return s
}

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
