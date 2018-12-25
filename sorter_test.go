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

func ExampleStringFoldSlice() {
	ss := []string{"bármi", "Áron", "áram"}
	sort.Strings(ss)
	fmt.Println(ss)

	weights := WeightsFold("aábcdeéfghiíjklmnoóöőpqrstuúüűvwxyz")

	ss = []string{"bármi", "Áron", "áram"}
	strslice := &StringFoldSlice{
		Weights: weights,
		Slice:   ss,
	}
	sort.Sort(strslice)
	fmt.Println(ss)

	// Output:
	// [bármi Áron áram]
	// [áram Áron bármi]
}
