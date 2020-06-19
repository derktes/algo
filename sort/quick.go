package sort

func Quick(val []interface{}, less LessFunc, swap SwapFunc) {
	subSort(val, less, swap)
}

func subSort(val []interface{}, less LessFunc, swap SwapFunc) {
	if len(val) < 2 {
		return
	}
	k := 0
	i := 1
	j := len(val) - 1
	for i <= j {
		for i < len(val) {
			if less(val[k], val[i]) {
				break
			}
			i++
		}
		for j > 0 {
			if less(val[j], val[k]) {
				break
			}
			j--
		}
		if i <= j {
			swap(val, i, j)
		} else {
			swap(val, k, j)
		}
	}
	subSort(val[:j], less, swap)
	subSort(val[(j+1):], less, swap)
}
