package tank90

import (
	"log"
	"math/rand"
)

/*
人工智能部分
闲逛模式的机器人相关代码
*/

// GetRandomTarget 随机找一个坦克要去的目地
func GetRandomTarget(gmap [26][26]int, shipable bool) Point {
	for {
		seed := rand.Intn(26 * 26)
		y := seed / 26
		x := seed % 26

		if CheckTankStayable(gmap, x, y, shipable) {
			return Point{Row: y, Col: x}
		}
	}
}

// CheckTankStayable 检查指定的位置，坦克是否可停留
// x,y 的坐标是坦克左上角的坐标
// shipable 目前坦克是否可船运，可船运的可以停留在水域，否则不可以。
func CheckTankStayable(mapArr [26][26]int, x, y int, shipable bool) bool {
	if checkStayable(mapArr, x, y, shipable) &&
		checkStayable(mapArr, x+1, y, shipable) &&
		checkStayable(mapArr, x, y+1, shipable) &&
		checkStayable(mapArr, x+1, y+1, shipable) {
		return true
	}
	return false
}

// 检查一个节点
func checkStayable(mapArr [26][26]int, x, y int, shipable bool) bool {
	if y < 0 || y > 25 {
		return false // 越界了
	}
	if x < 0 || x > 25 {
		return false // 越界
	}
	if mapArr[y][x] == 0 {
		return true
	}
	if mapArr[y][x] == 3 && shipable {
		return true
	}
	return false
}

// getNeighborsTowward 从相邻的 source 到 target， 是那个朝向走动， 基于 source
func getNeighborsTowward(source, target Point) string {
	if source.Row == target.Row { // 行相同， Y 轴值不变。
		if source.Col < target.Col {
			return "right"
		}
		return "left"

	} else if source.Col == target.Col { // 列相同， X 轴值不变
		if source.Row < target.Row {
			return "bottom"
		}
		return "top"
	} else {
		log.Fatalf("不是相邻的两点: %v %v", source, target)
	}
	return ""
}
