package tools

// AbsInt 取正数的绝对值
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
