package sliceopt

import (
	"fmt"
	"github.com/libbasic/common"
	"reflect"
)

const (
	SLICE_POS_INVALID = -1
)

type compare func(a, b interface{}) int

func checkParam(listA interface{}, listB interface{}, cmp compare) error {
	if nil == cmp {
		return fmt.Errorf("compare should not be nil")
	}
	if _, err := common.IsValidInterface(listA, reflect.Array, "listA"); err != nil {
		return err
	}
	if _, err := common.IsValidInterface(listB, reflect.Array, "listB"); err != nil {
		return err
	}
	if reflect.ValueOf(listA).Index(0).Kind() != reflect.ValueOf(listB).Index(0).Kind() {
		return fmt.Errorf("type of element in listA is not equal with it in listB")
	}

	return nil
}

func SliceAnd(listA interface{}, listB interface{}, cmp compare) ([]interface{}, error) {

	if err := checkParam(listA, listB, cmp); err != nil {
		return nil, err
	}

	lenA := reflect.ValueOf(listA).Len()
	lenB := reflect.ValueOf(listB).Len()
	result := make([]interface{}, 0, lenA)

	valueA := reflect.ValueOf(listA)
	valueB := reflect.ValueOf(listB)
	for i, j := 0, 0; i < lenA && j < lenB; {
		for ; i < lenA; {
			r := 1
			if j < lenB {
				r = cmp(valueA.Index(i), valueB.Index(j))
			}

			if r == 0 {
				result = append(result, valueA.Index(i))
				i++
				j++
			} else if r > 0 {
				break
			} else {
				i++
			}
		}

		for ; j < lenB; {
			r := -1
			if i < lenA {
				r = cmp(valueA.Index(i), valueB.Index(j))
			}

			if r == 0 {
				result = append(result, valueB.Index(j))
				j++
				i++
			} else if r < 0 {
					break
			} else {
				j++
			}
		}
	}

	return result, nil
}

func SliceOr(listA interface{}, listB interface{}, cmp compare) ([]interface{}, error) {
	if err := checkParam(listA, listB, cmp); err != nil {
		return nil, err
	}

	lenA := reflect.ValueOf(listA).Len()
	lenB := reflect.ValueOf(listB).Len()
	result := make([]interface{}, 0, lenA + lenB)

	valueA := reflect.ValueOf(listA)
	valueB := reflect.ValueOf(listB)
	for i, j := 0, 0; i < lenA || j < lenB; {
		for ; i < lenA; {
			/* 如果listB已经全部遍历过了，只需要把listA剩余的节点加入结果集即可 */
			r := -1
			if j < lenB {
				r = cmp(valueA.Index(i), valueB.Index(j))
			}

			if r == 0 {
				result = append(result, valueA.Index(i))
				i++
				j++
			} else if r > 0 {
					break
			} else {
				result = append(result, valueA.Index(i))
				i++
			}
		}
		for ; j < lenB; {
			r := 1
			if i < lenA {
				r = cmp(valueA.Index(i), valueB.Index(j))
			}

			if r == 0 {
				result = append(result, valueA.Index(i))
				i++
				j++
			} else if r < 0 {
					break
			} else {
				result = append(result, valueB.Index(j))
				j++
			}
		}
	}

	return result, nil
}

func SliceSub(listA interface{}, listB interface{}, cmp compare) ([]interface{}, error) {
	if err := checkParam(listA, listB, cmp); err != nil {
		return nil, err
	}

	lenA := reflect.ValueOf(listA).Len()
	lenB := reflect.ValueOf(listB).Len()
	result := make([]interface{}, 0, lenA)

	valueA := reflect.ValueOf(listA)
	valueB := reflect.ValueOf(listB)
	for i, j := 0, 0; i < lenA; {
		r := -1
		if j < lenB {
			r = cmp(valueA.Index(i), valueB.Index(j))
		}
		if r == 0 {
			i++
			j++
		} else if r > 0 {
				j++
		} else {
					result = append(result, valueA.Index(i))
					i++
		}
	}

	return result, nil
}

/*
  说明：在有序的切片中检索，可以是升序或降序。如果有多个重复的元素，则返回的元素位置是不确定的。
  输入：
      s：必须是数组类型，数组元素类型必须与v的类型保持一致。
      v：需要检索的数据
  start：数组中查找的起始位置，必须在切片中可用，否则会panic
    end：数组中查找的结束位置，必须在切片中可用，否则会panic
    cmp：元素间比较函数，返回值必须与升降序保持一致
  输出：
     1.命中的元素下标。未命中时返回不可用值 SLICE_POS_INVALID
     2.未命中时下一个元素的下标，使用前必须先判断是否是可用值。
     3.错误信息
*/
func SearchInSortedSlice(s interface{}, v interface{}, start int, end int, cmp compare) (int, int, error) {
	if _, err := common.IsValidInterface(s, reflect.Array, "s"); err != nil {
		return SLICE_POS_INVALID, SLICE_POS_INVALID, err
	}
	if _, err := common.IsValidInterface(v, reflect.ValueOf(s).Index(start).Kind(), "v"); err != nil {
		return SLICE_POS_INVALID, SLICE_POS_INVALID, err
	}

	if start < 0 || start > end {
		return SLICE_POS_INVALID, SLICE_POS_INVALID, fmt.Errorf("invalid param: start(%d)", start)
	}

	if end < 0 {
		return SLICE_POS_INVALID, SLICE_POS_INVALID, fmt.Errorf("invalid param: end(%d)", end)
	}

	SliceValue := reflect.ValueOf(s)
	r := cmp(SliceValue.Index(start), reflect.ValueOf(v))
	if r == 0 {
		return start, SLICE_POS_INVALID, nil
	} else if r > 0 {
		return SLICE_POS_INVALID, SLICE_POS_INVALID, fmt.Errorf("not existed")
	}

	r = cmp(SliceValue.Index(end), reflect.ValueOf(v))
	if r == 0 {
		return end, SLICE_POS_INVALID, nil
	} else if r < 0 {
		return SLICE_POS_INVALID, SLICE_POS_INVALID, fmt.Errorf("not existed")
	}

	for ; start < end; {
		mid := (start + end) / 2
		r := cmp(SliceValue.Index(mid), reflect.ValueOf(v))
		if r == 0 {
			return mid, SLICE_POS_INVALID, nil
		} else if r > 0 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return SLICE_POS_INVALID, end, fmt.Errorf("not existed")
}