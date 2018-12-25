package abcsort

import (
	"fmt"
	"sort"
)

func ExampleStringSlice() {
	weights := Weights("bac")

	ss := []string{"abc", "bac", "cba", "CCC"}
	strslice := &StringSlice{
		Weights: weights,
		Slice:   ss,
	}
	sort.Sort(strslice)
	fmt.Println(ss)

	// Output:
	// [CCC bac abc cba]
}
