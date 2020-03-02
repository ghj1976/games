package main

import (
	"math/rand"
	"strconv"
	"time"
)

// 准备演示的数据, 随机地图
func prepareData1(rows, cols int) *AStarMap {

	amap := &AStarMap{}
	amap.MapRows = rows
	amap.MapCols = cols
	// 二维切片的初始化 https://blog.csdn.net/jiang_mingyi/article/details/81567740
	amap.nodeMap = make([][]int, rows)
	for y := 0; y < rows; y++ {
		amap.nodeMap[y] = make([]int, cols)
		for x := 0; x < cols; x++ {
			amap.nodeMap[y][x] = 0
		}
	}

	// 随机给 200 ~300 个点设置障碍
	rand.Seed(time.Now().Unix())
	for i := 0; i < 200+rand.Intn(100); i++ {
		m := rand.Intn(rows*cols - 1)
		y := m / rows
		x := m % rows
		if x == 0 && y == 0 {
			continue // 起点位置不能设置障碍
			// 终点可以设置障碍
		}
		amap.nodeMap[y][x] = 1
	}

	return amap
}

// 导入坦克大战的地图
func prepareData2(rows, cols int, mt string) *AStarMap {

	amap := &AStarMap{}
	amap.MapRows = rows
	amap.MapCols = cols
	// 二维切片的初始化 https://blog.csdn.net/jiang_mingyi/article/details/81567740
	amap.nodeMap = make([][]int, rows)
	for y := 0; y < rows; y++ {
		amap.nodeMap[y] = make([]int, cols)
		for x := 0; x < cols; x++ {
			amap.nodeMap[y][x] = 0
		}
	}

	// 加载数据
	x := 0
	y := 0
	for _, rv := range mt {

		v, err := strconv.Atoi(string(rv))
		if err != nil {
			continue
		}

		if x >= cols || y >= rows {
			//log.Println(v)
			break
		} else {
			amap.nodeMap[y][x] = v
		}

		// 下一轮的 x y 计算
		if x < cols-1 {
			x++
		} else {
			x = 0
			y++
		}
	}

	return amap

}
