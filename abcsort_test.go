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

func TestLess_LessFold(t *testing.T) {
	const alphabet = "cdb"
	weights := Weights(alphabet)
	weightsFold := WeightsFold(alphabet)

	cases := []struct {
		title          string
		s1, s2         string
		less, lessFold bool
	}{
		{
			title: "empty < empty",
			s1:    "", s2: "",
			less: false, lessFold: false,
		},
		{
			title: `"b" < empty`,
			s1:    "b", s2: "",
			less: false, lessFold: false,
		},
		{
			title: `"x" < empty`,
			s1:    "x", s2: "",
			less: false, lessFold: false,
		},
		{
			title: `empty < "b"`,
			s1:    "", s2: "b",
			less: true, lessFold: true,
		},
		{
			title: `empty < "x"`,
			s1:    "", s2: "x",
			less: true, lessFold: true,
		},
		{
			title: `"b" < "c"`,
			s1:    "b", s2: "c",
			less: false, lessFold: false,
		},
		{
			title: `"c" < "b"`,
			s1:    "c", s2: "b",
			less: true, lessFold: true,
		},
		{
			title: `"d" < "b"`,
			s1:    "d", s2: "b",
			less: true, lessFold: true,
		},
		{
			title: `"c" < "b"`,
			s1:    "c", s2: "b",
			less: true, lessFold: true,
		},
		{
			title: `"cdb" < "cbd"`,
			s1:    "cdb", s2: "cbd",
			less: true, lessFold: true,
		},
		{
			title: `"cx" < "cx"`,
			s1:    "cx", s2: "cx",
			less: false, lessFold: false,
		},
		{
			title: `"cx" < "cy"`,
			s1:    "cx", s2: "cy",
			less: true, lessFold: true,
		},
		{
			title: `"cy" < "cx"`,
			s1:    "cy", s2: "cx",
			less: false, lessFold: false,
		},
		{
			title: `"cb" < "cba"`,
			s1:    "cb", s2: "cba",
			less: true, lessFold: true,
		},
		{
			title: `"cb" < "cbx"`,
			s1:    "cb", s2: "cbx",
			less: true, lessFold: true,
		},
		{
			title: `"cba" < "cb"`,
			s1:    "cba", s2: "cb",
			less: false, lessFold: false,
		},
		{
			title: `"cbx" < "cb"`,
			s1:    "cbx", s2: "cb",
			less: false, lessFold: false,
		},
		{
			title: `"x" < "x"`,
			s1:    "x", s2: "x",
			less: false, lessFold: false,
		},
		{
			title: `"x" < "y"`,
			s1:    "x", s2: "y",
			less: true, lessFold: true,
		},
		{
			title: `"y" < "x"`,
			s1:    "y", s2: "x",
			less: false, lessFold: false,
		},
		{
			title: `"xc" < "xb"`,
			s1:    "xc", s2: "xb",
			less: true, lessFold: true,
		},
		{
			title: `"xb" < "xc"`,
			s1:    "xb", s2: "xc",
			less: false, lessFold: false,
		},
		{
			title: `"xcy" < "xby"`,
			s1:    "xcy", s2: "xby",
			less: true, lessFold: true,
		},
		{
			title: `"xcz" < "xby"`,
			s1:    "xcz", s2: "xby",
			less: true, lessFold: true,
		},
		{
			title: `"xcy" < "xbz"`,
			s1:    "xcy", s2: "xbz",
			less: true, lessFold: true,
		},
		{
			title: `"xby" < "xcy"`,
			s1:    "xby", s2: "xcy",
			less: false, lessFold: false,
		},
		{
			title: `"xbz" < "xcy"`,
			s1:    "xbz", s2: "xcy",
			less: false, lessFold: false,
		},
		{
			title: `"xby" < "xcz"`,
			s1:    "xby", s2: "xcz",
			less: false, lessFold: false,
		},
	}

	for _, c := range cases {
		less := Less(c.s1, c.s2, weights)
		if less != c.less {
			t.Errorf("[%s] Expected less: %t, got: %t", c.title, c.less, less)
		}

		lessFold := LessFold(c.s1, c.s2, weightsFold)
		if lessFold != c.lessFold {
			t.Errorf("[%s] Expected less fold: %t, got: %t", c.title, c.lessFold, lessFold)
		}
	}
}

func TestLessFold_2(t *testing.T) {
	const alphabet = "bCa"
	weights := Weights(alphabet)
	weightsFold := WeightsFold(alphabet)

	cases := []struct {
		title          string
		s1, s2         string
		less, lessFold bool
	}{
		{
			title: `"cb" < "Ca"`,
			s1:    "cb", s2: "Ca",
			less: false, lessFold: true,
		},
	}

	for _, c := range cases {
		less := Less(c.s1, c.s2, weights)
		if less != c.less {
			t.Errorf("[%s] Expected less: %t, got: %t", c.title, c.less, less)
		}

		lessFold := LessFold(c.s1, c.s2, weightsFold)
		if lessFold != c.lessFold {
			t.Errorf("[%s] Expected less fold: %t, got: %t", c.title, c.lessFold, lessFold)
		}
	}
}
