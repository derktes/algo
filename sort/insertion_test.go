package sort

import "testing"

func TestInsertionSortInt(t *testing.T) {
	values, expectedValues := loadComparableIntSlice()
	Insertion(values, intComparator, intSwapper)
	expectSliceEqual(t, expectedValues, values)
}

func TestInsertionSortString(t *testing.T) {
	values, expectedValues := loadComparableStringSlice()
	Insertion(values, stringComparator, stringSwapper)
	expectSliceEqual(t, expectedValues, values)
}
