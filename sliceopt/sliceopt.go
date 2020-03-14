package sliceopt

import (
	"fmt"
	"reflect"
)

type compare func(a, b interface{}) int

func checkParam(listA interface{}, listB interface{}, cmp compare) error {
	if nil == cmp {
		return fmt.Errorf("compare should not be nil")
	}
	if reflect.TypeOf(listA).Kind() != reflect.Array {
		return fmt.Errorf("listA should be array type")
	}
	if reflect.TypeOf(listB).Kind() != reflect.Array {
		return fmt.Errorf("listB should be array type")
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
