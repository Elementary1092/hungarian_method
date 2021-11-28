package hungarian_method

import (
	c "coord"
	ll "linked_list"
)

func (t *Table) CoverZeros() (uint, []*c.Coord) {
	zerosCoords := t.coordinatesOfZeros()
	zerosArrangedByRows := arrangeByRows(len(t.values), zerosCoords)
	zerosArrangedByCols := arrangeByCols(len(t.values), zerosCoords)
	var linesCount uint
	var coveredRowsAndCols []*c.Coord

	for _, zerosInRow := range zerosArrangedByRows {
		if zerosInRow.Len() != 0 {
			maxZerosInCol := maxAmountOfZerosInCol(zerosArrangedByCols)
			if zerosInRow.Len() < maxZerosInCol && zerosInRow.Len() < 3 {
				linesCount += uint(zerosInRow.Len())

				cur := zerosInRow.GetFirstNode()
				for cur != nil {
					deleteCol(zerosArrangedByRows, zerosArrangedByCols, cur.Coord)

					coveredRowsAndCols = append(
						coveredRowsAndCols, 
						c.NewCoord(-1, cur.Coord.Col()),
					)
					cur = cur.Next()
				}
			} else {
				rowNum := zerosInRow.GetFirstNode().Coord.Row()

				cur := zerosInRow.GetFirstNode()
				for cur != nil {
					deleteRow(zerosArrangedByCols, cur.Coord)
					cur = cur.Next()
				}

				coveredRowsAndCols = append(
					coveredRowsAndCols, 
					c.NewCoord(rowNum, -1),
				)

				linesCount++
			}
		}
	}

	return linesCount, coveredRowsAndCols
}

func arrangeByRows(numberOfRows int, zerosCoords *ll.List) []*ll.List {
	rowsByArrangment := make([]*ll.List, numberOfRows)
	for i := 0; i < numberOfRows; i++ {
		rowsByArrangment[i] = ll.NewList()
	}

	cur := zerosCoords.GetFirstNode()
	for cur != nil {
		add := ll.NewNode(cur.Coord)
		rowsByArrangment[cur.Coord.Row()].Add(add)
		cur = cur.Next()
	}

	return rowsByArrangment
}

func arrangeByCols(numberOfCols int, zerosCoords *ll.List) []*ll.List {
	colsByArrangment := make([]*ll.List, numberOfCols)
	for i := 0; i < numberOfCols; i++ {
		colsByArrangment[i] = ll.NewList()
	}

	cur := zerosCoords.GetFirstNode()
	for cur != nil {
		add := ll.NewNode(cur.Coord)
		colsByArrangment[cur.Coord.Col()].Add(add)
		cur = cur.Next()
	}

	return colsByArrangment
}

func maxAmountOfZerosInCol(zerosInCols []*ll.List) int {
	var max int

	for _, zerosInCol := range zerosInCols {
		if zerosInCol.Len() > max {
			max = zerosInCol.Len()
		}
	}

	return max
}

func deleteCol(zerosInRow, zerosInCol []*ll.List, p *c.Coord) {
	for i := 0; i < len(zerosInCol); i++ {
		if zerosInCol[i].Len() == 0 {
			continue
		}

		zerosInCol[i].DeleteEntriesWithCol(p.Col())
	}

	for i := range zerosInRow {
		zerosInRow[i].DeleteEntriesWithCol(p.Col())
	}
}

func deleteRow(zerosInCol []*ll.List, p *c.Coord) {
	for i := range zerosInCol {
		zerosInCol[i].DeleteEntriesWithRow(p.Row())
	}
}