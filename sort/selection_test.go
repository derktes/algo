package sort

import (
	"testing"
)

func TestSelectionSortInt(t *testing.T) {
	cases := loadIntCases()
	for _, c := range cases {
		Selection(c.in,
			func(a, b interface{}) bool {
				return a.(int) < b.(int)
			},
			func(v []interface{}, i, j int) {
				v[i], v[j] = v[j], v[i]
			})
		expectSliceEqual(t, c.want, c.in)
	}
}

func TestSelectionSortString(t *testing.T) {
	cases := loadStringCases()
	for _, c := range cases {
		Selection(c.in,
			func(a, b interface{}) bool {
				return a.(string) < b.(string)
			},
			func(v []interface{}, i, j int) {
				v[i], v[j] = v[j], v[i]
			})
		expectSliceEqual(t, c.want, c.in)

	}
}
