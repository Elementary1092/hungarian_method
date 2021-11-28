package main

import (
	out "display_results"
	alg "hungarian_method"
	in "input_data"
)

func main() {
	input := &in.CMDInput{}
	data := input.GetData()

	t := alg.NewTable(data)
	out.Display(t, t.Solve())
}