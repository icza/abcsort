/*

Package abcsort is a sorting library that uses a custom alphabet.

*/
package abcsort

// Weights returns a map containing runes of the given alphabet, position of
// each rune is used as the weight.
func Weights(alphabet string) map[rune]int {
	// TODO
	return nil
}

// WeightsFold returns a map containing lower and upper versions of all runes
// of the given alphabet, position of each rune is used as the weight.
func WeightsFold(alphabet string) map[rune]int {
	// TODO
	return nil
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
