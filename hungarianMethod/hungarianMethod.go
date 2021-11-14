package hungarianmethod

type pair struct {
	row   uint
	col   uint
}

func (p *pair) Row() uint {
	return p.row
}

func (p *pair) Col() uint {
	return p.col
}

func (t *Table) Solve() *pair {
	t.Step1()
	if t.LinesNeededToCoverZeros() == uint(len(t.values)) {
		return nil
	}

	return nil
}