package pqueue

import "errors"

var (
	ErrNoMinKeyFound = errors.New("No min key found")
	ErrIndexOutOfBound = errors.New("Index out of bound")
)

// IndexMinPQ is a variant of minimum-oriented
// priority queue that is indexed
type IndexMinPQ struct {
	pq   []int
	qp   []int
	keys []int
	size int
}

// New construct a new instance of IndexMinPQ
func NewIndexMinPQ(capacity int) IndexMinPQ {
	if capacity < 1 {
		capacity = 20
	}
	// create an instance of IndexMinPQ by initializing pq, qp and keys
	p := IndexMinPQ{
		make([]int, capacity + 1),
		make([]int, capacity),
		make([]int, capacity), 0,
	}
	// assign all elements in qp to -1 indicating index not set
	for i := 0; i < len(p.qp); i++ {
		p.qp[i] = -1
	}
	return p
}

func (p *IndexMinPQ) Size() int {
	return p.size
}

// Insert puts a new key at index i
func (p *IndexMinPQ) Insert(i int, key int) error {
	p.size++
	if i >= 0 && i < len(p.keys) {
		p.pq[p.size] = i
		p.qp[i] = p.size
		p.keys[i] = key
		p.swim(p.size)
		return nil
	}
	return ErrIndexOutOfBound
	//p.qp = insertInverse(p.qp, i, p.size)
	//fmt.Printf("p.qp %v\n", p.qp)
	//p.keys = insertKey(p.keys, i, key)
}

func (p *IndexMinPQ) Contains(i int) bool {
	return p.qp[i] != -1
}

func (p *IndexMinPQ) MinKey() (int, error) {
	if p.size > 0 {
		return p.keys[p.pq[1]], nil
	}
	return -1, ErrNoMinKeyFound
}

func (p *IndexMinPQ) DelMin() (int, error) {
	if p.size > 0 {
		i := p.pq[1]
		p.qp[i] = -1
		p.pq[1] = p.pq[p.size]
		p.pq[p.size] = 0
		p.sink(1)
		p.size--
		return i, nil
	}
	return -1, ErrNoMinKeyFound
}

func (p *IndexMinPQ) swim(k int) {
	// continue to swim when k != 1 i.e. haven't reach root
	if k != 1 {
		// swap with current node's parents if current node
		// value is greater than parent's and continue to
		// swim, otherwise we have restored the heap order
		//fmt.Printf("p.keys[%d] == %d\n", p.pq[k], p.keys[p.pq[k]])
		//fmt.Printf("p.keys[%d] == %d\n", p.pq[k/2], p.keys[p.pq[k/2]])
		if p.keys[p.pq[k]] < p.keys[p.pq[k/2]] {
			p.pq[k], p.pq[k/2] = p.pq[k/2], p.pq[k]
			p.swim(k / 2)
		}
	}
}

func (p *IndexMinPQ) sink(k int) {
	j := 2 * k
	if j < p.size {
		if j + 1 <= p.size {
			if p.keys[p.pq[j-1]] < p.keys[p.pq[j]] {
				j++
			}
		}
		if p.keys[p.pq[j]] < p.keys[p.pq[k]] {
			p.pq[k], p.pq[j] = p.pq[j], p.pq[k]
			p.sink(j)
		}
	}
}

func insertInverse(a []int, i int, n int) []int {
	if i >= cap(a) {
		tmp := make([]int, 2*cap(a))
		copy(tmp[0:len(a)], a)
		for t := len(a); t < len(tmp); t++ {
			tmp[t] = -1
		}
		a = tmp
	}
	a[i] = n
	return a
}

func insertKey(a []int, i int, k int) []int {
	if i >= cap(a) {
		tmp := make([]int, 2*cap(a))
		copy(tmp[0:len(a)], a)
		a = tmp
	}
	a[i] = k
	return a
}
