package hungarian_method

import (
	"sync"
	c "coord"
)

func (t *Table) Modify(covered []*c.Coord) {
	rowsLeft, colsLeft := t.identifyUncoveredPart(covered)

	minVal := t.findMinInUncoveredPart(rowsLeft, colsLeft)

	for _, i := range rowsLeft {
		for _, j := range colsLeft {
			t.values[i][j].weight -= minVal
		}
	}

	coveredRows := make([]int, 0)
	coveredCols := make([]int, 0)

	for _, p := range covered {
		if p.Row() != -1 {
			coveredRows = append(coveredRows, p.Row())
		} else {
			coveredCols = append(coveredCols, p.Col())
		}
	}

	for _, row := range coveredRows {
		for _, col := range coveredCols {
			t.values[row][col].weight += minVal
		}
	}
}

func (t *Table) identifyUncoveredPart(covered []*c.Coord) ([]int, []int) {
	rowsLeft := make([]int, len(t.values))
	colsLeft := make([]int, len(t.values))

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		numerate(rowsLeft)
	} ()

	go func() {
		defer wg.Done()
		numerate(colsLeft)
	} ()

	wg.Wait()

	for _, p := range covered {
		if p.Row() != -1 {
			rowsLeft[p.Row()] = 0
		} else {
			colsLeft[p.Col()] = 0
		}
	}

	rowsLeft = deleteZeros(rowsLeft)
	colsLeft = deleteZeros(colsLeft)

	return rowsLeft, colsLeft
}

func numerate(arr []int) {
	for i := range arr {
		arr[i] = i + 1
	}
}

func deleteZeros(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr = append(arr[:i], arr[(i+1):]...)
			i--
		} else {
			arr[i]--
		}
	}

	return arr
}