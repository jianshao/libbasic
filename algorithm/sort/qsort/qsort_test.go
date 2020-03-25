package qsort

import (
	"testing"
)

type SortInt []int

func (s SortInt)Len() int {
	return len(s)
}

func (s SortInt)Swap(i, j int)  {

	s[i], s[j] = s[j], s[i]
}

func (s SortInt)Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	} else {
		return false
	}
}

func TestQsort(t *testing.T) {
	t.Helper()
	var list = []int{10, 1,3,9,4,5}
	value := make(SortInt, len(list))
	value = append(list, 8)
	for k, v := range value {
		t.Error("k: ",k, " v: ", v)
	}
	Qsort(value, 0, len(value) - 1)

	for k, v := range value {
		t.Error("k: ",k, " v: ", v)
	}
}
