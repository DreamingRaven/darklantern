// Multi-Directed Acyclic Graph interfacer
package dager

// nulti-Directed Acyclic-enforcing Graph intERface
type Dager interface {
	List() string
	IsNode(n *Node) bool
	AddNode(n *Node) error
	AddEdge(from, to *Node) error
	RmNode(n *Node) error
	RmEdge(from, to *Node) error
}
