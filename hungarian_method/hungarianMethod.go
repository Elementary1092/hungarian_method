package hungarian_method

import (
	c "coord"
)

func (t *Table) IsDone(linesToCover uint) bool {
	return uint(len(t.values)) == linesToCover
}

func (t *Table) Solve() []*c.Coord {
	t.Step1()

	minLinesNeededToCoverZeros, covered := t.CoverZeros()
	for !t.IsDone(minLinesNeededToCoverZeros) {
		t.Modify(covered)
		minLinesNeededToCoverZeros, covered = t.CoverZeros()
	}

	return t.ChooseOperators()
}