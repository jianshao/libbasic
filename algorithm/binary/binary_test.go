package binary

import (
	"github.com/libbasic/test"
	"testing"
)

type IntSearch []int

func (i *IntSearch)Compare(a, b interface{}) int {
	v1 := a.(int)
	v2 := b.(int)
	if v1 > v2 {
		return 1
	} else if v1 == v2 {
		return 0
	} else {
		return -1
	}
}

func (i IntSearch)Len() int {
	return len(i)
}

func (i IntSearch)GetByIndex(pos int) interface{} {
	return i[pos]
}

func (i *IntSearch)Add(data int)  {
	*i = append(*i, data)
}

func TestIntSearch(t *testing.T)  {
	var i = new(IntSearch)
	i.Add(1)
	i.Add(2)
	i.Add(3)
	i.Add(4)
	i.Add(10)
	i.Add(15)

	pos := Search(i, 10)
	test.AssertEqual(t, pos, 4)
	pos = Search(i, 1)
	test.AssertEqual(t, pos, 0)
	pos = Search(i, 20)
	test.AssertEqual(t, pos, -1)
}



type item struct {
	value int
}

type structSearch []item

func (s structSearch)Len() int {
	return len(s)
}

func (s structSearch)Compare(a, b interface{}) int {
	v1 := a.(item)
	v2 := b.(item)
	if v1.value > v2.value {
		return 1
	} else if v1.value == v2.value {
		return 0
	} else {
		return -1
	}
}

func (s structSearch)GetByIndex(pos int) interface{} {
	return s[pos]
}

func (s *structSearch)Add(i item)  {
	*s = append(*s, i)
}

func TestStructSearch(t *testing.T)  {
	s := new(structSearch)
	s.Add(item{value:1})
	s.Add(item{value:2})
	s.Add(item{value:10})
	s.Add(item{value:18})

	pos := Search(s, item{value:1})
	test.AssertEqual(t, pos, 0)
	pos = Search(s, item{value:10})
	test.AssertEqual(t, pos, 2)
	pos = Search(s, item{value:20})
	test.AssertEqual(t, pos, -1)
}