// This file contains the Sorter helper type.

package abcsort

import "sort"

// Sorter provides functionality to easily sort slices using a custom alphabet.
// A sorter is safe for concurrent use.
type Sorter struct {
	// Alphabet is the custom alphabet
	Alphabet string
	// Weights is used for "normal" sorting.
	Weights map[rune]int
	// WeightsFold is used for fold sorting.
	WeightsFold map[rune]int
}

// New returns a new Sorter that will use the given custom alphabet when sorting.
func New(alphabet string) *Sorter {
	return &Sorter{
		Alphabet:    alphabet,
		Weights:     Weights(alphabet),
		WeightsFold: WeightsFold(alphabet),
	}
}

// Strings sorts a string slice.
func (s *Sorter) Strings(ss []string) {
	sort.Sort(&StringSlice{Slice: ss, Weights: s.Weights})
}

// StringsFold sorts a string slice under Unicode folding.
func (s *Sorter) StringsFold(ss []string) {
	sort.Sort(&StringFoldSlice{Slice: ss, Weights: s.Weights})
}

// Slice sorts a slice.
// getField is a function that must return the string value (e.g. a field) of
// the element at the ith index.
func (s *Sorter) Slice(slice interface{}, getField func(i int) string) {
	sort.Slice(slice, func(i int, j int) bool {
		return Less(getField(i), getField(j), s.Weights)
	})
}

// SliceFold sorts a slice under Unicode folding.
// getField is a function that must return the string value (e.g. a field) of
// the element at the ith index.
func (s *Sorter) SliceFold(slice interface{}, getField func(i int) string) {
	sort.Slice(slice, func(i int, j int) bool {
		return LessFold(getField(i), getField(j), s.WeightsFold)
	})
}

// StringSlice is a helper struct that implements sort.Interface.
type StringSlice struct {
	Slice   []string
	Weights map[rune]int
}

func (ss *StringSlice) Len() int           { return len(ss.Slice) }
func (ss *StringSlice) Less(i, j int) bool { return Less(ss.Slice[i], ss.Slice[j], ss.Weights) }
func (ss *StringSlice) Swap(i, j int)      { ss.Slice[i], ss.Slice[j] = ss.Slice[j], ss.Slice[i] }

// StringFoldSlice is a helper struct that implements sort.Interface.
type StringFoldSlice struct {
	Slice   []string
	Weights map[rune]int
}

func (ss *StringFoldSlice) Len() int           { return len(ss.Slice) }
func (ss *StringFoldSlice) Less(i, j int) bool { return LessFold(ss.Slice[i], ss.Slice[j], ss.Weights) }
func (ss *StringFoldSlice) Swap(i, j int)      { ss.Slice[i], ss.Slice[j] = ss.Slice[j], ss.Slice[i] }
