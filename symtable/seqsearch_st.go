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
	key   string
	value int
}

type SequentialSearchST struct {
	ls *list.List
}

func New() *SequentialSearchST {
	return &SequentialSearchST{list.New()}
}

func (st *SequentialSearchST) Get(key string) (int, error) {
	for e := st.ls.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*keyValue)
		if kv.key == key {
			return kv.value, nil
		}
	}
	return 0, ErrValueNotFound
}

func (st *SequentialSearchST) Put(key string, val int) {
	for e := st.ls.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*keyValue)
		if kv.key == key {
			kv.value = val
			return
		}
	}
	st.ls.PushFront(&keyValue{key, val})
}

func (st *SequentialSearchST) Keys() ([]string, error) {
	size := st.ls.Len()
	keys := make([]string, 0, size)
	for e := st.ls.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*keyValue)
		keys = append(keys, kv.key)
	}
	if len(keys) > 0 {
		sort.StringSlice(keys).Sort()
		return keys, nil
	}
	return nil, ErrNoKeys
}
