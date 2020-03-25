package bst

import (
	"github.com/libbasic/common"
	"github.com/libbasic/test"
	"testing"
)


func TestBst_Add(t *testing.T)  {
	b := NewBst(common.CmpInt)
	test.AssertNotNil(t, b)

	r, err := b.Add(nil)
	test.AssertFalse(t, r)
	test.AssertEqual(t, err.Error(), "invalid param")

	r, err = b.Add(1)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	r = b.IsExist(1)
	test.AssertTrue(t, r)

	r = b.IsExist(4)
	test.AssertFalse(t, r)

	r, err = b.Add(2)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	r, err = b.Add(3)
	test.AssertTrue(t, r)
	test.AssertNil(t, err)

	r = b.IsExist(3)
	test.AssertTrue(t, r)

	b.Delete(3)
	r = b.IsExist(3)
	test.AssertFalse(t, r)
}
