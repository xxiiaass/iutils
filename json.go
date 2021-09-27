package iutils

import (
	"encoding/json"
	"strings"
)

// 更方便的json数据解析，建议使用map.go文件中的toStruct方法，解析到结构体中


// 编码任意数据为json字符串
func JsonEncode(v Any) string {
	bt, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bt)
}

// 将json字符串解码成map结构 （此时，数据的结构会变成浮点数，超过浮点数精度的数据，会不准确）
func JsonDecode(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &mapResult)
	return mapResult
}

// 将json字符串解码成map结构，能保留大数的数据
func JsonDecodeSafe(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	ds := json.NewDecoder(strings.NewReader(jsonStr))
	ds.UseNumber()
	_ = ds.Decode(&mapResult)
	return mapResult
}

// 将json字符串解码成map结构，也可以是个数组，在两个返回值中返回一个，使用者需要知道数据是什么类型
func JsonDecodes(jsonStr string) (map[string]interface{}, []map[string]interface{}) {
	if jsonStr[0] == '[' {
		mapSlice := make([]map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(jsonStr), &mapSlice)
		return nil, mapSlice
	}
	var mapResult map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &mapResult)
	return mapResult, nil
}

// 可以解析任意json类型的数据，以interface{}类型返回
func JsonDecodeAny(jsonStr string) interface{} {
	var any interface{}
	_ = json.Unmarshal([]byte(jsonStr), &any)
	return any
}
