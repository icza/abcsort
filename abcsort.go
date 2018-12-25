/*

Package abcsort is a string sorting library that uses a custom,user-defined
alphabet.

*/
package abcsort

import (
	"strings"
	"unicode"
	"unicode/utf8"
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
	return less(s1, s2, weights, false)
}

// LessFold tells if s1 is less than s2 under Unicode case folding,
// interpreted as UTF-8 strings, using the given weights (which should be
// constructed using WeightsFold()).
func LessFold(s1, s2 string, weights map[rune]int) bool {
	return less(s1, s2, weights, true)
}

// less tells if s1 is less than s2, interpreted as UTF-8 strings,
// using the given weights.
//
// If fold is true comparison is performed under Unicode case folding
// in which case weights should be constructed using WeightsFold(), otherwise
// using Weights().
func less(s1, s2 string, weights map[rune]int, fold bool) bool {
	for {
		switch e1, e2 := len(s1) == 0, len(s2) == 0; {
		case e1 && e2:
			return false // Both empty, they are equal (not less)
		case !e1 && e2:
			return false // s1 not empty but s2 is: s1 is greater (not less)
		case e1 && !e2:
			return true // s1 empty but s2 is not: s1 is less
		}

		r1, size1 := utf8.DecodeRuneInString(s1)
		r2, size2 := utf8.DecodeRuneInString(s2)

		// Check if both are custom, in which case we use custom order:
		var (
			w1     int
			custom bool
		)
		if w1, custom = weights[r1]; custom {
			var w2 int
			if w2, custom = weights[r2]; custom {
				if w1 != w2 {
					return w1 < w2
				}
			}
		}
		if !custom {
			// Fallback to numeric rune comparison.
			// If fold, transform runes:
			if fold {
				r1, r2 = unicode.ToLower(r1), unicode.ToLower(r2)
			}
			if r1 != r2 {
				return r1 < r2
			}
		}

		// Cut off first runes, and continue:
		s1, s2 = s1[size1:], s2[size2:]
	}
}
