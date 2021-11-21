package hungarianmethod_test

import (
	"fmt"
	"testing"

	alg "hungarian_method"
)

func TestStep3(t *testing.T) {
	data := [][][]int64{
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
		{
			{5, 0, 3, 3, 7},
			{6, 8, 0, 0, 0},
			{0, 4, 4, 4, 1},
			{0, 1, 4, 4, 0},
			{2, 0, 1, 1, 1},
		},
	}

	for i := range data {
		t := alg.NewTable(data[i])
		t.Step1()
		minLinesNeededToCoverZeros, covered := t.CoverZeros()
		for !t.IsDone(minLinesNeededToCoverZeros) {
			t.Modify(covered)
			minLinesNeededToCoverZeros, covered = t.CoverZeros()
		}

		fmt.Println(t)
	}
}
