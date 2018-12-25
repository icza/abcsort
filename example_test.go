package abcsort_test

import (
	"fmt"

	"github.com/icza/abcsort"
)

// Example shows how to do custom sorting using the Sorter helper type.
func Example() {
	sorter := abcsort.New("bac")

	ss := []string{"abc", "bac", "cba", "CCC"}
	sorter.Strings(ss)
	fmt.Println(ss)

	ss = []string{"abc", "bac", "cba", "CCC"}
	sorter.StringsFold(ss)
	fmt.Println(ss)

	type Person struct {
		Name string
		Age  int
	}

	ps := []Person{{Name: "alice", Age: 21}, {Name: "bob", Age: 12}}
	sorter.Slice(ps, func(i int) string { return ps[i].Name })
	fmt.Println(ps)

	ps = []Person{{Name: "Alice", Age: 21}, {Name: "Bob", Age: 12}}
	sorter.SliceFold(ps, func(i int) string { return ps[i].Name })
	fmt.Println(ps)

	// Output:
	// [CCC bac abc cba]
	// [bac abc cba CCC]
	// [{bob 12} {alice 21}]
	// [{Bob 12} {Alice 21}]
}
