package hungarian_method

type Table struct {
	values            [][]*Pair
}

func NewTable(vals [][]int64) *Table {
	/*if len(vals) != len(vals[0]) {
		return makeUpSqureTable(vals)
	}*/

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
/*
func makeUpSqureTable(vals [][]int64) *Table {
	if len(vals) < len(vals[0]) {
		table := &Table{
			values:            make([][]int64, len(vals[0])),
		}

		rowsCopied := copy(table.values, vals)

		for i := rowsCopied; i < len(table.values); i++ {
			table.values[i] = make([]int64, len(table.values))
		}

		return table
	}
	
	table := &Table{
		values:         make([][]int64, len(vals)),
	}

	for i := 0; i < len(vals); i++ {
		table.values[i] = make([]int64, len(vals))
		copy(table.values[i], vals[i])
	}

	return table
}*/