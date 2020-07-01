package symtable

import (
	"container/list"
	"errors"
	"sort"
)

var (
	ErrValueNotFound = errors.New("Value not found")
	ErrNoKeys        = errors.New("No keys found")
)

type keyValue struct {
	key   byte
	value int
}

type SequentialSearchST struct {
	ls *list.List
}

func New() *SequentialSearchST {
	return &SequentialSearchST{list.New()}
}

func (st *SequentialSearchST) Get(key byte) (int, error) {
	for e := st.ls.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*keyValue)
		if kv.key == key {
			return kv.value, nil
		}
	}
	return 0, ErrValueNotFound
}

func (st *SequentialSearchST) Put(key byte, val int) {
	for e := st.ls.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*keyValue)
		if kv.key == key {
			kv.value = val
			return
		}
	}
	st.ls.PushFront(&keyValue{key, val})
}

func (st *SequentialSearchST) Keys() ([]byte, error) {
	size := st.ls.Len()
	ikeys := make([]int, 0, size)
	for e := st.ls.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*keyValue)
		ikeys = append(ikeys, int(kv.key))
	}
	if len(ikeys) > 0 {
		islice := sort.IntSlice(ikeys)
		islice.Sort()
		keys := make([]byte, 0, size)
		for _, k := range ikeys {
			keys = append(keys, byte(k))
		}
		return keys, nil
	}
	return nil, ErrNoKeys
}
