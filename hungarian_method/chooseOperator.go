package hungarian_method

func (t *Table) ChooseOperators() []*pair {
	coordsOfZeros := t.coordinatesOfZeros()
	rowsWithZeros := make([][]*pair, len(t.values))
	
	for _, coord := range coordsOfZeros {
		rowsWithZeros[coord.row] = append(rowsWithZeros[coord.row], coord)
	}
	
	return getPossibleResults(rowsWithZeros, len(t.values))
}

func getPossibleResults(zeros [][]*pair, shouldConsider int) []*pair {
	if len(zeros) == 0 {
		return nil
	}

	if len(zeros) == 1 {
		return zeros[0]
	}

	res := make([]*pair, 0)
	for j := range zeros[0] {
		temp := getPossibleResults(
			returnElemsExceptFirstRowAndCol(zeros, zeros[0][j].col),
			shouldConsider - 1,
		)

		if len(temp) == shouldConsider - 1 {
			res = append(temp, zeros[0][j])
			break
		}
	}

	return res
}

func returnElemsExceptFirstRowAndCol(zeros [][]*pair, colNum int) [][]*pair {
	res := make([][]*pair, 0)
	
	for i := 1; i < len(zeros); i++ {
		temp := make([]*pair, 0)

		for j := range zeros[i] {
			if zeros[i][j].col != colNum {
				temp = append(temp, zeros[i][j])
			}
		}

		if len(temp) != 0 {
			res = append(res, temp)
		}
	}

	return res
}
/*
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
}*/