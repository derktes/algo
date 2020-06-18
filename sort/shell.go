package sort

// Shell implement shell sorting algorithm
func Shell(values []interface{}, cmp Comparator, swp Swapper) {
	// compute initial increment sequence value k
	t := len(values) / 3
	h := 1
	for h < t {
		h = 3*h + 1
	}
	// h-sort each subsequence of the array
	for h >= 1 {
		for i := h; i < len(values); i++ {
			for j := i; j >= h; j -= h {
				if cmp(values[j-h], values[j]) < 0 {
					break
				}
				swp(values, j-h, j)
			}
		}
		h = h / 3
	}
}
