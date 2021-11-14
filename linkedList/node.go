package list

type Node struct {
	value  int32
	weight int32
	next   *Node
}

//New node is created with nullified next value
func NewNode(value int32) *Node {
	return &Node{
		value: value,
		weight: value,
	}
} 

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Value() int32 {
	return n.value
}

func (n *Node) Weight() int32 {
	return n.weight
}

func (n *Node) ChangeValue(newValue int32) {
	n.value = newValue
}

func (n *Node) ChangeNext(next *Node) {
	n.next = next
}

func (n *Node) ChangeWeight(newWeight int32) {
	n.weight = newWeight
}