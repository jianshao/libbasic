package test

import (
	"github.com/libbasic/datastruct"
	"testing"
)

func TestDLinkList_New(t *testing.T)  {
	d := datastruct.NewDLinkList()
	AssertNotNil(t, d)

	data, err := d.Get(1)
	AssertNil(t, data)
	AssertEqual(t, err, "not existed")

	data, err = d.GetHead()
	AssertNil(t, data)
	AssertEqual(t, err, "not existed")

	data, err = d.GetTail()
	AssertNil(t, data)
	AssertEqual(t, err, "not existed")

	c := d.GetNodeCount()
	AssertEqual(t, c, 0)
}

func TestDLinkListAdd(t *testing.T)  {
	d := datastruct.NewDLinkList()
	AssertNotNil(t, d)

	r, err := d.Add2Head(nil)
	AssertFalse(t, r)
	AssertEqual(t, err, "invalid param")

	r, err = d.Add2Head(1)
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err := d.GetHead()
	AssertNotNil(t, data)
	AssertEqual(t, data, 1)
	AssertNil(t, err)

	data, err = d.GetTail()
	AssertEqual(t, data, 1)
	AssertNil(t, err)

	r, err = d.Add2Head(2)
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err = d.GetHead()
	AssertEqual(t, data, 2)
	AssertNil(t, err)

	data, err = d.Get(1)
	AssertEqual(t, data, 1)
	AssertNil(t, err)

	data, err = d.GetTail()
	AssertEqual(t, data, 1)
	AssertNil(t, err)

	r, err = d.DeleteHead()
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err = d.GetHead()
	AssertEqual(t, data, 1)
	AssertNil(t, err)

	r, err = d.Add2Tail(3)
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err = d.GetTail()
	AssertEqual(t, data, 3)
	AssertNil(t, err)

	r, err = d.DeleteTail()
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err = d.GetTail()
	AssertEqual(t, data, 1)
	AssertNil(t, err)
	
}

func TestDLinkList_Get(t *testing.T)  {
	d := datastruct.NewDLinkList()
	AssertNotNil(t, d)

	r, err := d.Append(1)
	AssertTrue(t, r)
	AssertNil(t, err)

	r, err = d.Append(2)
	AssertTrue(t, r)
	AssertNil(t, err)

	r, err = d.Append(3)
	AssertTrue(t, r)
	AssertNil(t, err)

	r, err = d.Append(4)
	AssertTrue(t, r)
	AssertNil(t, err)

	r, err = d.Append(5)
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err := d.Get(1)
	AssertEqual(t, data, 1)
	AssertNil(t, err)

	r, err = d.Delete(1)
	AssertTrue(t, r)
	AssertNil(t, err)

	data, err = d.Get(1)
	AssertNil(t, data)
	AssertEqual(t, err, "not existed")
}