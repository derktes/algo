package pqueue

import (
	"testing"
	//"fmt"
)


type tuple2 [2]int

func TestIndexMinPQInsertion(t *testing.T) {
	type testCase struct {
		in []tuple2
		want [3][]int
	}
	cases := []testCase{
		{
			[]tuple2{{1, 1}, {2, 2}},
			[3][]int{
				{0, 1, 2, 0, 0, 0, 0, 0, 0},
				{-1, 1, 2, -1, -1, -1, -1, -1},
				{0, 1, 2, 0, 0, 0, 0, 0},
			},
		},
		{
			[]tuple2{{2, 7}, {4, 10}, {6, 13}, {7, 8}},
			[3][]int{
				{0, 2, 7, 6, 4, 0, 0, 0, 0},
				{-1, -1, 1, -1, 2, -1, 3, 4},
				{0, 0, 7, 0, 10, 0, 13, 8},
			},
		},
	}
	for _, c := range cases {
		p := NewIndexMinPQ(8)
		// test insertion with c.in
		for _, k := range c.in {
			p.Insert(k[0], k[1])
			//fmt.Printf("p.pq %v\n", p.pq)
			//fmt.Printf("p.keys %v\n", p.keys)
		}
		assertEqual(t, c.want[0], p.pq)
		assertEqual(t, c.want[1], p.qp)
		assertEqual(t, c.want[2], p.keys)
	}

}

func TestIndexMinPQContains(t *testing.T) {
	type testCase struct {
		in   []tuple2
		want []bool
	}
	cases := []testCase{
		{
			[]tuple2{{1, 10}, {2, 12}},
			[]bool{
				false, true, true, false,
				false, false, false, false,
			},
		},
		{
			[]tuple2{{3, 15}, {0, 9}, {7, 100}, {6, 19}},
			[]bool{
				true, false, false, true,
				false, false, true, true,
			},
		},
	}
	for _, c := range cases {
		p := NewIndexMinPQ(8)
		max := -1
		for _, tp := range c.in {
			p.Insert(tp[0], tp[1])
			if max < tp[0] {
				max = tp[0]
			}
		}
		for i := 0; i < max+1; i++ {
			if p.Contains(i) != c.want[i] {
				t.Logf("p.pq %v\n", p.pq)
				t.Logf("p.qp %v\n", p.qp)
				t.Logf("p.keys %v\n", p.keys)
				t.Fatalf("For %v, no key found at %d\n", c.in, i)
			}
		}
	}
}

func TestIndexMinPQMinKey(t *testing.T) {
	type testCase struct {
		in []tuple2
		want int
	}
	cases := []testCase{
		{[]tuple2{{5, 13}, {9, 7}}, 7},
	}
	for _, c := range cases {
		p := NewIndexMinPQ(10)
		for _, tp := range c.in {
			p.Insert(tp[0], tp[1])
		}
		min, _ := p.MinKey()
		if min != c.want {
			t.Fatalf("For %v, minkey not %d but %d\n", c.in, c.want, min)
		}
	}
}

func TestIndexMinPQDelMin(t *testing.T) {
	type testCase struct {
		in []tuple2
		want [3][]int
	}
	cases := []testCase{
		{
			[]tuple2{{2, 1}, {1, 2}},
			[3][]int{
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{-1, 2, -1, -1, -1, -1, -1, -1},
				{0, 2, 1, 0, 0, 0, 0, 0},
			},
		},
		{
			[]tuple2{{7, 12}, {3, 19}, {0, 3}, {5, 1}},
			[3][]int{
				{0, 0, 3, 7, 0, 0, 0, 0, 0},
				{3, -1, -1, 2, -1, -1, -1, 1},
				{3, 0, 0, 19, 0, 1, 0, 12},
			},
		},
	}
	for _, c := range cases {
		p := NewIndexMinPQ(8)
		for _, tp := range c.in {
			p.Insert(tp[0], tp[1])
		}
		p.DelMin()
		assertEqual(t, c.want[0], p.pq)
		assertEqual(t, c.want[1], p.qp)
		assertEqual(t, c.want[2], p.keys)
	}

}


