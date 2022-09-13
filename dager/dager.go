// Multi-Directed Acyclic Graph interfacer
// abstracts graphs and their nodes' data in a nice format for neural networks
package dager

import (
	"gitlab.com/deepcypher/darklantern/node"
)

// nulti-Directed Acyclic-enforcing Graph intERface
type Dager interface {
	List() string
	IsNode(n *node.Node) bool
	AddNode(n *node.Node) error
	AddEdge(from, to *node.Node) error
	RmNode(n *node.Node) error
	RmEdge(from, to *node.Node) error
}
