package calculator

import "testing"

var AddResultSet = []struct{ Input1, Input2, Result int }{{1, 2, 3}, {4, 5, 9}}
var SubtractResultSet = []struct{ Input1, Input2, Result int }{{3, 2, 10}, {6, 3, 3}}
var MultiplyResultSet = []struct{ Input1, Input2, Result int }{{3, 2, 6}, {6, 3, 18}}

func TestAdd(t *testing.T) {
	for _, v := range AddResultSet {
		result := Add(v.Input1, v.Input2)
		expected := v.Result

		if result != expected {
			t.Fail()
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, v := range SubtractResultSet {
		result := Subtract(v.Input1, v.Input2)
		expected := v.Result

		if result != expected {
			t.Fail()
		}
	}
}

func TestMultiply(t *testing.T) {
	for _, v := range MultiplyResultSet {
		result := Multiply(v.Input1, v.Input2)
		expected := v.Result

		if result != expected {
			t.Fail()
		}
	}
}
