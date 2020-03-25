package qsort

import (
	"github.com/libbasic/common"
	"reflect"
)

type QSort interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

func Qsort(q QSort, start, end int)  {
	if start > end {
		return
	}
	old := start
	for i, j := start + 1, end; i <= j; {
		for ;i <= j; j--{
			if q.Less(j, old) {
				q.Swap(old, j)
				old = j
				break
			}
		}
		for ; i <= j; i++ {
			if q.Less(old, i) {
				q.Swap(old, i)
				old = i
				break
			}
		}
	}
	Qsort(q, start, old - 1)
	Qsort(q, old + 1, end)
}

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

/*
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
}*/