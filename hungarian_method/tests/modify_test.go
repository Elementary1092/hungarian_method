package hungarianmethod_test

import (
	"fmt"
	"testing"

	alg "hungarian_method"
)

func TestModify(t *testing.T) {
	data := [][][]int64{
		{
			{7, 11, 3, 0},
			{0, 11, 13, 13},
			{13, 0, 0, 0},
			{9, 12, 11, 0},
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
		_, covered := t.CoverZeros()
		t.Modify(covered)
		minLines, _ := t.CoverZeros()
		fmt.Println(minLines)
	}
}