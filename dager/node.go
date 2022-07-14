package dager

import "fmt"

type Node struct {
	name string
	item any
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.name)
}
