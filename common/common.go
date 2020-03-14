package common

import "reflect"

type Compare func(a, b interface{}) int

func CmpInt(a, b interface{}) int {
	valueA := reflect.ValueOf(a).Int()
	valueB := reflect.ValueOf(b).Int()
	if valueA > valueB {
		return 1
	} else if valueA == valueB {
		return 0
	} else {
		return -1
	}
}

func CmpString(a, b interface{}) int {
	valueA := reflect.ValueOf(a).String()
	valueB := reflect.ValueOf(b).String()
	if valueA > valueB {
		return 1
	} else if valueA == valueB {
		return 0
	} else {
		return -1
	}
}

func ConvertInterface2String(data []interface{}) []string {
	result := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i].(string)
	}
	return result
}

func ConvertInterface2Int(data []interface{}) []int {
	result := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i].(int)
	}
	return result
}