package input_data

import (
	"fmt"
)

type CMDInput struct {}

func (c *CMDInput) GetData() [][]int64 {
	var n int
	fmt.Print("Input number of works and workers: ")
	fmt.Scan(&n)

	data := make([][]int64, n)
	for i := 0; i < n; i++ {
		row := make([]int64, n)
		for j := 0; j < n; j++ {
			fmt.Printf(
				"Enter cost of the work %d if %dth worker does this work: ", 
				i, 
				j,
			)
			
			fmt.Scan(&row[j])
		}

		data[i] = row
	}

	return data
}