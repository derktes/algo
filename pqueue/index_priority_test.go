package pqueue

import (
	"testing"
	//"fmt"
)

type idxMinPQtestCase struct {
	in   [][2]int
	want [3][]int
}

func TestIndexMinPQInsertion(t *testing.T) {
	cases := []idxMinPQtestCase{
		{
			[][2]int{{1, 1}, {2, 2}},
			[3][]int{{2, 1}, {-1, 1, 2}, {0, 1, 2}},
		},
		{
			[][2]int{{2, 7}, {4, 10}, {6, 13}, {8, 8}},
			[3][]int{
				{6, 8, 4, 2},
				{-1, -1, 1, -1, 2, -1, 3, -1, 4, -1, -1, -1},
				{0, 0, 7, 0, 10, 0, 13, 0, 8, 0, 0, 0},
			},
		},
	}
	for _, c := range cases {
		// create an instance of IndexMinPQ by initializing pq, qp and keys
		p := IndexMinPQ{make([]int, 0, 3), make([]int, 3), make([]int, 3), 0}
		// first element of pq is not used, so fill it away
		p.pq = append(p.pq, 0)
		// assign all elements in qp to -1 indicating index not set
		for i := 0; i < len(p.qp); i++ {
			p.qp[i] = -1
		}
		// test insertion with c.in
		for _, k := range c.in {
			p.Insert(k[0], k[1])
			//fmt.Printf("p.pq %v\n", p.pq)
			//fmt.Printf("p.keys %v\n", p.keys)
		}
		assertEqual(t, c.want[0], p.pq[1:])
		assertEqual(t, c.want[1], p.qp)
		assertEqual(t, c.want[2], p.keys)
	}

}
