package pqueue

import "testing"

type testCase struct {
	in   []int
	want []int
}

func assertEqual(t *testing.T, want []int, computed []int) {
	equal := true
	if len(want) != len(computed) {
		equal = false
	}
	if equal {
		for i := range want {
			if want[i] != computed[i] {
				equal = false
				break
			}
		}
	}
	if !equal {
		t.Fatalf("Slice not equal, expected %v but got %v\n", want, computed)
	}
}

func TestPriorityQueueInsertion(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2}, []int{2, 1}},
		{[]int{2, 1}, []int{2, 1}},
		{[]int{2, 6, 1, 9}, []int{9, 6, 1, 2}},
	}
	for _, c := range cases {
		pq := New(8)
		for _, i := range c.in {
			pq.Insert(i)
		}
		assertEqual(t, c.want, pq.keys[1:])
	}
}

func TestPriorityQueueDelMax(t *testing.T) {
	cases := []testCase{
		{[]int{0, 2}, []int{}},
		{[]int{0, 2, 1}, []int{1}},
		{[]int{0, 9, 6, 1, 2}, []int{6, 2, 1}},
		{[]int{0, 11, 4, 8, 1, 2, 3}, []int{8, 4, 3, 1, 2}},
	}
	for _, c := range cases {
		pq := PriorityQueue{c.in, len(c.in) - 1}
		pq.DelMax()
		assertEqual(t, c.want, pq.keys[1:])
	}
}

func TestEmptyPriorityQueueDelMax(t *testing.T) {
	pq := New(8)
	_, err := pq.DelMax()
	if err != ErrEmptyPriorityQueue {
		t.Fatalf("DelMax on empty PriorityQueue, expects %v, but got %v\n", ErrEmptyPriorityQueue, err)
	}
}
