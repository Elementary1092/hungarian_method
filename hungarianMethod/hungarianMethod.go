package hungarianmethod

import "fmt"

type pair struct {
	row int
	col int
}

func (p *pair) Row() int {
	return p.row
}

func (p *pair) Col() int {
	return p.col
}

func (p *pair) String() string {
	return fmt.Sprintf("row: %d; col: %d |", p.row, p.col)
}

func (t *Table) IsDone(linesToCover uint) bool {
	return uint(len(t.values)) == linesToCover
}

func (t *Table) Solve() []*pair {
	t.Step1()

	minLinesNeededToCoverZeros, covered := t.CoverZeros()
	for !t.IsDone(minLinesNeededToCoverZeros) {
		t.Modify(covered)
		minLinesNeededToCoverZeros, covered = t.CoverZeros()
	}

	return t.ChooseOperators()
}