package iutils

import (
	"math/rand"
	"time"
)

// 权重工具

type WrEle struct {
	Key    string
	Weight int
}

type WrList []WrEle

// 获取总权重
func (list WrList) TotalWeight() int {
	total := 0
	for _, i := range list {
		total += i.Weight
	}
	return total
}

// 获取随机权重的key
func (list WrList) Rand() string {
	total := list.TotalWeight()

	rand.Seed(int64(time.Now().UnixNano()))
	randRet := rand.Int() % total
	cur := 0
	for _, v := range list {
		cur += v.Weight
		if randRet < cur {
			return v.Key
		}
	}
	panic("错误随机")
}
