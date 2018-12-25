# abcsort

[![Build Status](https://travis-ci.org/icza/abcsort.svg?branch=master)](https://travis-ci.org/icza/abcsort)
[![GoDoc](https://godoc.org/github.com/icza/abcsort?status.svg)](https://godoc.org/github.com/icza/abcsort)
[![Go Report Card](https://goreportcard.com/badge/github.com/icza/abcsort)](https://goreportcard.com/report/github.com/icza/abcsort)
[![codecov](https://codecov.io/gh/icza/abcsort/branch/master/graph/badge.svg)](https://codecov.io/gh/icza/abcsort)

Go string sorting library that uses a custom, user-defined alphabet.

Implementation does not convert the input strings into byte or rune slices, so
performance is rather good.

Custom sorting can be easiest achieved by using the Sorter helper type, for example:

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


The essence of sorting, the `less()` function required by the standard lib's `sort`
package is also exposed, and may be used "manually" like this:

	weights := Weights("bac")

	ss := []string{"abc", "bac", "cba", "CCC"}
	sort.Slice(ss, func(i int, j int) bool {
		return Less(ss[i], ss[j], weights)
	})
	fmt.Println(ss)

	// Output:
	// [CCC bac abc cba]
