package test

import (
	"github.com/libbasic/common"
	"github.com/libbasic/datastruct"
	"testing"
)


func TestBst_Add(t *testing.T)  {
	b := datastruct.NewBst(common.CmpInt)
	AssertNotNil(t, b)

	r, err := b.Add(nil)
	AssertFalse(t, r)
	AssertEqual(t, err, "invalid param")

	r, err = b.Add(1)
	AssertTrue(t, r)
	AssertNil(t, err)

	r = b.IsExist(1)
	AssertTrue(t, r)

	r = b.IsExist(4)
	AssertFalse(t, r)

	r, err = b.Add(2)
	AssertTrue(t, r)
	AssertNil(t, err)

	r, err = b.Add(3)
	AssertTrue(t, r)
	AssertNil(t, err)

	r = b.IsExist(3)
	AssertTrue(t, r)

	b.Delete(3)
	r = b.IsExist(3)
	AssertFalse(t, r)
}
