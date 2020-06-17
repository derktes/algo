package sort

// Selection implements the standard selection sort algoirthm which
// is an inefficient algorithm for sorting large number of keys
func Selection(val []interface{}, cmp Comparator, swp Swapper) {
	if len(val) < 1 {
		return
	}
	for i := 0; i < len(val)-1; i++ {
		kMin := i
		for j := i + 1; j < len(val); j++ {
			if cmp(val[j], val[kMin]) < 0 {
				kMin = j
			}
		}
		swp(val, i, kMin)
	}
}
