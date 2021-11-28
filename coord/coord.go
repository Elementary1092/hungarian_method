package coord

import "fmt"

type Coord struct {
	row int
	col int
}

func NewCoord(row int, col int) *Coord {
	return &Coord{
		row: row,
		col: col,
	}
}

func (p *Coord) Row() int {
	return p.row
}

func (p *Coord) Col() int {
	return p.col
}

func (p *Coord) String() string {
	return fmt.Sprintf("| row: %d; col: %d |", p.row, p.col)
}
