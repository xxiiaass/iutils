# 简介

本项目用于集成平时go项目中常用的一些小工具函数，减少各个项目中，重复的代码，让写代码多一点点快乐

* as_type 中有类型转换的函数，建议使用cast库
* calculate 有一些数字计算的方法
* hash 中是关于base64、加解密、md5的工具函数
* json 中是关于json格式化，编码、解码函数
* location 关于地域的工具函数, 省、市的处理
* map 关于结构体、map类型的快捷方法（重点推荐）
* random 关于随机数的快捷方法
* slice 对于数组的一些处理、查询、转换 （推荐）
* time 时间的便捷函数 
* tools 其他杂七杂八的
* uuid 唯一id生成
* weight_random 权重计算



```go
// 类型转换相关
// 建议使用更完善的工具库 https://github.com/spf13/cast

// 判断该值是否为表示空值，0，"0"， false, "false"等
func Empty(v interface{}) bool

// 判断一堆基础类型的值是不是都是空
func AllEmpty(vl ...interface{}) bool

// 判断interface是否为nil
func IsNil(i interface{}) bool

// 如果是字符串，则转化为int64, 浮点数丢失精度，无法转化则为0
func AsInt64(v interface{}) int64

func AsBool(v interface{}) bool

func AsInt64Slice(v interface{}) []int64

func AsInt64OrDefault(v interface{}, d int64) int64

func AsInt(v interface{}) int

func AsFloat64(v interface{}) float64

func AsString(v interface{}) string

func AsStringOrDefault(v interface{}, d string) string

func AsH(v interface{}) H
func AsStringSlice(v interface{}) []string

func SliceString2Int(src []string) ([]int, error)

// 字符串切片转换为 int64 切片
func SliceString2Int64(src []string) ([]int64, error)

// "s1, s2, s3, s4" => []string{"s1", "s2", "s3", "s4"}
// "[s1, s2, s3, s4]" => []string{"s1", "s2", "s3", "s4"}
// "["s1", "s2", "s3", "s4"]" => []string{"s1", "s2", "s3", "s4"}
func StrToStrSlice(str string) []string

func SliceToStr(arr []int64) string

func StrSliceToStr(arr []string) string

// "1,2,3,4" => []int{1, 2, 3, 4}
// "[1,2,3,4]" => []int{1, 2, 3, 4}
func StrToSlice(str string) []int64

// 将字符串转换为byte数组，默认使用标准转换方式
// 可以传参true，要求使用强转换，此时返回的byte数组将不能修改，否则会报错!!!
func StrToBytes(str string, useUnsafe ...bool) []byte

// 将byte数组转换为字符串，默认使用标准转换方式
// 可以传参true，要求使用强转换，此时返回的string将不再是只读，有可能会被修改!!
func BytesToStr(bs []byte, useUnsafe ...bool) string

// 计算绝对值
func Abs(num1, num2 int64) int64

// 最大值
func MaxNum(argus ...int64) int64

// 最小值
func MinNum(argus ...int64) int64

// 计算md5值
func MD5(str string) string

// 计算base64
func Base64(str string) string

// 计算Sha256
func Sha256(str string) string

func _PKCS5Padding(ciphertext []byte, blockSize int) []byte

// 使用aes-cbc加密方法加密数据
func AesCbcEncrypt(str string, key []byte, IV []byte) ([]byte, error)

// 使用aes-cbc加密方法解密数据

func AesCbcDecrypt(encrypted []byte, key []byte, IV []byte) ([]byte, error)

// 获取一个将字符串转为int64的算法句柄
func HashCode() hashCode

// 更方便的json数据解析，建议使用map.go文件中的toStruct方法，解析到结构体中

// 编码任意数据为json字符串
func JsonEncode(v Any) string

// 将json字符串解码成map结构 （此时，数据的结构会变成浮点数，超过浮点数精度的数据，会不准确）
func JsonDecode(jsonStr string) map[string]interface{}

// 将json字符串解码成map结构，能保留大数的数据
func JsonDecodeSafe(jsonStr string) map[string]interface{}

// 将json字符串解码成map结构，也可以是个数组，在两个返回值中返回一个，使用者需要知道数据是什么类型
func JsonDecodes(jsonStr string) (map[string]interface{}, []map[string]interface{})

// 可以解析任意json类型的数据，以interface{}类型返回
func JsonDecodeAny(jsonStr string) interface{}

// 将城市去除 `市`的后缀，省份去除`省`的后缀
func SimpleName(o string) string

// 将m的值，复制给s结构体中同名json-tag的字段，s必须为结构体指针
func Merge(m H, h H) H

// 将多个类型的数据，映射到结构体中
// @param m string|map|struct|[]map 源数据
// @param s *struct|*[]struct 接收数据的结构体，可以是结构体切片，但是必须是指针类型
// @param tagNames string 做字段映射的时候，默认使用json的tag，可以在这里指定使用什么tag
func ToStruct(m interface{}, s interface{}, tagNames ...string)

// 将json字符串，使用mapstruct实现
// @param json string 源数据json
// @param s *struct|*[]struct 接收数据的结构体，可以是结构体切片，但是必须是指针类型
// @param tagNames string 做字段映射的时候，默认使用json的tag，可以在这里指定使用什么tag
func JsonToStruct(json string, s interface{}, tagNames ...string)

// 能够多层级获取map中的数据
// @param m 待获取的数据源
// @param key 字段，格式为 a.b.c.d
// @return v 获取到的值
// @return ok 是否成功获取到了值
func DeepGet(m map[string]interface{}, key string) (v interface{}, ok bool)

// 能够多层级获取map中的数据，无法获取时抛出错误
// @param m 待获取的数据源
// @param key 字段，格式为 a.b.c.d
// @return 获取到的值
func DeepGetMust(m map[string]interface{}, key string) interface{}

// 能够多层级获取map中的数据，无法获取时返回nil
// @param m 待获取的数据源
// @param key 字段，格式为 a.b.c.d
// @return 获取到的值
func DeepGetShould(m map[string]interface{}, key string) interface{}

// 能够多层级设置map中的数据，失败时抛出异常
// @param m 待设置的数据源
// @param key 字段，格式为 a.b.c.d
// @param val 值
func DeepMustSet(m map[string]interface{}, key string, val interface{})

// 获取两个值之间的随机数
func RandInt64(min, max int64) int64

// 若size <= len , 生成常规的切片
// 若size > len , 生成切片的值为常规切片中的随机值
func RandSliceIdx(size, len int) []int

// 从切片中获取随机值，和剩余的其他值
func RandSlice(s []int) (int, []int)

// 通过自定义函数比较是否在切片中
// param @val 待判断的值
// param @array 目标切片
// param equalFunc 自定义比较函数
// return @exist 是否存在
// return @index 若存在返回在数组所在位置
func InArrayByFunc(val interface{}, array interface{}, equalFunc func(val, arrayItem interface{}) bool) (exist bool, index int)

// 判断val 是否在切片中
func IsExists(val interface{}, array interface{}) bool

// 判断某个值是否在切片中
func IsExistsByFunc(val interface{}, array interface{}, equalFunc func(val, arrayItem interface{}) bool) bool

// 获取val在切片中的位置
// return 若不存在返回值为-1
func IndexOf(val interface{}, array interface{}) int

// 通过自定义函数获取val在切片中的位置
// return 若不存在返回值为-1
func IndexOfByFunc(val interface{}, array interface{}, equalFunc func(val, arrayItem interface{}) bool) int

// 切片值去重
func Deduplication(arr []int64) []int64

// 删除切片中某个索引的值
func RemoveNum(arr []int64, target int64) []int64

// 删除交集
func RemoveIntersection(arr1, arr2 []int64) []int64

// 切片中的所有元素是否都满足指定方法
func All(array interface{}, function func(item interface{}) bool) bool

// 切片中是否存在满足指定方法的元素，如果不传方法，则判断切片中是否有元素，传多个方法只会取第一个
func AnyOne(array interface{}, function ...func(item interface{}) bool) bool

// 将切片中的每个元素都作为入参传入指定方法中，收集方法返回值并放入切片返回
func Select(array interface{}, function func(item interface{}) interface{}) []interface{}

// 返回切片中满足指定方法的元素个数
func Count(array interface{}, function func(item interface{}) bool) int64

// 返回切片中满足指定方法的元素
func Where(array interface{}, function func(item interface{}) bool) []interface{}

// 返回切片中第一个满足指定方法的元素，如不存在则返回nil
func First(array interface{}, function func(item interface{}) bool) interface{}

// 返回切片中第一个满足指定方法的元素，如不存在则返回入参中的默认值
func FirstOrDefault(array interface{}, function func(item interface{}) bool, def interface{}) interface{}

// 判断时间段是否跨天
func IsCrossDay(begin int64, ends ...int64) bool

// 日期时间点(YYYY-MM-DD HH:dd:ss)转时间戳
func DatetimeToUnix(date string) int64

// 年月日(YYYY-MM-DD)转时间戳
func DateToUnix(date string) int64

// 时间戳转日期时间点(YYYY-MM-DD HH:dd:ss)
func UnixToDate(unix int64) string

// 获取今天凌晨对应的时间戳
func TodayBegin() time.Time

// 获取某个时间对应当天凌晨时间戳
func Begin(t time.Time) time.Time

// 获取当天23:59:59 时间点对应的时间戳
func TodayEnd() time.Time

// 获取某个时间对应当天23:59:59 时间点对应的时间戳
func End(t time.Time) time.Time

// 获取最近几天相同时刻
func GetLatestDaysBaseNow(days int, desc ...bool) []time.Time

// 获取自某个时刻最近几天的相同时刻（包含当日）
// base 起点时刻
// days 最近的几天，< 0 表示过去几天；> 0 表示之后几天
// desc[0] = true 倒序输出
func GetLatestDays(base time.Time, days int, desc ...bool) (dates []time.Time)

// 获取该时间当天的起止时间点
func GetThisDayBeginEnd(t time.Time) (begin, end time.Time)

// 获取距离明天凌晨的时间长度
func DurationUntilTomorrow() time.Duration

// 获取当前时间戳，秒级
func GetNowUnix() int64

// 获取当前时间戳，毫秒级
func GetNowUnixMilli() int64
// 获取当前时间戳，纳秒级
func GetNowUnixNano() int64

// 获取本周所在年份、第几周，起止时间点
func WeekInfo(t ...time.Time) (year, isoWeek int, weekBegin, weekEnd time.Time)

// 简易的年月日转字符串
// 使用前需自行校验 year, month, day 有效性
func YmdStr(year, month, day int, sep ...string) string

// 获取两个时间间隔天数
func SubDay(time1, time2 time.Time) int64

// 获取某个时间对应的小时
func GetDayHourNum(t int64) int

// 获取时间错在当天已过的时长秒
func GetOneDaySec(t int64) int64

// 获取最近的指定月份、日的当天的起始时间, 可以指定特定时间为需要的"当前"时间
func LastMonthDay(month, day int, nows ...time.Time) time.Time

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time

// 获取当前 时分秒 的秒数
func GetTimeSeconds() int64

// 版本比较
func VersionThan(va string, vb string) int

// AppendStuct 将give的结构体中，同名的成员赋值给reciv
func AppendStuct(reciv Any, give Any)

// 组装 http url 的key 和value
// param @v url 的值
// param @keys url key列表
// param @withOutEncode 是否需要编码
// return 返回组装完成的 http url key/value
func EncodeUrlWithoutSort(v url.Values, keys []string, withOutEncode bool) string

// 判断val 是否在arry 中
func InArray(val interface{}, array interface{}) (exists bool, index int)

// 将下划线风格的单词变为驼峰命名的单词
func UnderLineToCamel(line string) string

// 合并两个map
func MergeMap(argus ...map[string]interface{}) map[string]interface{}

// http get 请求
// return http get response body
// return error
func HttpGet(url string) (map[string]interface{}, error)

// 设置连接超时时间
// param @cTimeout conntect timeout second
// param @rwTimeout read/write timeout second
// return func
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error)

// 获取 map key 列表
func MapKeys(item map[string]interface{}) []string

// http post 请求
// param @url http url
// param @params http post body
// param @seconds connect/read/write timeout second
// param @header http post header
// return response body
// return post error
func HttpPost(url string, params interface{}, seconds int, header ...interface{}) (map[string]interface{}, error)

// 模仿 php 中的 mb_strimwidth
// 字显示宽度为 1 或 2
func SubStrByShowLen(s, suffix string, l int) string

// map 深度拷贝
func CopyMapTopLevel(src H) H

// 比较来个slice 的差异
func Diff(arr1 []int64, arr2 []int64) []int64
// 获取slice 中最大值
func MaxInSlice(arr []int64) int64

// 获取当前路径
func GetCurPath() string

// 取模
func Ceil(n int, c int) int

// 获取文件后缀
func GetFileExt(path string) string

// 按显示长度截取字符串
func ShowSubstr(s string, l int) string

// 压缩字符串，去除空格、制表符、换页符等字符
func CompressStr(str string) string

// You can get a substring of a UTF-8 string without allocating additional memory (you don't have to convert it to a rune slice)
// @from: https://stackoverflow.com/questions/28718682/how-to-get-a-substring-from-a-string-of-runes-in-golang
func Substr(s string, start int, end int) string

// Range(n) return [0,1,...,n-1,n] with []int;
// Range(n,m) return [n,n+1,...,m-1,m] with []int
func Range(first int, second ...int) []int

// Range64(n) return [0,1,...,n-1,n] with []int64;
// Range64(n,m) return [n,n+1,...,m-1,m] with []int64
func Range64(first int64, second ...int64) []int64

// 驼峰转蛇形
func SnakeString(s string) string

// 过滤字符串数组中的空字符串
func FilterEmptyStr(arr []string) []string

// 将字符串以逗号分隔开
func SplitComma(str string) []string

// 数值转人民币
func Number2Chinese(number int64, money ...bool) (chinese string)

// 宽字符长度计算
func GetStringLenInfo(str string) (runeLen int, displayLen int)

func (u *UUID) setVersion(version byte)

// UUID3 generates a version 3 UUID
func UUID3(namespace UUID, name string) UUID

// UUID4 generates a version 4 UUID
func UUID4() UUID

// UUID5 generates a version 5 UUID
func UUID5(namespace UUID, name string) UUID

// NewUUID 默认使用 UUID4
func NewUUID() string

```