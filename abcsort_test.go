package abcsort

import (
	"reflect"
	"testing"
)

func TestWeights_WeightsFold(t *testing.T) {
	cases := []struct {
		title       string
		alphabet    string
		weights     map[rune]int
		weightsFold map[rune]int
	}{
		{
			title:       "empty",
			alphabet:    "",
			weights:     map[rune]int{},
			weightsFold: map[rune]int{},
		},
		{
			title:       "normal",
			alphabet:    "abc",
			weights:     map[rune]int{'a': 0, 'b': 1, 'c': 2},
			weightsFold: map[rune]int{'a': 0, 'b': 1, 'c': 2, 'A': 0, 'B': 1, 'C': 2},
		},
		{
			title:       "mixed",
			alphabet:    "aBc",
			weights:     map[rune]int{'a': 0, 'B': 1, 'c': 2},
			weightsFold: map[rune]int{'a': 0, 'b': 1, 'c': 2, 'A': 0, 'B': 1, 'C': 2},
		},
		{
			title:       "upper",
			alphabet:    "ABC",
			weights:     map[rune]int{'A': 0, 'B': 1, 'C': 2},
			weightsFold: map[rune]int{'a': 0, 'b': 1, 'c': 2, 'A': 0, 'B': 1, 'C': 2},
		},
		{
			title:       "invalid",
			alphabet:    "\xfa\xfb\xfc",
			weights:     map[rune]int{},
			weightsFold: map[rune]int{},
		},
	}

	for _, c := range cases {
		weights := Weights(c.alphabet)
		if !reflect.DeepEqual(weights, c.weights) {
			t.Errorf("[%s] Expected weights: %v, got: %v", c.title, c.weights, weights)
		}

		weightsFold := WeightsFold(c.alphabet)
		if !reflect.DeepEqual(weightsFold, c.weightsFold) {
			t.Errorf("[%s] Expected weights fold: %v, got: %v", c.title, c.weightsFold, weightsFold)
		}
	}
}
