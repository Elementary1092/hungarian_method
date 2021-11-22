package hungarian_method

import "fmt"

type Coord struct {
	row int
	col int
}

func (p *Coord) Row() int {
	return p.row
}

func (p *Coord) Col() int {
	return p.col
}

func (p *Coord) String() string {
	return fmt.Sprintf("row: %d; col: %d |", p.row, p.col)
}
