package hungarian_method

import "fmt"

type Pair struct {
	row int
	col int
}

func (p *Pair) Row() int {
	return p.row
}

func (p *Pair) Col() int {
	return p.col
}

func (p *Pair) String() string {
	return fmt.Sprintf("row: %d; col: %d |", p.row, p.col)
}

func (t *Table) IsDone(linesToCover uint) bool {
	return uint(len(t.values)) == linesToCover
}

func (t *Table) Solve() []*Pair {
	t.Step1()

	minLinesNeededToCoverZeros, covered := t.CoverZeros()
	for !t.IsDone(minLinesNeededToCoverZeros) {
		t.Modify(covered)
		minLinesNeededToCoverZeros, covered = t.CoverZeros()
	}

	return t.ChooseOperators()
}