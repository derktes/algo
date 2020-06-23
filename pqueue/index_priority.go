package pqueue

//import "fmt"

// IndexMinPQ is a variant of minimum-oriented
// priority queue that is indexed
type IndexMinPQ struct {
	pq   []int
	qp   []int
	keys []int
	size int
}

// Insert puts a new key at index i
func (p *IndexMinPQ) Insert(i int, key int) {
	p.size++
	p.pq = append(p.pq, i)
	p.qp = insertInverse(p.qp, i, p.size)
	//fmt.Printf("p.qp %v\n", p.qp)
	p.keys = insertKey(p.keys, i, key)
	p.swim(p.size)
}

func (p *IndexMinPQ) swim(k int) {
	// continue to swim when k != 1 i.e. haven't reach root
	if k != 1 {
		// swap with current node's parents if current node
		// value is greater than parent's and continue to
		// swim, otherwise we have restored the heap order
		//fmt.Printf("p.keys[%d] == %d\n", p.pq[k], p.keys[p.pq[k]])
		//fmt.Printf("p.keys[%d] == %d\n", p.pq[k/2], p.keys[p.pq[k/2]])
		if p.keys[p.pq[k]] > p.keys[p.pq[k/2]] {
			p.pq[k], p.pq[k/2] = p.pq[k/2], p.pq[k]
			p.swim(k / 2)
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
