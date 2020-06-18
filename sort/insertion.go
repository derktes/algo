package sort

// Insertion implements insertion sort
func Insertion(val []interface{}, less LessFunc, swap SwapFunc) {
	if len(val) < 1 {
		return
	}
	for i := 1; i < len(val); i++ {
		for j := i; j > 0; j-- {
			if less(val[j-1], val[j]) {
				break
			}
			swap(val, j-1, j)
		}
	}
}
