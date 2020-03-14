package sliceopt

import (
	"github.com/libbasic/common"
)



func StringSliceAnd(listA, listB []string) ([]string, error) {
	r, err := SliceAnd(listA, listB, common.CmpString)
	if err != nil {
		return nil, err
	}

	result := common.ConvertInterface2String(r)
	return result, nil
}

func StringSliceOr(listA, listB []string) ([]string, error) {
	r, err := SliceOr(listA, listB, common.CmpString)
	if err != nil {
		return nil, err
	}

	result := common.ConvertInterface2String(r)
	return result, nil
}

func StringSliceSub(listA, listB []string) ([]string, error) {
	r, err := SliceSub(listA, listB, common.CmpString)
	if err != nil {
		return nil, err
	}

	result := common.ConvertInterface2String(r)
	return result, nil
}