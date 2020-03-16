package test

import (
	"github.com/libbasic/datastruct"
	"reflect"
	"testing"
)

func intGetOffset(key interface{}, hashMask int) int {
	if reflect.TypeOf(key).Kind() != reflect.Int {
		return datastruct.HashTableInvalidOffset
	}
	k := key.(int)
	return k % hashMask
}

type hashTableTest struct {
	key    interface{}
	data   interface{}
	result interface{}
	err    interface{}
}

var tests_Add = []hashTableTest{
	{nil, nil, false, "invalid param"},
	{"123", nil, false, "invalid param: key(type: string)"},
	{1, 1, true, nil},
	{2, 2, true, nil},
	{1024, 1024, true, nil},
	{1025, 1025, true, nil},
	{1025, 1025, false, "already existed"},
}

var tests_Search = []hashTableTest{
	{nil, nil, nil, "invalid param"},
	{"1", nil, nil, "invalid param: key(type: string)"},
	{1, nil, 1, nil},
	{2, nil, 2, nil},
	{1025, nil, 1025, nil},
	{10, nil, nil, "not existed"},
}

var tests_Delete = []hashTableTest{
	{nil, nil, false, "invalid param"},
	{"12", nil, false, "invalid param: key(type: string)"},
	{1, nil, true, nil},
	{10, nil, true, nil},
}

func TestHashTable_Add(t *testing.T)  {
	h := datastruct.NewHashTable(1024, intGetOffset)
	AssertNotNil(t, h)

	for i := 0; i < len(tests_Add); i++ {
		r, err := h.Add(tests_Add[i].key, tests_Add[i].data)
		if r != tests_Add[i].result || ( nil != err && err.Error() != tests_Add[i].err) {
			t.Errorf("Hash Table Add: key(%v) data(%v) r(%v) err(%v)", tests_Add[i].key, tests_Add[i].data, r,err)
		}
	}


	for i := 0; i < len(tests_Search); i++ {
		r, err := h.Search(tests_Search[i].key)
		if r != tests_Search[i].result || ( nil != err && err.Error() != tests_Search[i].err) {
			t.Errorf("Hash Table Search: key(%v) r(%v) err(%v)", tests_Search[i].key, r, err)
		}
	}

	for i := 0; i < len(tests_Delete); i++ {
		r, err := h.Delete(tests_Delete[i].key)
		if r != tests_Delete[i].result || ( nil != err && err.Error() != tests_Delete[i].err) {
			t.Errorf("Hash Table Delete: key(%v)", tests_Delete[i].key)
		}
	}


	/* 增常添加节点，并查找，删除 */
	r, err := h.Add(111, 111)
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err := h.Search(111)
	AssertEqual(t, data, 111)
	AssertNil(t, err)

	r, err = h.Delete(111)
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err = h.Search(111)
	//t.Errorf("data(%v) err(%v)", ndata, nerr)
	AssertNil(t, data)
	AssertEqual(t, err.Error(), "not existed")


}
