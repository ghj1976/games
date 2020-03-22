package tools

import "testing"

func TestBrickAttacked(t *testing.T) {

	if BrickAttacked(15, 2, "top") != 0 {
		t.Error("威力为2的坦克没能击毁墙！")
	}

	a := 0b1111
	b := 0b0111
	from := "top"
	if BrickAttacked(a, 1, from) != b {
		t.Errorf("来自 %s 的攻击，无法把 %b 攻击成 %b ", from, a, b)
	}
	a = 0b0111
	b = 0b0000
	if BrickAttacked(a, 1, from) != b {
		t.Errorf("来自 %s 的攻击，无法把 %b 攻击成 %b ", from, a, b)
	}

	a = 0b0111
	b = 0b0101
	from = "left"
	if BrickAttacked(a, 1, from) != b {
		t.Errorf("来自 %s 的攻击，无法把 %b 攻击成 %b ", from, a, b)
	}

	a = 0b0101
	b = 0b0000
	from = "left"
	if BrickAttacked(a, 1, from) != b {
		t.Errorf("来自 %s 的攻击，无法把 %b 攻击成 %b ", from, a, b)
	}

}

func TestGetBrickArea(t *testing.T) {
	area := 0b1111
	top, bottom, left, right := GetBrickArea(area)
	rtop, rbottom, rleft, rright := 1, 2, 1, 2
	if top != rtop || bottom != rbottom || left != rleft || right != rright {
		t.Errorf("来自 %b 无法转换成显示模式 %d%d%d%d  ", area, top, bottom, left, right)
	}

	area = 0b0000
	top, bottom, left, right = GetBrickArea(area)
	rtop, rbottom, rleft, rright = 0, 0, 0, 0
	if top != rtop || bottom != rbottom || left != rleft || right != rright {
		t.Errorf("来自 %b 无法转换成显示模式 %d%d%d%d  ", area, top, bottom, left, right)
	}

	area = 0b1001
	top, bottom, left, right = GetBrickArea(area)
	rtop, rbottom, rleft, rright = 1, 1, 2, 2
	if top != rtop || bottom != rbottom || left != rleft || right != rright {
		t.Errorf("来自 %b 无法转换成显示模式 %d%d%d%d  ", area, top, bottom, left, right)
	}

	area = 0b1010
	top, bottom, left, right = GetBrickArea(area)
	rtop, rbottom, rleft, rright = 1, 1, 1, 1
	if top != rtop || bottom != rbottom || left != rleft || right != rright {
		t.Errorf("来自 %b 无法转换成显示模式 %d%d%d%d  ", area, top, bottom, left, right)
	}

	area = 0b1011
	top, bottom, left, right = GetBrickArea(area)
	rtop, rbottom, rleft, rright = 1, 1, 1, 2
	if top != rtop || bottom != rbottom || left != rleft || right != rright {
		t.Errorf("来自 %b 无法转换成显示模式 %d%d%d%d  ", area, top, bottom, left, right)
	}
}
