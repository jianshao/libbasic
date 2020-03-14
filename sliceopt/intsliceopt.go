package sliceopt

import (
	"github.com/libbasic/common"
)


func IntSliceAnd(listA, listB []int) ([]int, error) {
	r, err := SliceAnd(listA, listB, common.CmpInt)
	if err != nil {
		return nil, err
	}

	result := common.ConvertInterface2Int(r)
	return result, nil
}

func IntSliceOr(listA, listB []int) ([]int, error) {
	r, err := SliceOr(listA, listB, common.CmpInt)
	if err != nil {
		return nil, err
	}

	result := common.ConvertInterface2Int(r)
	return result, nil
}

func IntSliceSub(listA, listB []int) ([]int, error) {
	r, err := SliceSub(listA, listB, common.CmpInt)
	if err != nil {
		return nil, err
	}

	result := common.ConvertInterface2Int(r)
	return result, nil
}