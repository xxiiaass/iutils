# 简介

本项目用于集成平时go项目中常用的一些小工具函数，减少各个项目中，重复的代码，让写代码多一点点快乐

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
func InArrayByFunc(val interface}, array interface}, equalFunc func(val, arrayItem interface}) bool) (exist bool, index int)

// 判断val 是否在切片中
func IsExists(val interface}, array interface}) bool

// 判断某个值是否在切片中
func IsExistsByFunc(val interface}, array interface}, equalFunc func(val, arrayItem interface}) bool) bool

// 获取val在切片中的位置
// return 若不存在返回值为-1
func IndexOf(val interface}, array interface}) int

// 通过自定义函数获取val在切片中的位置
// return 若不存在返回值为-1
func IndexOfByFunc(val interface}, array interface}, equalFunc func(val, arrayItem interface}) bool) int

// 切片值去重
func Deduplication(arr []int64) []int64

// 删除切片中某个索引的值
func RemoveNum(arr []int64, target int64) []int64

// 删除交集
func RemoveIntersection(arr1, arr2 []int64) []int64

// 切片中的所有元素是否都满足指定方法
func All(array interface}, function func(item interface}) bool) bool

// 切片中是否存在满足指定方法的元素，如果不传方法，则判断切片中是否有元素，传多个方法只会取第一个
func AnyOne(array interface}, function ...func(item interface}) bool) bool

// 将切片中的每个元素都作为入参传入指定方法中，收集方法返回值并放入切片返回
func Select(array interface}, function func(item interface}) interface}) []interface}

// 返回切片中满足指定方法的元素个数
func Count(array interface}, function func(item interface}) bool) int64

// 返回切片中满足指定方法的元素
func Where(array interface}, function func(item interface}) bool) []interface}

// 返回切片中第一个满足指定方法的元素，如不存在则返回nil
func First(array interface}, function func(item interface}) bool) interface}

// 返回切片中第一个满足指定方法的元素，如不存在则返回入参中的默认值
func FirstOrDefault(array interface}, function func(item interface}) bool, def interface}) interface}

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
func InArray(val interface}, array interface}) (exists bool, index int)

// 将下划线风格的单词变为驼峰命名的单词
func UnderLineToCamel(line string) string

// 合并两个map
func MergeMap(argus ...map[string]interface}) map[string]interface}

// http get 请求
// return http get response body
// return error
func HttpGet(url string) (map[string]interface}, error)

// 设置连接超时时间
// param @cTimeout conntect timeout second
// param @rwTimeout read/write timeout second
// return func
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error)

// 获取 map key 列表
func MapKeys(item map[string]interface}) []string

// http post 请求
// param @url http url
// param @params http post body
// param @seconds connect/read/write timeout second
// param @header http post header
// return response body
// return post error
func HttpPost(url string, params interface}, seconds int, header ...interface}) (map[string]interface}, error)

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

// 获取UUID 字符串内容
func (u UUID) String() string

// UUID3 generates a version 3 UUID
func UUID3(namespace UUID, name string) UUID

// UUID4 generates a version 4 UUID
func UUID4() UUID

// UUID5 generates a version 5 UUID
func UUID5(namespace UUID, name string) UUID

// NewUUID 默认使用 UUID4
func NewUUID() string

// 获取总权重
func (list WrList) TotalWeight() int

// 获取随机权重的key
func (list WrList) Rand() string 



