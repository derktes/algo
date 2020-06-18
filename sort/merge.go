package sort

// Merge performs merge sort
func Merge(val []interface{}, less LessFunc, swap SwapFunc) {
	if len(val) < 2 {
		return
	}
	// compute mid and (mid-1)
	mid := len(val) / 2
	Merge(val[:mid], less, swap)
	Merge(val[mid:], less, swap)
	//fmt.Printf("sub-merge: %v\n", val[:mid])
	//fmt.Printf("sub-merge: %v\n", val[mid:])
	valCopy := make([]interface{}, len(val))
	copy(valCopy, val)
	k := 0
	i := 0
	j := mid
	for i < mid && j < len(val) {
		if less(valCopy[i], valCopy[j]) {
			val[k] = valCopy[i]
			i++
		} else {
			val[k] = valCopy[j]
			j++
		}
		k++
	}
	for i < mid {
		val[k] = valCopy[i]
		i++
		k++
	}
	for j < len(val) {
		val[k] = valCopy[j]
		j++
		k++
	}
	//fmt.Printf("%v\n", val)
}
