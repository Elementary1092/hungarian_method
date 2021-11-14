package list

type List struct {
	head *Node
}

func NewList() *List {
	return &List{
		head: nil,
	}
}

func (l *List) PushBack(n *Node) {
	curr := l.head

	if curr == nil {
		l.head = n
		return
	}

	for curr.Next() != nil {
		curr = curr.next
	}

	curr.next = n
}

func (l *List) PushFront(n *Node) {
	prevHead := l.head
	l.head = n
	n.next = prevHead
}

func (l *List) DeleteHead() {
	prevHead := l.head
	l.head = l.head.next
	prevHead.next = nil
}

func (l *List) FindZeroValuedNodes() []*Node {
	zeroValuedNodes := make([]*Node, 0)

	curr := l.head
	for curr != nil {
		if curr.weight == 0 {
			zeroValuedNodes = append(zeroValuedNodes, curr)
		}

		curr = curr.next
	}

	return zeroValuedNodes
}

func (l *List) GetNodeWithIndex(index uint32) *Node {
	curr := l.head

	for index > 0 {
		curr = curr.next
		index--
	}

	return curr
}

func (l *List) GetIndex(n *Node) int64 {
	curr := l.head

	for i := int64(0); curr != nil; i++ {
		if curr == n {
			return i
		}
	}

	return -1
}

func (l *List) TraverseListNodesWithAction(fn func(*int32)) {	
	curr := l.head
	for curr != nil {
		fn(&curr.weight)

		curr = curr.next
	}

}

func (l *List) FindMin() (min int32) {
	curr := l.head

	if (curr != nil) {
		min = curr.weight
		curr = curr.next
	}

	for curr != nil {
		if curr.value < min {
			min = curr.value
		}
	}

	return
}