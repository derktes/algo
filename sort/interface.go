package sort

// Comparator must be implemented to provide type specific
// comparision.
type Comparator func(interface{}, interface{}) int

// Swapper must be implemented to provide type specific
// exchange of value.
type Swapper func([]interface{}, int, int)
