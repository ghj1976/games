package main

import (
	"math/rand"
	"strconv"
	"time"
)

// 初始化 AStarMap
// mapType 地图类型
// hScoreType A*算法中，H值的计算方法
func Prepare(demoMapType, hScoreType int) *AStarMap {
	amap := &AStarMap{}
	amap.HScoreType = hScoreType
	amap.DemoMapType = demoMapType

	return amap
}

// Reset 重新开始
func (m *AStarMap) Reset(demoMapType, hScoreType int) {
	if m.StopCh != nil {
		close(m.StopCh)
	}

	// m.GoRunStatus = "stop"
	m.HScoreType = hScoreType
	m.DemoMapType = demoMapType

	if demoMapType == 1 {
		// 随机地图
		m.MapRows, m.MapCols = 50, 50
		m.prepareData1()
	} else if demoMapType == 2 {

		// 坦克大战的地图
		m.MapRows, m.MapCols = 26, 26
		mt := "0000000000000000000000000000000000000000000000000000001100110011001100110011000011001100110011001100110000110011001000010011001100001100110010000100110011000011001100115511001100110000110011001155110011001100001100110011001100110011000011001100000000001100110000110011000000000011001100000000000011001100000000000000000000110011000000000011001111000000000011110011550011110000000000111100550000000000110011000000000000000000001111110000000000001100110011111100110011000011001100110011001100110000110011001100110011001100001100110011001100110011000011001100000000001100110000110011000000000011001100001100110001111000110011000000000000019910000000000000000000000199100000000000"
		m.prepareData2(mt)
	} else {

		// U 型地图
		m.MapRows, m.MapCols = 26, 26
		mt := `
	00000000000000000000000000
	00000000000000000000000000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00000000000000000000001000
	00111111111111111111111000
	00000000000000000000000000
	00000000000000000000000000
	`
		m.prepareData2(mt)
	}

	m.MapWidth = 2 + (5+2)*m.MapCols
	m.MapHeight = 2 + (5+2)*m.MapRows

	m.StopCh = make(chan struct{})

	go m.FindPath(m.StopCh, Point{Row: 0, Col: 0}, Point{Row: m.MapRows - 1, Col: m.MapCols - 1})

}

// 准备演示的数据, 随机地图
func prepareData1(rows, cols, hScoreType int) *AStarMap {

	amap := &AStarMap{}
	amap.HScoreType = hScoreType
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
func prepareData2(rows, cols int, mt string, hScoreType int) *AStarMap {

	amap := &AStarMap{}
	amap.HScoreType = hScoreType
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

// prepareData1 按照之前给定的 m.MapRows ， m.MapCols 随机产生地图
func (m *AStarMap) prepareData1() {

	// 二维切片的初始化 https://blog.csdn.net/jiang_mingyi/article/details/81567740
	m.nodeMap = make([][]int, m.MapRows)
	for y := 0; y < m.MapRows; y++ {
		m.nodeMap[y] = make([]int, m.MapCols)
		for x := 0; x < m.MapCols; x++ {
			m.nodeMap[y][x] = 0
		}
	}

	// 随机给 200 ~300 个点设置障碍
	rand.Seed(time.Now().Unix())
	for i := 0; i < 200+rand.Intn(100); i++ {
		mr := rand.Intn(m.MapRows*m.MapCols - 1)
		y := mr / m.MapRows
		x := mr % m.MapRows
		if x == 0 && y == 0 {
			continue // 起点位置不能设置障碍
			// 终点可以设置障碍
		}
		m.nodeMap[y][x] = 1
	}
}

// 按照字符串的定义，给地图赋值
// 注意地图的大小事先在 m.MapRows ， m.MapCols 来指定
func (m *AStarMap) prepareData2(mt string) {
	// 二维切片的初始化 https://blog.csdn.net/jiang_mingyi/article/details/81567740
	m.nodeMap = make([][]int, m.MapRows)
	for y := 0; y < m.MapRows; y++ {
		m.nodeMap[y] = make([]int, m.MapCols)
		for x := 0; x < m.MapCols; x++ {
			m.nodeMap[y][x] = 0
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

		if x >= m.MapCols || y >= m.MapRows {
			//log.Println(v)
			break
		} else {
			m.nodeMap[y][x] = v
		}

		// 下一轮的 x y 计算
		if x < m.MapCols-1 {
			x++
		} else {
			x = 0
			y++
		}
	}

}
