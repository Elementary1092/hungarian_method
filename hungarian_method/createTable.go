package hungarian_method

type Table struct {
	values            [][]*Pair
}

func NewTable(vals [][]int64) *Table {
	if len(vals) != len(vals[0]) {
		panic("Expected square matrix")
	}

	var table = &Table{
		values:         make([][]*Pair, len(vals)),
	}

	for i, row := range vals {
		copiedRow := make([]*Pair, len(vals))
		for j, val := range row {
			copiedRow[j] = NewPair(val, val)
		}
		
		table.values[i] = copiedRow
	}
	
	return table
}

func (t *Table) ValueAt(row, col int) int64 {
	return t.values[row][col].value
}