package hungarian_method

func (t *Table) CoverZeros() (uint, []*Coord) {
	zerosCoordinates := t.coordinatesOfZeros()
	zerosArrangedByRows := arrangeByRows(len(t.values), zerosCoordinates)
	zerosArrangedByCols := arrangeByCols(len(t.values), zerosCoordinates)
	var linesCount uint
	var coveredRowsAndCols []*Coord

	for _, zerosInRow := range zerosArrangedByRows {
		if len(zerosInRow) != 0 {
			maxZerosInCol := maxAmountOfZerosInCol(zerosArrangedByCols)
			if len(zerosInRow) < maxZerosInCol && len(zerosInRow) < 3 {
				linesCount += uint(len(zerosInRow))

				for _, coord := range zerosInRow {
					zerosArrangedByRows, zerosArrangedByCols = deleteCol(
						zerosArrangedByRows,
						zerosArrangedByCols,
						coord,
					)

					coveredRowsAndCols = append(
						coveredRowsAndCols, 
						&Coord{
							row: -1,
							col: coord.col,
						},
					)
				}
			} else {
				rowNum := zerosInRow[0].row

				for _, coord := range zerosInRow {
					zerosArrangedByCols = deleteRow(
						zerosArrangedByCols,
						coord,
					)
				}

				coveredRowsAndCols = append(
					coveredRowsAndCols, 
					&Coord{
						row: rowNum,
						col: -1,
					},
				)

				linesCount++
			}
		}
	}

	return linesCount, coveredRowsAndCols
}

func arrangeByRows(numberOfRows int, zerosCoordinates []*Coord) [][]*Coord {
	rowsByArrangment := make([][]*Coord, numberOfRows)

	for _, coords := range zerosCoordinates {
		rowsByArrangment[coords.row] = append(rowsByArrangment[coords.row], coords)
	}

	return rowsByArrangment
}

func arrangeByCols(numberOfCols int, zerosCoordinates []*Coord) [][]*Coord {
	colsByArrangment := make([][]*Coord, numberOfCols)

	for _, coords := range zerosCoordinates {
		colsByArrangment[coords.col] = append(colsByArrangment[coords.col], coords)
	}

	return colsByArrangment
}

func maxAmountOfZerosInCol(zerosInCols [][]*Coord) int {
	var max int

	for _, zerosInCol := range zerosInCols {
		if len(zerosInCol) > max {
			max = len(zerosInCol)
		}
	}

	return max
}

func deleteCol(zerosInRow, zerosInCol [][]*Coord, p *Coord) ([][]*Coord, [][]*Coord) {
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

func deleteRow(zerosInCol [][]*Coord, p *Coord) [][]*Coord {
	for i := range zerosInCol {
		for j := 0; j < len(zerosInCol[i]); j++ {
			if zerosInCol[i][j].row == p.row {
				zerosInCol[i] = deleteFromArray(zerosInCol[i], j)
				j--
			}
		}
	}

	return zerosInCol
}