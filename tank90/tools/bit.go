package tools

// ValueAtBit 获得num数字二进制时，bit位的值
func ValueAtBit(num, bit int) int {
	return (num >> (bit - 1)) & 1
}
