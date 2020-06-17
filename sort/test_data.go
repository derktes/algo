package sort

func intComparator(a, b interface{}) int {
	x := a.(int)
	y := b.(int)
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}

func intSwapper(v []interface{}, i, j int) {
	v[i], v[j] = v[j], v[i]
}

func loadComparableIntSlice() (values, expectedValues []interface{}) {
	n := []int{5, 4, 3, 2, 1}
	m := []int{1, 2, 3, 4, 5}
	for _, i := range n {
		values = append(values, i)
	}
	for _, i := range m {
		expectedValues = append(expectedValues, i)
	}
	return
}

func stringComparator(a, b interface{}) int {
	x := a.(string)
	y := b.(string)
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}

func stringSwapper(v []interface{}, i, j int) {
	v[i], v[j] = v[j], v[i]
}

func loadComparableStringSlice() (values, expectedValues []interface{}) {
	n := []string{"command", "keyboard", "angular", "mern", "browserify"}
	m := []string{"angular", "browserify", "command", "keyboard", "mern"}
	for _, i := range n {
		values = append(values, i)
	}
	for _, i := range m {
		expectedValues = append(expectedValues, i)
	}
	return
}
