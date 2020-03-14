package algorithm

import (
	"fmt"
	"github.com/libbasic/common"
	"reflect"
)


func sort(newData interface{}, start, end int, cmp common.Compare)  {

	old := start
	temp := reflect.ValueOf(newData).Index(old)
	for i, j := start + 1, end; i < j; {
		for ;;j--{
			if cmp(temp, reflect.ValueOf(newData).Index(j)) > 0 {
				reflect.ValueOf(newData).Index(old).Set(temp)
				old = j
				break
			}
		}
		for ;;i++ {
			if cmp(temp, reflect.ValueOf(newData).Index(i)) < 0 {
				reflect.ValueOf(newData).Index(old).Set(temp)
				old = i
				break
			}
		}
	}
	reflect.ValueOf(newData).Index(old).Set(temp)
	sort(newData, start, old - 1, cmp)
	sort(newData, old + 1, end, cmp)
}

func Qsort(data interface{}, start, end int, cmp common.Compare) (interface{}, error) {

	if _, err := common.IsValidInterface(data, reflect.Array, "data"); err != nil {
		return nil, err
	}
	if start < 0 {
		return nil, fmt.Errorf("invalid param: start should not be less than 0")
	}
	if end < 0 {
		return  nil, fmt.Errorf("invalid param: end should not be less than 0")
	}
	if start > end {
		return nil, fmt.Errorf("invalid param: end should be bigger than start")
	}
	if reflect.ValueOf(data).Len() <= end {
		return nil, fmt.Errorf("invalid param: end should be less length of data")
	}
	if cmp == nil {
		return nil, fmt.Errorf("invalid param: cmp should not be nil")
	}

	sort(data, start, end, cmp)
	return data, nil
}

func QsortInt(data []int, start, end int) ([]int, error) {
	r, err := Qsort(data, start, end, common.CmpInt)
	if err != nil {
		return nil, err
	}

	return r.([]int), nil
}

func QsortString(data []string, start, end int) ([]string, error) {
	r, err := Qsort(data, start, end, common.CmpString)
	if err != nil {
		return nil, err
	}
	return r.([]string), nil
}