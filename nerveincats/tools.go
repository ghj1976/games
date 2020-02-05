package nerveincats

import "math/rand"

// MaxInt 找出两个int的最大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinInt 找出两个int的最小值
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinIntValue 返回多个数据的最小值
func MinIntValue(arr ...int) int {
	min := arr[0]
	for _, v := range arr {
		if min > v {
			min = v
		}
	}
	return min
}

// RandomString 从给定的字符数组中随机出一个
func RandomString(arr []string) string {
	return arr[rand.Intn(len(arr))]
}
