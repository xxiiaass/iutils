package iutils

import "reflect"

func InArrayByFunc(val interface{}, array interface{}, equalFunc func(val, arrayItem interface{}) bool) (exist bool, index int) {
	exist = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if equalFunc(val, s.Index(i).Interface()) {
				index = i
				exist = true
				return
			}
		}
	}
	return
}

func IsExists(val interface{}, array interface{}) bool {
	e, _ := InArray(val, array)
	return e
}

func IsExistsByFunc(val interface{}, array interface{}, equalFunc func(val, arrayItem interface{}) bool) bool {
	e, _ := InArrayByFunc(val, array, equalFunc)
	return e
}

func IndexOf(val interface{}, array interface{}) int {
	_, i := InArray(val, array)
	return i
}

func IndexOfByFunc(val interface{}, array interface{}, equalFunc func(val, arrayItem interface{}) bool) int {
	_, i := InArrayByFunc(val, array, equalFunc)
	return i
}

func Deduplication(arr []int64) []int64 {
	helper := make(map[int64]bool)
	res := make([]int64, 0)
	for _, item := range arr {
		if _, ok := helper[item]; !ok {
			res = append(res, item)
			helper[item] = true
		}
	}
	return res
}


func RemoveNum(arr []int64, target int64) []int64 {
	index := -1
	for i, v := range arr {
		if v == target {
			index = i
			break
		}
	}
	if index == -1 {
		return arr
	}
	if index == len(arr)-1 {
		return arr[:len(arr)-1]
	}
	return append(arr[:index], arr[index+1:]...)
}



// 删除交集
func RemoveIntersection(arr1, arr2 []int64) []int64 {
	arr2m := make(map[int64]struct{})
	for _, i := range arr2 {
		arr2m[i] = struct{}{}
	}
	result := make([]int64, 0)
	for _, i := range arr1 {
		if _, ok := arr2m[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}


// 切片中的所有元素是否都满足指定方法
func All(array interface{}, function func(item interface{}) bool) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if !function(s.Index(i).Interface()) {
				return false
			}
		}
	default:
		return false
	}
	return true
}

// 切片中是否存在满足指定方法的元素，如果不传方法，则判断切片中是否有元素，传多个方法只会取第一个
func AnyOne(array interface{}, function ...func(item interface{}) bool) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		f := func(_ interface{}) bool { return true }
		if len(function) > 0 {
			f = function[0]
		}
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if f(s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}


// 将切片中的每个元素都作为入参传入指定方法中，收集方法返回值并放入切片返回
func Select(array interface{}, function func(item interface{}) interface{}) []interface{} {
	res := make([]interface{}, 0)
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			res = append(res, function(s.Index(i).Interface()))
		}
	}
	return res
}

// 返回切片中满足指定方法的元素个数
func Count(array interface{}, function func(item interface{}) bool) int64 {
	var res int64
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if function(s.Index(i).Interface()) {
				res++
			}
		}
	}
	return res
}

// 返回切片中满足指定方法的元素
func Where(array interface{}, function func(item interface{}) bool) []interface{} {
	res := make([]interface{}, 0)
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if function(s.Index(i).Interface()) {
				res = append(res, s.Index(i).Interface())
			}
		}
	}
	return res
}

// 返回切片中第一个满足指定方法的元素，如不存在则返回nil
func First(array interface{}, function func(item interface{}) bool) interface{} {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if function(s.Index(i).Interface()) {
				return s.Index(i).Interface()
			}
		}
	}
	return nil
}

// 返回切片中第一个满足指定方法的元素，如不存在则返回入参中的默认值
func FirstOrDefault(array interface{}, function func(item interface{}) bool, def interface{}) interface{} {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if function(s.Index(i).Interface()) {
				return s.Index(i).Interface()
			}
		}
	}
	return def
}

