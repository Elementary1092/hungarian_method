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
		if val.weight < min {
			min = val.weight
		}
	}

	return min
}

func (t *Table) findMinInColumn(col uint) int64 {
	min := int64(math.MaxInt64)

	for _, row := range t.values {
		if row[col].weight < min {
			min = row[col].weight
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
					int(t.values[i][j].value),
				),
			)

			buffer.WriteRune(' ')
		}
		
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

func (t *Table) coordinatesOfZeros() []*Coord {
	coordinatesOfZeros := make([]*Coord, 0)

	for i := range t.values {
		for j := range t.values[i] {
			if t.values[i][j].weight == 0 {
				coordinatesOfZeros = append(
					coordinatesOfZeros,
					&Coord{
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
			if t.values[i][j].weight < min {
				min = t.values[i][j].weight
			}
		}
	}

	return min
}

func deleteFromArray(arr []*Coord, toBeDeleted int) []*Coord {
	arr = append(arr[:toBeDeleted], arr[(toBeDeleted+1):]...)

	return arr
}

func deleteFrom2DArray(arr [][]*Coord, toBeDeleted int) [][]*Coord {
	arr = append(arr[:toBeDeleted], arr[(toBeDeleted + 1):]...)

	return arr
}