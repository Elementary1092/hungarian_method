package hungarian_method

func (t *Table) IsDone(linesToCover uint) bool {
	return uint(len(t.values)) == linesToCover
}

func (t *Table) Solve() []*Coord {
	t.Step1()

	minLinesNeededToCoverZeros, covered := t.CoverZeros()
	for !t.IsDone(minLinesNeededToCoverZeros) {
		t.Modify(covered)
		minLinesNeededToCoverZeros, covered = t.CoverZeros()
	}

	return t.ChooseOperators()
}