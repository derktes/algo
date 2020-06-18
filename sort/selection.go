package sort

// Selection implements the standard selection sort algoirthm which
// is an inefficient algorithm for sorting large number of keys
func Selection(val []interface{}, less LessFunc, swap SwapFunc) {
	if len(val) < 1 {
		return
	}
	for i := 0; i < len(val)-1; i++ {
		kMin := i
		for j := i + 1; j < len(val); j++ {
			if less(val[j], val[kMin]) {
				kMin = j
			}
		}
		swap(val, i, kMin)
	}
}
