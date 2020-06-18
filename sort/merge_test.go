package sort

import "testing"

func TestMergeSort(t *testing.T) {
	cases := loadIntCases()
	for _, c := range cases {
		Merge(c.in, func(a, b interface{}) bool {
			x := a.(int)
			y := b.(int)
			return x < y
		}, nil)
		expectSliceEqual(t, c.want, c.in)
	}
}
