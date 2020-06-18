package sort

import (
	"testing"
)

func TestSelectionSortInt(t *testing.T) {
	values, expectedValues := loadComparableIntSlice()
	Selection(values, intLess, intSwapper)
	expectSliceEqual(t, expectedValues, values)
}

func TestSelectionSortString(t *testing.T) {
	values, expectedValues := loadComparableStringSlice()
	Selection(values, stringLess, stringSwapper)
	expectSliceEqual(t, expectedValues, values)
}
