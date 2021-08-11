package iutils

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// 版本比较
func VersionThan(va string, vb string) int {
	vas := strings.Split(va, ".")
	vbs := strings.Split(vb, ".")

	count := len(vas)
	if len(vas) < len(vbs) {
		count = len(vbs)
	}
	for i := 0; i < count; i++ {
		a := 0
		if len(vas) < i+1 {
			a = 0
		} else {
			a, _ = strconv.Atoi(vas[i])
		}

		b := 0
		if len(vbs) < i+1 {
			b = 0
		} else {
			b, _ = strconv.Atoi(vbs[i])
		}
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
	}
	return 0
}

// AppendStuct 将give的结构体中，同名的成员赋值给reciv
func AppendStuct(reciv Any, give Any) {
	gtyp := reflect.TypeOf(give).Elem()
	gval := reflect.ValueOf(give).Elem()

	rtyp := reflect.TypeOf(reciv).Elem()
	rval := reflect.ValueOf(reciv).Elem()

	for i := 0; i < gtyp.NumField(); i++ {
		gname := gtyp.Field(i).Name
		for j := 0; j < rtyp.NumField(); j++ {
			if gname == rtyp.Field(j).Name {
				rval.Field(j).Set(gval.Field(i))
				break
			}
		}
	}
}

