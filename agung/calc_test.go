package agung

import (
	"testing"
)

type basicCalcTest struct {
	first    int
	second   int
	expected int
}

func TestAdd(t *testing.T) {
	table := []basicCalcTest{
		{first: 1, second: 2, expected: 3},
		{first: 2, second: 3, expected: 5},
	}

	for _, test := range table {
		res := Add(test.first, test.second)

		if res != test.expected {
			t.Error("expected", test.first, "+", test.second, "equal to", test.expected, "but got", res)
		}
	}
}

func TestSub(t *testing.T) {
	table := []basicCalcTest{
		{first: 10, second: 1, expected: 9},
		{first: 1, second: 2, expected: -1},
	}

	for _, test := range table {
		res := Sub(test.first, test.second)

		if res != test.expected {
			t.Error("expected", test.first, "-", test.second, "equal to", test.expected, "but got", res)
		}
	}
}

func TestTime(t *testing.T) {
	table := []basicCalcTest{
		{first: 1, second: 2, expected: 2},
		{first: 10, second: 2, expected: 20},
	}

	for _, test := range table {
		res := Time(test.first, test.second)

		if res != test.expected {
			t.Error("expected", test.first, "*", test.second, "equal to", test.expected, "but got", res)
		}
	}
}

func TestDiv(t *testing.T) {
	defer func() {
		recover()
	}()

	table := []basicCalcTest{
		{first: 10, second: 5, expected: 2},
		{first: 9, second: 3, expected: 3},
		{first: 9, second: 0, expected: 3},
	}

	for _, test := range table {
		res := Div(test.first, test.second)

		if res != test.expected {
			t.Error("expected", test.first, "/", test.second, "equal to", test.expected, "but got", res)
		}
	}

	t.Error("panic not raised")
}
