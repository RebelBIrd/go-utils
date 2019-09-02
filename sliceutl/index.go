package sliceutl

import (
	"reflect"
	"sort"
	"strings"
)

func InArray(key interface{}, sli []interface{}) bool {
	for _, v := range sli {
		if key == v {
			return true
		}
	}
	return false
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	sort.Strings(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func IsStringInArray(val string, array []string) bool {
	var isIn = false
	for _, v := range array {
		if val == v {
			isIn = true
			break
		} else if strings.Contains(val, v) {
			isIn = true
			break
		}
	}
	return isIn
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
