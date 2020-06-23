package pqueue

import "errors"

// PriorityQueue is an implementation of array-based
// binary heap
type PriorityQueue struct {
	keys []int
	size int
}

// New return an instance of PriorityQueue with the
// given initial capacity
func New(capacity int) PriorityQueue {
	if capacity < 1 {
		capacity = 8
	}
	pq := PriorityQueue{make([]int, 0, capacity), 0}
	// fill pq.keys[0] because we only need
	// pq.keys[1] until pq.keys[len(pq.keys) - 1]
	pq.keys = append(pq.keys, 0)
	return pq
}

// Insert add a new key
func (pq *PriorityQueue) Insert(i int) {
	// push a new node
	pq.keys = append(pq.keys, i)
	pq.size++
	// restore the heap order by swimming
	pq.swim(len(pq.keys) - 1)
}

// ErrEmptyPriorityQueue error occurs when trying to invoke
// DelMax on an empty priority queue
var ErrEmptyPriorityQueue = errors.New("PriorityQueue is empty")

// DelMax remove the largest key
func (pq *PriorityQueue) DelMax() (int, error) {
	if pq.size > 0 {
		// pop the largest key at root node
		v := pq.keys[1]
		// move the last node to the root
		pq.keys[1] = pq.keys[len(pq.keys)-1]
		pq.keys = pq.keys[:len(pq.keys)-1]
		pq.size--
		// restore heap order by sinking
		pq.sink(1)
		return v, nil
	}
	return 0, ErrEmptyPriorityQueue
}

func (pq *PriorityQueue) swim(k int) {
	// continue to swim when k != 1 i.e. haven't reach root
	if k != 1 {
		// swap with current node's parents if current node
		// value is greater than parent's and continue to
		// swim, otherwise we have restored the heap order
		if pq.keys[k] > pq.keys[k/2] {
			pq.keys[k], pq.keys[k/2] = pq.keys[k/2], pq.keys[k]
			pq.swim(k / 2)
		}
	}
}

func (pq *PriorityQueue) sink(k int) {
	// continue to sink if not yet reached bottom
	j := 2 * k
	if j <= pq.size {
		// pick the greater among children if there is one
		if j < pq.size && pq.keys[j] < pq.keys[j+1] {
			j++
		}
		// swap with the greater child and continue to
		// attempt to sink
		if pq.keys[k] < pq.keys[j] {
			pq.keys[k], pq.keys[j] = pq.keys[j], pq.keys[k]
			pq.sink(j)
		}
	}
}
