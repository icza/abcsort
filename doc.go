/*

Package abcsort is a string sorting library that uses a custom, user-defined
alphabet.

Implementation does not convert the input strings into byte or rune slices, so
performance is rather good.

Custom sorting can be easiest achieved by using the Sorter helper type.

The essence of sorting, the less() function required by the standard lib's sort
package is also exposed, and may be used "manually" like this:

	weights := Weights("bac")

	ss := []string{"abc", "bac", "cba", "CCC"}
	sort.Slice(ss, func(i int, j int) bool {
		return Less(ss[i], ss[j], weights)
	})
	fmt.Println(ss)

	// Output:
	// [CCC bac abc cba]

Or via the StringSlice type:

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

*/
package abcsort
