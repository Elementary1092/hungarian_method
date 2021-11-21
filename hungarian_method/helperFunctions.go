package hungarian_method

import (
	"math"
	"strconv"
	"strings"
)

func (t *Table) findMinInRow(row uint) int64 {
	searchIn := t.values[row]
	min := int64(math.MaxInt64)

	for _, val := range searchIn {
		if val < min {
			min = val
		}
	}

	return min
}

func (t *Table) findMinInColumn(col uint) int64 {
	min := int64(math.MaxInt64)

	for _, row := range t.values {
		if row[col] < min {
			min = row[col]
		}
	}

	return min
}

func (t *Table) String() string {
	var buffer strings.Builder

	for i := range t.values {
		for j := range t.values[i] {
			buffer.WriteString(
				strconv.Itoa(
					int(t.values[i][j]),
				),
			)

			buffer.WriteRune(' ')
		}
		
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

func (t *Table) coordinatesOfZeros() []*Pair {
	coordinatesOfZeros := make([]*Pair, 0)

	for i := range t.values {
		for j := range t.values[i] {
			if t.values[i][j] == 0 {
				coordinatesOfZeros = append(
					coordinatesOfZeros,
					&Pair{
						row: i,
						col: j,
					},
				)
			}
		}
	}

	return coordinatesOfZeros
}

func (t *Table) findMinInUncoveredPart(rowsLeft []int, colsLeft []int) int64 {
	var min int64 = math.MaxInt64

	for _, i := range rowsLeft {
		for _, j := range colsLeft {
			if t.values[i][j] < min {
				min = t.values[i][j]
			}
		}
	}

	return min
}

func deleteFromArray(arr []*Pair, toBeDeleted int) []*Pair {
	arr = append(arr[:toBeDeleted], arr[(toBeDeleted+1):]...)

	return arr
}

func deleteFrom2DArray(arr [][]*Pair, toBeDeleted int) [][]*Pair {
	arr = append(arr[:toBeDeleted], arr[(toBeDeleted + 1):]...)

	return arr
}