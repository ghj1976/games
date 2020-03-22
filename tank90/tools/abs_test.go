package tools

import "testing"

func TestAbsInt(t *testing.T) {
	if AbsInt(0) != 0 {
		t.Error("0的绝对值应该也是0！")
	}

	if AbsInt(-9) != 9 {
		t.Error("-9的绝对值应该是9！")
	}

	if AbsInt(8) != 8 {
		t.Error("8的绝对值应该是8！")
	}
}
