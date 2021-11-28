package linked_list

import (
	c "coord"
)

type Node struct {
	Coord *c.Coord
	next  *Node 
}

func NewNode(coord *c.Coord) *Node {
	return &Node{
		Coord: coord,
		next:  nil,
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

func (n *Node) HasNext() bool {
	return n.next != nil
}