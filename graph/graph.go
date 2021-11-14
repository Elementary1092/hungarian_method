package graph

import (
	"errors"

	list "linkedList"
)

type Graph struct {
	values []*list.List
}

func NewGraph() *Graph {
	return &Graph{
		values: make([]*list.List, 0, 4),
	}
}

func (g *Graph) RowsLength() int {
	return len(g.values)
}

func (g *Graph) AddList(l *list.List) {
	if (len(g.values) == cap(g.values)) {
		g.values = append(g.values, l)
		return
	}

	g.values[len(g.values)] = l
}

func (g *Graph) TraverseAllGraphRowsWithAction(fn func(*int32)) {
	for _, l := range g.values {
		l.TraverseListNodesWithAction(fn)
	}
}

func (g *Graph) GetMinInColumn(column uint32) (int32, error) {
	if column < uint32(len(g.values)) {
		var min = g.values[0].GetNodeWithIndex(column)

		for i := 1; i < len(g.values); i++ {
			value := g.values[i].GetNodeWithIndex(column)
			if min.Weight() > value.Weight() {
				min = value
			}
		}

		return min.Weight(), nil
	}

	return 0, errors.New("Out of bound")
}

func (g *Graph) TraverseWithActionOnColumn(column uint32, fn func(*int32)) {
	for _, l := range g.values {
		n := l.GetNodeWithIndex(column)
		val := n.Weight()
		fn(&val)
		n.ChangeWeight(val)
	}
}

func (g *Graph) GetNodesWithZeroValuesInRow(row uint32) []*list.Node {
	return g.values[row].FindZeroValuedNodes()
}

func (g *Graph) GetNodesWithZeroValuesInColumn(column uint32) []*list.Node {
	zeroValuedNodes := make([]*list.Node, len(g.values))

	for i, l := range g.values {
		n := l.GetNodeWithIndex(column)
		if n.Weight() == 0 {
			zeroValuedNodes[i] = n
		}
	}

	return zeroValuedNodes
}

func (g *Graph) GetNodesWithZeroValuesInGraph() [][]*list.Node {
	res := make([][]*list.Node, len(g.values))

	for i, l := range g.values {
		nodes := l.FindZeroValuedNodes()
		row := make([]*list.Node, len(g.values))

		for _, n := range nodes {
			row[(l.GetIndex(n))] = n
		}

		res[i] = row
	}

	return nil
}