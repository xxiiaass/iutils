package iutils

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

// 类型转换相关

// 判断该值是否为表示空值，0，"0"， false, "false"等
func Empty(v interface{}) bool {
	switch v.(type) {
	case uint32:
		return v.(uint32) == 0
	case uint64:
		return v.(uint64) == 0
	case int:
		return v.(int) == 0
	case int32:
		return v.(int32) == 0
	case int64:
		return v.(int64) == int64(0)
	case float64:
		return v.(float64) == float64(0)
	case float32:
		return v.(float32) == float32(0)
	case string:
		return v.(string) == "" || v.(string) == "0" || v.(string) == "false"
	case bool:
		return v.(bool) == false
	default:
		return false
	}
}

// 判断一堆基础类型的值是不是都是空
func AllEmpty(vl ...interface{}) bool {
	for _, v := range vl {
		if !Empty(v) {
			return false
		}
	}
	return true
}

// 判断interface是否为nil
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if exists, _ := InArray(vi.Kind(), []reflect.Kind{reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice}); exists {
		return vi.IsNil()
	}

	return i == nil
}

func ValNotNilOrDefault(v, d interface{}) interface{} {
	if v != nil {
		return v
	}
	return d
}

// 如果是字符串，则转化为int64, 浮点数丢失精度，无法转化则为0
func AsInt64(v interface{}) int64 {
	switch v.(type) {
	case uint32:
		return int64(v.(uint32))
	case uint64:
		return int64(v.(uint64))
	case int:
		return int64(v.(int))
	case int32:
		return int64(v.(int32))
	case int64:
		return v.(int64)
	case float64:
		return int64(v.(float64))
	case float32:
		return int64(v.(float32))
	case json.Number:
		vv, err := strconv.ParseInt(string(v.(json.Number)), 10, 64)
		if err != nil {
			return 0
		}
		return vv
	case string:
		vv, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			return 0
		}
		return vv
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func AsBool(v interface{}) bool {
	switch nv := v.(type) {
	case uint32:
		return nv != 0
	case uint64:
		return nv != 0
	case int:
		return nv != 0
	case int32:
		return nv != 0
	case int64:
		return nv != 0
	case float64:
		return nv != 0
	case float32:
		return nv != 0
	case json.Number:
		vv, err := strconv.ParseInt(string(v.(json.Number)), 10, 64)
		if err != nil {
			return false
		}
		return vv != 0
	case string:
		vv, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			return false
		}
		return vv != 0
	case bool:
		return nv
	default:
		return false
	}
}

func AsInt64Slice(v interface{}) []int64 {
	slice := reflect.ValueOf(v)
	if slice.Kind() != reflect.Slice {
		return []int64{}
	}
	ret := make([]int64, 0)
	if slice.Kind() == reflect.Ptr {
		slice = slice.Elem()
	}
	for i := 0; i < slice.Len(); i++ {
		ret = append(ret, AsInt64(slice.Index(i).Addr().Elem().Interface()))
	}
	return ret
}

func AsInt64OrDefault(v interface{}, d int64) int64 {
	switch v.(type) {
	case uint32:
		return int64(v.(uint32))
	case uint64:
		return int64(v.(uint64))
	case int:
		return int64(v.(int))
	case int32:
		return int64(v.(int32))
	case int64:
		return v.(int64)
	case float64:
		return int64(v.(float64))
	case float32:
		return int64(v.(float32))
	case string:
		vv, err := strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			return 0
		}
		return vv
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return d
	}
}

