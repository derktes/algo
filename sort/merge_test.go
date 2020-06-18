package sort

import "testing"

func TestMergeSort(t *testing.T) {
	cases := loadIntCases()
	for _, c := range cases {
		Merge(c.in, func(a, b interface{}) bool {
			return a.(int) < b.(int)
		})
		expectSliceEqual(t, c.want, c.in)
	}
}

func TestMergeSortString(t *testing.T) {
	cases := loadStringCases()
	for _, c := range cases {
		Merge(c.in, func(a, b interface{}) bool {
			return a.(string) < b.(string)
		})
		expectSliceEqual(t, c.want, c.in)
	}
}
