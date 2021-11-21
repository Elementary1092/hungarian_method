package display_results

import (
	"fmt"
	alg "hungarian_method"
)

func Display(data [][]int64, coords []*alg.Pair) {
	for _, coord := range coords {
		fmt.Printf(
			"Job %d should be done by %d worker and it will cost %d\n",
			coord.Row() + 1,
			coord.Col() + 1,
			data[coord.Row()][coord.Col()],
		)
	}
}