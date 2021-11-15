package hungarianmethod

func (t *Table) LinesNeededToCoverZeros() uint {
	zerosCoordinates := t.coordinatesOfZeros()
	zerosArrangedByRows := arrangeByRows(len(t.values), zerosCoordinates)
	zerosArrangedByCols := arrangeByCols(len(t.values), zerosCoordinates)
	var linesCount uint

	for _, zerosInRow := range zerosArrangedByRows {
		if len(zerosInRow) != 0 {
			maxZerosInCol := maxAmountOfZerosInCol(zerosArrangedByCols)
			if len(zerosInRow) < maxZerosInCol {
				linesCount += uint(len(zerosInRow))

				for _, coord := range zerosInRow {
					zerosArrangedByRows, zerosArrangedByCols = deleteCol(
						zerosArrangedByRows,
						zerosArrangedByCols,
						coord,
					)
				}
			} else {
				for _, coord := range zerosInRow {
					zerosArrangedByCols = deleteRow(
						zerosArrangedByCols,
						coord,
					)
				}
				
				linesCount++
			}
		}
	}

	return linesCount
}

func (t *Table) coordinatesOfZeros() []*pair {
	coordinatesOfZeros := make([]*pair, 0)

	for i := range t.values {
		for j := range t.values[i] {
			if t.values[i][j] == 0 {
				coordinatesOfZeros = append(
					coordinatesOfZeros,
					&pair{
						row: uint(i),
						col: uint(j),
					},
				)
			}
		}
	}

	return coordinatesOfZeros
}

func arrangeByRows(numberOfRows int, zerosCoordinates []*pair) [][]*pair {
	rowsByArrangment := make([][]*pair, numberOfRows)

	for _, coords := range zerosCoordinates {
		rowsByArrangment[coords.row] = append(rowsByArrangment[coords.row], coords)
	}

	return rowsByArrangment
}

func arrangeByCols(numberOfCols int, zerosCoordinates []*pair) [][]*pair {
	colsByArrangment := make([][]*pair, numberOfCols)

	for _, coords := range zerosCoordinates {
		colsByArrangment[coords.col] = append(colsByArrangment[coords.col], coords)
	}

	return colsByArrangment
}

func maxAmountOfZerosInCol(zerosInCols [][]*pair) int {
	var max int

	for _, zerosInCol := range zerosInCols {
		if len(zerosInCol) > max {
			max = len(zerosInCol)
		}
	}

	return max
}

func deleteCol(zerosInRow, zerosInCol [][]*pair, p *pair) ([][]*pair, [][]*pair) {
	for i := 0; i < len(zerosInCol); i++ {
		if len(zerosInCol[i]) == 0 {
			continue
		}

		if zerosInCol[i][0].col == p.col {
			zerosInCol = deleteFrom2DArray(zerosInCol, i)

			break
		}
	}

	for i := range zerosInRow {
		for j := 0; j < len(zerosInRow[i]); j++ {
			if zerosInRow[i][j].col == p.col {
				zerosInRow[i] = deleteFromArray(zerosInRow[i], j)
				j--
			}
		}
	}

	return zerosInRow, zerosInCol
}

func deleteRow(zerosInCol [][]*pair, p *pair) [][]*pair {
	for i := range zerosInCol {
		for j := 0; i < len(zerosInCol[i]); j++ {
			if zerosInCol[i][j].row == p.row {
				zerosInCol[i] = deleteFromArray(zerosInCol[i], j)
				j--
			}
		}
	}

	return zerosInCol
}