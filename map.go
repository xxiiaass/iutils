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


// 将多个类型的数据，映射到结构体中
// @param m string|map|struct|[]map 源数据
// @param s *struct|*[]struct 接收数据的结构体，可以是结构体切片，但是必须是指针类型
// @param tagNames string 做字段映射的时候，默认使用json的tag，可以在这里指定使用什么tag
func ToStruct(m interface{}, s interface{}, tagNames ...string) {
	if str, ok := m.(string); ok {
		JsonToStruct(str, s, tagNames...)
		return
	}
	// 非字符串内容，由他自己处理
	mapstructureExport(m, s, tagNames...)
}

// 将json字符串，使用mapstruct实现
// @param json string 源数据json
// @param s *struct|*[]struct 接收数据的结构体，可以是结构体切片，但是必须是指针类型
// @param tagNames string 做字段映射的时候，默认使用json的tag，可以在这里指定使用什么tag
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


// 能够多层级获取map中的数据
// @param m 待获取的数据源
// @param key 字段，格式为 a.b.c.d
// @return v 获取到的值
// @return ok 是否成功获取到了值
func DeepGet(m map[string]interface{}, key string) (v interface{}, ok bool) {
	return deepGet(m, key)
}

// 能够多层级获取map中的数据，无法获取时抛出错误
// @param m 待获取的数据源
// @param key 字段，格式为 a.b.c.d
// @return 获取到的值
func DeepGetMust(m map[string]interface{}, key string) interface{} {
	v, ok := deepGet(m, key)
	if !ok {
		panic("没有该值")
	}
	return v
}

// 能够多层级获取map中的数据，无法获取时返回nil
// @param m 待获取的数据源
// @param key 字段，格式为 a.b.c.d
// @return 获取到的值
func DeepGetShould(m map[string]interface{}, key string) interface{} {
	v, _ := deepGet(m, key)
	return v
}


// 能够多层级设置map中的数据，失败时抛出异常
// @param m 待设置的数据源
// @param key 字段，格式为 a.b.c.d
// @param val 值
func DeepMustSet(m map[string]interface{}, key string, val interface{}) {
	succ := deepSet(m, key, val)
	if !succ {
		panic("设置失败")
	}
}