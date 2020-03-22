package tools

import "testing"

func TestValueAtBit(t *testing.T) {

	a := 1
	if ValueAtBit(a, 2) != 0 {
		t.Errorf("%d 的二进制数字 %b 第 2 位不是0", a, a)
	}
	if ValueAtBit(a, 1) != 1 {
		t.Errorf("%d 的二进制数字 %b 第 1 位不是1", a, a)
	}

	a = 15
	if ValueAtBit(a, 1) != 1 {
		t.Errorf("%d 的二进制数字 %b 第 1 位不是1", a, a)
	}
	if ValueAtBit(a, 2) != 1 {
		t.Errorf("%d 的二进制数字 %b 第 2 位不是1", a, a)
	}
	if ValueAtBit(a, 3) != 1 {
		t.Errorf("%d 的二进制数字 %b 第 3 位不是1", a, a)
	}
	if ValueAtBit(a, 4) != 1 {
		t.Errorf("%d 的二进制数字 %b 第 4 位不是1", a, a)
	}
}
