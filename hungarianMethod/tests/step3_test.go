package hungarianmethod_test

import (
	"fmt"
	"testing"

	alg "hungarianmethod"
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
