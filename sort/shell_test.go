package sort

import "testing"

func TestShellSortInt(t *testing.T) {
	values, expectedValues := loadComparableIntSlice()
	Shell(values, intLess, intSwapper)
	expectSliceEqual(t, expectedValues, values)
}

func TestShellSortString(t *testing.T) {
	values, expectedValues := loadComparableStringSlice()
	Shell(values, stringLess, stringSwapper)
	expectSliceEqual(t, expectedValues, values)
}
