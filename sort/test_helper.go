package sort

import "testing"

func expectSliceEqual(t *testing.T, expected []interface{}, computed []interface{}) {
	isEqual := true
	if len(computed) != len(expected) {
		isEqual = false
	}
	if len(computed) <= 0 || len(expected) <= 0 {
		isEqual = false
	}
	for i, v := range computed {
		if v != expected[i] {
			isEqual = false
			break
		}
	}
	if !isEqual {
		t.Fatalf("Expected %v but got %v\n", expected, computed)
	}
}
