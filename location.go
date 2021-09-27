package iutils

import "regexp"

// 将城市去除 `市`的后缀，省份去除`省`的后缀
func SimpleName(o string) string {
	reg := regexp.MustCompile("市$")
	o = reg.ReplaceAllString(o, "")
	reg = regexp.MustCompile("省$")
	return reg.ReplaceAllString(o, "")
}
