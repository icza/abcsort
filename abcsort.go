/*

Package abcsort is a sorting library that uses a custom alphabet.

*/
package abcsort

import (
	"strings"
)

// Weights returns a map containing runes of the given alphabet, position of
// each rune is used as the weight.
func Weights(alphabet string) map[rune]int {
	weights := map[rune]int{}
	addWeights(alphabet, weights)
	return weights
}

// WeightsFold returns a map containing lower and upper versions of all runes
// of the given alphabet, position of each rune is used as the weight.
func WeightsFold(alphabet string) map[rune]int {
	weights := map[rune]int{}
	addWeights(strings.ToLower(alphabet), weights)
	addWeights(strings.ToUpper(alphabet), weights)
	return weights
}

// addWeights uses the given map to add the weights to.
func addWeights(alphabet string, weights map[rune]int) {
	for i, r := range alphabet {
		// NOTE: i is the byte index (not the rune index), but since it is also
		// monotone increasing, it will just just as fine for rune weight.
		if r != 0xfffd {
			weights[r] = i
		}
	}
}

// Less tells if s1 is less than s2, interpreted as UTF-8 strings,
// using the given weights (which should be constructed using Weights()).
func Less(s1, s2 string, weights map[rune]int) bool {
	// TODO
	return false
}

// LessFold tells if s1 is less than s2 under Unicode case folding,
// interpreted as UTF-8 strings, using the given weights (which should be
// constructed using WeightsFold()).
func LessFold(s1, s2 string, weights map[rune]int) bool {
	// TODO
	return false
}
