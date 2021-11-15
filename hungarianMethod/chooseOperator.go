package hungarianmethod

func (t *Table) ChooseOperators() []*pair {
	coordsOfZeros := t.coordinatesOfZeros()
	rowsWithZeros := make([][]*pair, len(t.values))
	
	for _, coord := range coordsOfZeros {
		rowsWithZeros[coord.row] = append(rowsWithZeros[coord.row], coord)
	}

	res := []*pair{}
	for len(rowsWithZeros) != 0 {
		for i := 0; i < len(rowsWithZeros); i++ {
			if len(rowsWithZeros[i]) == 1 {
				res = append(res, rowsWithZeros[i]...)
				rowsWithZeros = deleteCols(rowsWithZeros, rowsWithZeros[i][0].col)
				rowsWithZeros = append(rowsWithZeros[:i], rowsWithZeros[(i+1):]...)
				i--
			}
		}
	}

	return res
}

func deleteCols(zeros [][]*pair, colNum int) [][]*pair {
	for i := range zeros {
		for j := 0; j < len(zeros[i]); j++ {
			if zeros[i][j].col == colNum {
				zeros[i] = append(zeros[i][:j], zeros[i][(j+1):]...)
				j--
			}
		}
	}

	return zeros
}