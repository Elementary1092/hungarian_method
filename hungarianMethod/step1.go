package hungarianmethod

func (t *Table) Step1() {
	t.reduceRows()
	t.reduceColumns()
}

func (t *Table) reduceRows() {
	for i := range t.values {
		minVal := t.findMinInRow(uint(i))

		for j := range t.values[i] {
			t.values[i][j] -= minVal
		} 
	}
}

func (t *Table) reduceColumns() {
	for j := 0; j < len(t.values); j++ {
		minVal := t.findMinInColumn(uint(j))

		for i := 0; i < len(t.values); i++ {
			t.values[i][j] -= minVal
		}
	}
}