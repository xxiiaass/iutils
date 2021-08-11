package iutils

import (
	"encoding/json"
	"strings"
)

func JsonEncode(v Any) string {
	bt, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bt)
}

func JsonDecode(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &mapResult)
	return mapResult
}

func JsonDecodeSafe(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	ds := json.NewDecoder(strings.NewReader(jsonStr))
	ds.UseNumber()
	_ = ds.Decode(&mapResult)
	return mapResult
}

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

func JsonDecodeAny(jsonStr string) interface{} {
	var any interface{}
	_ = json.Unmarshal([]byte(jsonStr), &any)
	return any
}
