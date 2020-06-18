package sort

import "testing"

func TestShellSortInt(t *testing.T) {
	values, expectedValues := loadComparableIntSlice()
	Shell(values, intComparator, intSwapper)
	expectSliceEqual(t, expectedValues, values)
}

func TestShellSortString(t *testing.T) {
	values, expectedValues := loadComparableStringSlice()
	Shell(values, stringComparator, stringSwapper)
	expectSliceEqual(t, expectedValues, values)
}
