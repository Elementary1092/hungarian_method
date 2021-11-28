package linked_list

type List struct {
	head *Node
	len  int
}

func NewList() *List {
	return &List{
		head: &Node{},
		len: 0,
	}
}

func (l *List) Add(n *Node) {
	if l.head.next == nil {
		l.head.next = n
		l.len++
		return
	}

	cur := l.head.next
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = n
	l.len++
}

func (l *List) Len() int {
	return l.len
} 

func (l *List) GetFirstNode() *Node {
	return l.head.next
}

func (l *List) DeleteEntriesWithCol(col int) {
	cur := l.head
	for cur.next != nil {
		if cur.next.Coord.Col() == col {
			cur.next = cur.next.next
			l.len--
		} else {
			cur = cur.next
		}
	}
}

func (l *List) DeleteEntriesWithRow(row int) {
	cur := l.head
	for cur.next != nil {
		if cur.next.Coord.Row() == row {
			cur.next = cur.next.next
			l.len--
		} else {
			cur = cur.next
		}
	}
}