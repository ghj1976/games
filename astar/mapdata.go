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

func prepareDataTank90(rows, cols int) *AStarMap {

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
	mt := "0000000000000000000000000000000000000000000000000000001100110011001100110011000011001100110011001100110000110011001000010011001100001100110010000100110011000011001100115511001100110000110011001155110011001100001100110011001100110011000011001100000000001100110000110011000000000011001100000000000011001100000000000000000000110011000000000011001111000000000011110011550011110000000000111100550000000000110011000000000000000000001111110000000000001100110011111100110011000011001100110011001100110000110011001100110011001100001100110011001100110011000011001100000000001100110000110011000000000011001100001100110001111000110011000000000000019910000000000000000000000199100000000000"
	x := 0
	y := 0
	for _, rv := range mt {

		v, err := strconv.Atoi(string(rv))
		if err != nil {
			continue
		}

		if x >= 26 || y >= 26 {
			//log.Println(v)
			break
		} else {
			amap.nodeMap[y][x] = v
		}

		// 下一轮的 x y 计算
		if x < 25 {
			x++
		} else {
			x = 0
			y++
		}
	}

	return amap

}
