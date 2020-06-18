package sort

// Comparator must be implemented to provide type specific
// comparision.
type Comparator func(interface{}, interface{}) int

// SwapFunc must be implemented to provide type specific
// exchange of value.
type SwapFunc func([]interface{}, int, int)

// LessFunc must be implemented to provide type specific
// element pair comparison
type LessFunc func(interface{}, interface{}) bool
