package calculator

import "testing"

var AddResultSet = []struct{ Input1, Input2, Result int }{{1, 2, 3}, {4, 5, 9}}
var SubtractResultSet = []struct{ Input1, Input2, Result int }{{3, 2, 1}, {6, 3, 3}}

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
