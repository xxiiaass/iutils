package iutils

import (
	"github.com/mitchellh/mapstructure"
	"strings"
)


// 将m的值，复制给s结构体中同名json-tag的字段，s必须为结构体指针
func Merge(m H, h H) H {
	for k, v := range h {
		m[k] = v
	}
	return m
}


func mapstructureExport(m interface{}, s interface{}, tagNames ...string) {
	tagName := "json"
	if len(tagNames) > 0 {
		tagName = tagNames[0]
	}

	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           s,
		TagName:          tagName,
		WeaklyTypedInput: true,
	}
	decoder, _ := mapstructure.NewDecoder(config)
	decoder.Decode(m)
}



func ToStruct(m interface{}, s interface{}, tagNames ...string) {
	if str, ok := m.(string); ok {
		JsonToStruct(str, s, tagNames...)
		return
	}
	// 非字符串内容，由他自己处理
	mapstructureExport(m, s, tagNames...)
}

// 将json字符串，使用mapstruct实现
func JsonToStruct(json string, s interface{}, tagNames ...string) {
	m := JsonDecodeSafe(json)
	if m == nil {
		return
	}
	mapstructureExport(m, s, tagNames...)
}



func deepGet(m map[string]interface{}, key string) (interface{}, bool) {
	keys := strings.Split(key, ".")
	var ret interface{} = m
	for _, k := range keys {
		v, ok := ret.(map[string]interface{})
		if !ok {
			return nil, false
		}
		ret, ok = v[k]
		if !ok {
			return nil, false
		}
	}
	return ret, true
}

func deepSet(m map[string]interface{}, key string, val interface{}) bool {
	keys := strings.Split(key, ".")
	var ret = m
	for i, k := range keys {
		if i == len(keys)-1 {
			ret[k] = val
			return true
		}
		tmpRet, ok := ret[k]
		if !ok {
			return false
		}
		tmpRetm, ok := tmpRet.(map[string]interface{})
		if !ok {
			return false
		}
		ret = tmpRetm
	}
	return false
}



func DeepGet(m map[string]interface{}, key string) (interface{}, bool) {
	return deepGet(m, key)
}

func DeepGetMust(m map[string]interface{}, key string) interface{} {
	v, ok := deepGet(m, key)
	if !ok {
		panic("没有该值")
	}
	return v
}

func DeepGetShould(m map[string]interface{}, key string) interface{} {
	v, _ := deepGet(m, key)
	return v
}

func DeepMustSet(m map[string]interface{}, key string, val interface{}) {
	succ := deepSet(m, key, val)
	if !succ {
		panic("设置失败")
	}
}