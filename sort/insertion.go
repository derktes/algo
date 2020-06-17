package sort

// Insertion implements insertion sort
func Insertion(val []interface{}, cmp Comparator, swp Swapper) {
	if len(val) < 1 {
		return
	}
	for i := 1; i < len(val); i++ {
		for j := i; j > 0; j-- {
			if cmp(val[j-1], val[j]) < 0 {
				break
			}
			swp(val, j-1, j)
		}
	}
}
