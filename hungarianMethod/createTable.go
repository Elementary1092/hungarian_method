package hungarianmethod;

type Table struct {
	values            [][]int64
	shouldConsider    uint
	shouldBeTrasposed bool
}

func NewTable(vals [][]int64) *Table {
	if len(vals) != len(vals[0]) {
		return makeUpSqureTable(vals)
	}

	var table = &Table{
		values:         make([][]int64, len(vals)),
		shouldConsider: uint(len(vals)),
	}

	copy(table.values, vals)
	
	return table
}

func makeUpSqureTable(vals [][]int64) *Table {
	if len(vals) < len(vals[0]) {
		table := &Table{
			values:            make([][]int64, len(vals[0])),
			shouldConsider:    uint(len(vals)),
			shouldBeTrasposed: true,
		}

		rowsCopied := copy(table.values, vals)

		for i := rowsCopied; i < len(table.values); i++ {
			table.values[i] = make([]int64, len(table.values))
		}

		return table
	}
	
	table := &Table{
		values:         make([][]int64, len(vals)),
		shouldConsider: uint(len(vals[0])),
	}

	for i := 0; i < len(vals); i++ {
		table.values[i] = make([]int64, len(vals))
		copy(table.values[i], vals[i])
	}

	return table
}