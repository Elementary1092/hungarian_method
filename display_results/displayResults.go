package display_results

import (
	"fmt"

	alg "hungarian_method"
	c   "coord"
)

func Display(t *alg.Table, coords []*c.Coord) {
	var sum int64
	for _, coord := range coords {
		sum += t.ValueAt(coord.Row(), coord.Col())
		fmt.Printf(
			"Job %d should be done by %d worker and it will cost %d\n",
			coord.Row() + 1,
			coord.Col() + 1,
			t.ValueAt(coord.Row(), coord.Col()),
		)
	}
	fmt.Printf("Total cost: %d", sum)
}