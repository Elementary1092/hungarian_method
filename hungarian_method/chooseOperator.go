package hungarian_method

import (
	c "coord"
)

func (t *Table) ChooseOperators() []*c.Coord {
	coordsOfZeros := t.coordinatesOfZeros()
	rowsWithZeros := make([][]*c.Coord, len(t.values))
	
	cur := coordsOfZeros.GetFirstNode()
	for cur != nil {
		rowsWithZeros[cur.Coord.Row()] = append(rowsWithZeros[cur.Coord.Row()], cur.Coord)
		cur = cur.Next()
	}
	
	return getPossibleResults(rowsWithZeros, len(t.values))
}

func getPossibleResults(zeros [][]*c.Coord, shouldConsider int) []*c.Coord {
	if len(zeros) == 0 {
		return nil
	}

	if len(zeros) == 1 {
		return zeros[0]
	}

	res := make([]*c.Coord, 0)
	for j := range zeros[0] {
		temp := getPossibleResults(
			returnElemsExceptFirstRowAndCol(zeros, zeros[0][j].Col()),
			shouldConsider - 1,
		)

		if len(temp) == shouldConsider - 1 {
			res = append(temp, zeros[0][j])
			break
		}
	}

	return res
}

func returnElemsExceptFirstRowAndCol(zeros [][]*c.Coord, colNum int) [][]*c.Coord {
	res := make([][]*c.Coord, 0)
	
	for i := 1; i < len(zeros); i++ {
		temp := make([]*c.Coord, 0)

		for j := range zeros[i] {
			if zeros[i][j].Col() != colNum {
				temp = append(temp, zeros[i][j])
			}
		}

		if len(temp) != 0 {
			res = append(res, temp)
		}
	}

	return res
}