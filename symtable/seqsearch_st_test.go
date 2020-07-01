package symtable

import (
	"testing"
)

type testCase struct {
	in   []keyValue
	want []keyValue
}

func setupAndTest(t *testing.T, cases []testCase) {
	for _, c := range cases {
		st := New()
		// insert each key value pair
		for _, kv := range c.in {
			st.Put(kv.key, kv.value)
		}
		// retrieve each all values with each associated key
		for i, kv := range c.want {
			v, err := st.Get(kv.key)
			// fail the test when hit error or value differs
			if err != nil || v != kv.value {
				t.Logf("After multiple Put, list is not %v", c.want)
				t.Logf("Error encountered is '%s'", err)
				t.Fatalf("Value for key %v @ %d is %d", kv.key, i, v)
			}
		}
	}
}

func TestPutMultiple(t *testing.T) {
	cases := []testCase{
		{
			[]keyValue{{'E', 0}, {'A', 1}, {'R', 2}},
			[]keyValue{{'E', 0}, {'A', 1}, {'R', 2}},
		},
	}
	setupAndTest(t, cases)
}

func TestUpdateWithPut(t *testing.T) {
	cases := []testCase{
		{
			[]keyValue{{'B', 3}, {'Z', 4}, {'T', -2}, {'P', 6}, {'T', 5}},
			[]keyValue{{'B', 3}, {'Z', 4}, {'T', 5}, {'P', 6}},
		},
	}
	setupAndTest(t, cases)
}

func TestGetOnEmpty(t *testing.T) {
	st := New()
	_, err := st.Get('A')
	if err == nil {
		t.Log("After Get on empty symbol table, err != nil")
		t.Fatalf("Error encountered is '%s'", err)
	}
}

func TestKeysSorted(t *testing.T) {
	cases := []testCase{
		{
			[]keyValue{{'Y', 0}, {'M', 0}, {'C', 0}, {'A', 0}},
			[]keyValue{{'A', 0}, {'C', 0}, {'M', 0}, {'Y', 0}},
		},
	}
	for _, c := range cases {
		st := New()
		for _, kv := range c.in {
			st.Put(kv.key, kv.value)
		}
		keys, err := st.Keys()
		// assert no error
		if err != nil {
			t.Fatalf("After call Keys, error encountered is '%s'", err)
		}
		// assert equal length
		if len(keys) != len(c.want) {
			t.Fatalf("After call Keys, length is not %v", len(c.want))
		}
		// assert order
		for i, kv := range c.want {
			if kv.key != keys[i] {
				t.Fatalf("After call Keys, key @ %d is %v but want %v", i, keys[i], kv.key)
			}
		}
	}

}