func EncodeUrlWithoutSort(v url.Values, keys []string, withOutEncode bool) string {
	if v == nil {
		return ""
	}
	var buf string
	for _, k := range keys {
		vs, ok := v[k]
		if !ok {
			continue
		}
		keyEscaped := url.QueryEscape(k)
		if withOutEncode {
			keyEscaped = k
		}
		for _, vl := range vs {
			if len(buf) > 0 {
				buf += "&"
			}
			buf += keyEscaped
			buf += "="
			if withOutEncode {
				buf += vl
			} else {
				buf += url.QueryEscape(vl)
			}
		}
	}
	res := buf
	return res
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// 将下划线风格的单词变为驼峰命名的单词
func UnderLineToCamel(line string) string {
	words := strings.Split(line, "_")
	n := ""
	for _, w := range words {
		n += strings.ToUpper(w[0:1]) + w[1:]
	}
	return n
}

func MergeMap(argus ...map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for _, m := range argus {
		for k, v := range m {
			ret[k] = v
		}
	}
	return ret
}

func HttpGet(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rep := JsonDecode(string(body))
	return rep, nil
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func MapKeys(item map[string]interface{}) []string {
	ks := make([]string, 0)
	for k := range item {
		ks = append(ks, k)
	}
	return ks
}

func HttpPost(url string, params interface{}, seconds int, header ...interface{}) (map[string]interface{}, error) {
	body := JsonEncode(params)

	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if len(header) > 0 {
		headers := header[0].(map[string]string)
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	connectTimeout := time.Duration(seconds) * time.Second
	readWriteTimeout := time.Duration(seconds) * time.Second
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial:            TimeoutDialer(connectTimeout, readWriteTimeout),
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	rep := JsonDecode(string(content))
	return rep, nil
}

// 模仿 php 中的 mb_strimwidth
// 字显示宽度为 1 或 2
func SubStrByShowLen(s, suffix string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	suffixl := len(suffix)
	for _, r := range rs {
		rint := int(r) // 获取字节值
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl+suffixl > l {
			break
		}
		sl += rl
		ss += string(r)
	}

	if sl < suffixl {
		return ss
	}
	return ss + suffix
}

func CopyMapTopLevel(src H) H {
	res := make(H, len(src))
	for k, v := range src {
		res[k] = v
	}
	return res
}

func Diff(arr1 []int64, arr2 []int64) []int64 {
	ret := make([]int64, 0)
	for _, i := range arr1 {
		if !IsExists(i, arr2) {
			ret = append(ret, i)
		}
	}
	return ret
}

func MaxInSlice(arr []int64) int64 {
	max := int64(math.MinInt64)
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}

func Ceil(n int, c int) int {
	if n%c == 0 {
		return n
	}
	return int(math.Ceil(float64(n)/float64(c))) * c
}

func GetFileExt(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}

// 按显示长度截取字符串
func ShowSubstr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

// 压缩字符串，去除空格、制表符、换页符等字符
func CompressStr(str string) string {
	if str == "" {
		return ""
	}

	// \s 匹配任何空白字符，包括空格、制表符、换页符等等，等价于 \f\n\r\t\v
	// Unicode 正则表达式会匹配全角空格
	re := regexp.MustCompile("\\s+")

	return re.ReplaceAllString(str, "")
}

// You can get a substring of a UTF-8 string without allocating additional memory (you don't have to convert it to a rune slice)
// @from: https://stackoverflow.com/questions/28718682/how-to-get-a-substring-from-a-string-of-runes-in-golang
func Substr(s string, start int, end int) string {
	var a, b int
	for i := range s {
		if b == start {
			a = i
		}
		if b == end {
			return s[a:i]
		}
		b++
	}

	return s[a:]
}

// Range(n) return [0,1,...,n-1,n] with []int;
// Range(n,m) return [n,n+1,...,m-1,m] with []int
func Range(first int, second ...int) []int {
	var start, end int
	if len(second) == 0 {
		end = first
	} else {
		start, end = first, second[0]
	}
	if start > end {
		return []int{}
	}
	res := make([]int, end-start+1)
	for i := start; i <= end; i++ {
		res[i-start] = i
	}
	return res
}

// Range64(n) return [0,1,...,n-1,n] with []int64;
// Range64(n,m) return [n,n+1,...,m-1,m] with []int64
func Range64(first int64, second ...int64) []int64 {
	var start, end int64
	if len(second) == 0 {
		end = first
	} else {
		start, end = first, second[0]
	}
	if start > end {
		return []int64{}
	}
	res := make([]int64, end-start+1)
	for i := start; i <= end; i++ {
		res[i-start] = i
	}
	return res
}

// 驼峰转蛇形
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// 过滤字符串数组中的空字符串
func FilterEmptyStr(arr []string) []string {
	ret := make([]string, 0)
	for _, i := range arr {
		if strings.TrimSpace(i) == "" {
			continue
		}
		ret = append(ret, strings.TrimSpace(i))
	}
	return ret
}

// 将字符串以逗号分隔开
func SplitComma(str string) []string {
	str = strings.ReplaceAll(str, "，", ",")
	arr := strings.Split(str, ",")
	return FilterEmptyStr(arr)
}

func Number2Chinese(number int64, money ...bool) (chinese string) {
	isMoney := len(money) > 0 && money[0]
	if number == 0 {
		if isMoney {
			return "零圆整"
		}
		return "零"
	}
	if number < 0 {
		number *= -1
		defer func() {
			chinese = fmt.Sprintf("负%s", chinese)
		}()
	}
	chineseMap := []string{"%", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	chineseNum := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	if isMoney {
		chineseMap = []string{"圆整", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟"}
		chineseNum = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	} else {
		defer func() {
			chinese = strings.TrimSuffix(chinese, "%")
			if strings.HasPrefix(chinese, "一十") {
				chinese = strings.Replace(chinese, "一十", "十", 1)
			}
		}()
	}
	listNum := make([]int64, 0)
	for ; number > 0; number = number / 10 {
		listNum = append(listNum, number%10)
	}
	for i := len(listNum) - 1; i >= 0; i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}
	str := ""
	for {
		str = chinese
		str = strings.Replace(str, "零亿", "亿", 1)
		str = strings.Replace(str, "零万", "万", 1)
		if isMoney {
			str = strings.Replace(str, "零拾", "零", 1)
			str = strings.Replace(str, "零佰", "零", 1)
			str = strings.Replace(str, "零仟", "零", 1)
			str = strings.Replace(str, "零零", "零", 1)
			str = strings.Replace(str, "零圆", "圆", 1)
		} else {
			str = strings.Replace(str, "零十", "零", 1)
			str = strings.Replace(str, "零百", "零", 1)
			str = strings.Replace(str, "零千", "零", 1)
			str = strings.Replace(str, "零零", "零", 1)
			str = strings.Replace(str, "零%", "%", 1)
		}
		if str == chinese {
			chinese = str
			break
		}
		chinese = str
	}
	return
}

func GetStringLenInfo(str string) (runeLen int, displayLen int) {
	runeLen = utf8.RuneCountInString(str)
	s := []byte(str)
	for i := 0; i < len(str); {
		_, width := utf8.DecodeRune(s[i:])
		i += width
		if width > 1 {
			displayLen += 2
		} else {
			displayLen++
		}
	}
	return
}
