package tools

// BrickAttacked 砖墙被攻击后的剩余残留区域的计算
// 二进制存储的位信息如下：
// 4  3  2  1
// 上 下 左 右
// *  *  *  *
func BrickAttacked(oldArea, power int, from string) int {
	if power >= 2 {
		return 0
	}
	switch from {
	case "top":
		if ValueAtBit(oldArea, 4) == 0 {
			return 0
		}
		return oldArea &^ 0b1000

	case "bottom":
		if ValueAtBit(oldArea, 3) == 0 {
			return 0
		}
		return oldArea &^ 0b0100

	case "left":
		if ValueAtBit(oldArea, 2) == 0 {
			return 0
		}
		return oldArea &^ 0b0010

	case "right":
		if ValueAtBit(oldArea, 1) == 0 {
			return 0
		}
		return oldArea &^ 0b0001

	default:
		return oldArea
	}
}

// GetBrickArea 得到绘图区域位移参数
func GetBrickArea(area int) (top, bottom, left, right int) {
	if area == 0 {
		return 0, 0, 0, 0
	}

	top = 1
	if ValueAtBit(area, 4) == 0 {
		top = 2
	}

	bottom = 2
	if ValueAtBit(area, 3) == 0 {
		bottom = 1
	}

	left = 1
	if ValueAtBit(area, 2) == 0 {
		left = 2
	}

	right = 2
	if ValueAtBit(area, 1) == 0 {
		right = 1
	}
	return top, bottom, left, right
}
