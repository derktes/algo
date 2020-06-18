package sort

import "testing"

func TestInsertionSortInt(t *testing.T) {
	values, expectedValues := loadComparableIntSlice()
	Insertion(values, intLess, intSwapper)
	expectSliceEqual(t, expectedValues, values)
}

func TestInsertionSortString(t *testing.T) {
	values, expectedValues := loadComparableStringSlice()
	Insertion(values, stringLess, stringSwapper)
	expectSliceEqual(t, expectedValues, values)
}
