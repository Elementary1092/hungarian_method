package hungarianmethod_test

import (
	"fmt"
	"testing"

	alg "hungarianmethod"
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
		{
			{20, 28, 19, 13},
			{15, 30, 31, 28},
			{40, 21, 20, 17},
			{21, 28, 26, 12},
		},
		{
			{20, 15, 18, 20, 25},
			{18, 20, 12, 14, 15},
			{21, 23, 25, 27, 25},
			{17, 18, 21, 23, 20},
			{18, 18, 16, 19, 20},
		},
	}

	for i := range data {
		table := alg.NewTable(data[i])
		table.Step1()
		fmt.Println(table)
	}
}