func AsInt(v interface{}) int {
	switch v.(type) {
	case uint32:
		return int(v.(uint32))
	case uint64:
		return int(v.(uint64))
	case int:
		return v.(int)
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case float64:
		return int(v.(float64))
	case float32:
		return int(v.(float32))
	case string:
		vv, err := strconv.Atoi(v.(string))
		if err != nil {
			return 0
		}
		return vv
	case json.Number:
		vv, err := strconv.Atoi(string(v.(json.Number)))
		if err != nil {
			return 0
		}
		return vv
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func AsFloat64(v interface{}) float64 {
	switch v.(type) {
	case uint32:
		return float64(v.(uint32))
	case uint64:
		return float64(v.(uint64))
	case int:
		return float64(v.(int))
	case int32:
		return float64(v.(int32))
	case int64:
		return float64(v.(int64))
	case float64:
		return v.(float64)
	case float32:
		return float64(v.(float32))
	case string:
		vv, err := strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return 0
		}
		return vv
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func AsString(v interface{}) string {
	switch v.(type) {
	case uint32:
		return strconv.FormatInt(int64(v.(uint32)), 10)
	case uint64:
		return strconv.FormatInt(int64(v.(uint64)), 10)
	case int:
		return strconv.Itoa(v.(int))
	case int32:
		return strconv.Itoa(int(v.(int32)))
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case float64:
		return strconv.FormatFloat(v.(float64), 'E', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'E', -1, 64)
	case string:
		return v.(string)
	case bool:
		if v.(bool) {
			return "true"
		} else {
			return "false"
		}
	case map[string]interface{}:
		return JsonEncode(v)
	default:
		return ""
	}
}

func AsStringOrDefault(v interface{}, d string) string {
	switch v.(type) {
	case uint32:
		return strconv.FormatInt(int64(v.(uint32)), 10)
	case uint64:
		return strconv.FormatInt(int64(v.(uint64)), 10)
	case int:
		return strconv.Itoa(v.(int))
	case int32:
		return strconv.Itoa(int(v.(int32)))
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case float64:
		return strconv.FormatFloat(v.(float64), 'E', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v.(float32)), 'E', -1, 64)
	case string:
		return v.(string)
	case bool:
		if v.(bool) {
			return "true"
		} else {
			return "false"
		}
	default:
		return d
	}
}

func AsH(v interface{}) H {
	switch v.(type) {
	case H:
		return v.(H)
	}
	return nil
}
func AsStringSlice(v interface{}) []string {
	slice := reflect.ValueOf(v)
	if !slice.IsValid() {
		return make([]string, 0)
	}
	ret := make([]string, 0)
	if slice.Kind() == reflect.Ptr {
		slice = slice.Elem()
	}
	for i := 0; i < slice.Len(); i++ {
		ret = append(ret, AsString(slice.Index(i).Addr().Elem().Interface()))
	}
	return ret
}



func SliceString2Int(src []string) ([]int, error) {
	dst := make([]int, 0, len(src))
	for i := range src {
		n, err := strconv.Atoi(src[i])
		if err != nil {
			return nil, err
		}
		dst = append(dst, n)
	}
	return dst, nil
}

// 字符串切片转换为 int64 切片
func SliceString2Int64(src []string) ([]int64, error) {
	tmp, err := SliceString2Int(src)
	if err != nil {
		return nil, err
	}

	dst := make([]int64, 0, len(tmp))
	for _, i := range tmp {
		dst = append(dst, int64(i))
	}
	return dst, nil
}


// "s1, s2, s3, s4" => []string{"s1", "s2", "s3", "s4"}
// "[s1, s2, s3, s4]" => []string{"s1", "s2", "s3", "s4"}
// "["s1", "s2", "s3", "s4"]" => []string{"s1", "s2", "s3", "s4"}
func StrToStrSlice(str string) []string {
	if str == "" {
		return nil
	}
	if str[0] == '[' && str[len(str)-1] == ']' {
		str = str[1 : len(str)-1]
	}
	strs := strings.Split(str, ",")
	ids := make([]string, len(strs))
	for i, s := range strs {
		ids[i] = strings.Trim(strings.Trim(s, " "), "\"")
	}
	return ids
}

func SliceToStr(arr []int64) string {
	builder := strings.Builder{}
	for i, num := range arr {
		builder.WriteString(strconv.Itoa(int(num)))
		if i != len(arr)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

func StrSliceToStr(arr []string) string {
	builder := strings.Builder{}
	for _, str := range arr {
		if str == "" {
			continue
		}
		builder.WriteString(str)
		builder.WriteString(",")
	}

	res := builder.String()
	if res != "" && res[len(res)-1] == ',' {
		res = res[:len(res)-1]
	}
	return res
}

// "1,2,3,4" => []int{1, 2, 3, 4}
// "[1,2,3,4]" => []int{1, 2, 3, 4}
func StrToSlice(str string) []int64 {
	if str == "" {
		return nil
	}
	if str[0] == '[' && str[len(str)-1] == ']' {
		str = str[1 : len(str)-1]
	}
	strs := strings.Split(str, ",")
	ids := make([]int64, len(strs))
	for i, s := range strs {
		ids[i] = AsInt64(s)
	}
	return ids
}

