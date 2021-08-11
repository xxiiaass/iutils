package iutils


// 计算绝对值
func CalculateSubInt64Abs(num1, num2 int64) int64 {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}

// 最大值
func MaxNum(argus ...int64) int64 {
	max := argus[0]
	for _, i := range argus {
		if i > max {
			max = i
		}
	}
	return max
}

// 最小值
func MinNum(argus ...int64) int64 {
	min := argus[0]
	for _, i := range argus {
		if i < min {
			min = i
		}
	}
	return min
}


func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
