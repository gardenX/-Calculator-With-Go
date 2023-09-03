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

func TestCalc(t *testing.T) {
	type table struct {
		first    MicroValue
		operator string
		second   MicroValue
		expected MicroValue
	}

	step := []table{
		{first: 0, second: 300, operator: "+", expected: 300},
		{first: 300, second: 100, operator: "-", expected: 200},
		{first: 200, second: 200, operator: "*", expected: 400},
		{first: 400, second: 100, operator: "+", expected: 500},
		{first: 500, second: 200, operator: "/", expected: 250},
	}

	for _, test := range step {
		res := Calc(test.first, test.operator, test.second)

		if res != test.expected {
			t.Error("assert", test.first, test.operator, test.second, "equal to", test.expected, "but got", res)
		}
	}
}

func TestDisplayMacro(t *testing.T) {
	table := map[MicroValue]float32{100: 1, 200: 2, 350: 3.5, 442: 4.42}

	for i, test := range table {
		res := DisplayMacro(i)

		if res != test {
			t.Error("expected display macro for", i, "is", test, "but got", res)
		}
	}
}

func TestToMacro(t *testing.T) {
	table := map[MicroValue]int{100: 1, 200: 2, 330: 3}

	for i, value := range table {
		res := ToMacro(i)

		if res != value {
			t.Error("assert macro of", i, "are", value, "but got", res)
		}
	}
}

func TestToMicro(t *testing.T) {
	table := map[int]MicroValue{1: 100, 2: 200}

	for i, value := range table {
		res := ToMicro(i)

		if res != value {
			t.Error("assert micro of", i, "are", value, "but got", res)
		}
	}
}
