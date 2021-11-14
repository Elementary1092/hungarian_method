package hungarianmethod

func (t *Table) LinesNeededToCoverZeros() uint {
	var uncoveredZerosInRows = make([]int32, len(t.values))
	var uncoveredZerosInCols = make([]int32, len(t.values))
	var linesToCover uint

	for i := range t.values {
		for j := range  t.values[i] {
			if t.values[i][j] == 0 {
				uncoveredZerosInRows[i]++
				uncoveredZerosInCols[j]++
			}
		}
	}

	return linesToCover
}