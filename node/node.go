package node

import "fmt"

type Node struct {
	Name string
	Item any
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Name)
}
