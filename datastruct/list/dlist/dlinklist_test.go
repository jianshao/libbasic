package dlist

import (
	"github.com/libbasic/test"
	"testing"
)

func TestDLinkList_New(t *testing.T)  {
	d := NewDLinkList()
	test.AssertNotNil(t, d)

	data, err := Get(1)
	test.AssertNil(t, data)
	test.AssertEqual(t, err.Error(), "not existed")

	data, err = GetHead()
	test.AssertNil(t, data)
	test.AssertEqual(t, err.Error(), "not existed")

	data, err = GetTail()
	test.AssertNil(t, data)
	test.AssertEqual(t, err.Error(), "not existed")

	c := GetNodeCount()
	test.AssertEqual(t, c, 0)
}

func TestDLinkListAdd(t *testing.T)  {
	d := NewDLinkList()
	test.AssertNotNil(t, d)

	r, err := Add2Head(nil)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "invalid param")

	r, err = Add2Head(1)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	data, err := GetHead()
	test.AssertNotNil(t, data)
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)

	data, err = GetTail()
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)

	r, err = Add2Head(2)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	data, err = GetHead()
	test.AssertEqual(t, data, 2)
	test.AssertNil(t, err)

	data, err = Get(1)
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)

	data, err = GetTail()
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)

	r, err = DeleteHead()
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	data, err = GetHead()
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)

	r, err = Add2Tail(3)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	data, err = GetTail()
	test.AssertEqual(t, data, 3)
	test.AssertNil(t, err)

	r, err = DeleteTail()
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	data, err = GetTail()
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)
	
}

func TestDLinkList_Get(t *testing.T)  {
	d := NewDLinkList()
	test.AssertNotNil(t, d)

	r, err := Append(1)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	r, err = Append(2)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	r, err = Append(3)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	r, err = Append(4)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	r, err = Append(5)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	data, err := Get(1)
	test.AssertEqual(t, data, 1)
	test.AssertNil(t, err)

	r, err = Delete(1)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	data, err = Get(1)
	test.AssertNil(t, data)
	test.AssertEqual(t, err.Error(), "not existed")
}