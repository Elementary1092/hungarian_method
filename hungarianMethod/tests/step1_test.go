package hungarianmethod_test

import (
	"fmt"
	"testing"

	alg "hungarianMethod"
)

func TestStep1(t *testing.T) {
	data := [][][]int64{
		{
			{1, 5, 8, 9},
			{6, 7, 56, 15},
			{2, 57, 45, 12},
			{15, 21, 15, 16},
		},
		{	
			{1, 5, 8, 9},
			{6, 7, 56, 15},
			{2, 57, 45, 12},
		},
		{
			{1, 5, 8},
			{6, 7, 15},
			{2, 57, 45},
			{15, 21, 16},
		},
	}

	for i := range data {
		table := alg.NewTable(data[i])
		table.Step1()
		fmt.Println(table)
	}
}