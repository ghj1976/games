package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	// MapRows 地图的尺寸， 行数
	MapRows = 50
	// MapCols 地图的尺寸， 列数
	MapCols = 50
)

// Point 节点
type Point struct {
	Row int
	Col int
}

// PathPoint 路径中的点
type PathPoint struct {
	Point
	Parent *PathPoint
	HScore int
	GScore int
}

// GetFScore 获得 F  值
func (pp *PathPoint) GetFScore() int {
	return pp.HScore + pp.GScore
}

// AStarMap 演示 A* 寻路算法的类
type AStarMap struct {
	nodeMap [MapRows][MapCols]int // 地图 node 是 0 标示可以通行， 非零不可通行

	// 为了实时演示，把这些提取到这里，正常可以是局部变量
	openList  *sync.Map
	closeList *sync.Map
	currN     *PathPoint
}

// 准备演示的数据
func prepareData() *AStarMap {
	amap := &AStarMap{}
	amap.nodeMap = [50][50]int{}
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			amap.nodeMap[y][x] = 0
		}
	}

	// 随机给 200 ~300 个点设置障碍
	rand.Seed(time.Now().Unix())
	for i := 0; i < 200+rand.Intn(100); i++ {
		m := rand.Intn(Size*Size - 1)
		y := m / Size
		x := m % Size
		if x == 0 && y == 0 {
			continue // 起点位置不能设置障碍
			// 终点可以设置障碍
		}
		amap.nodeMap[y][x] = 1
	}

	return amap
}

// FindPath 通过A*算法寻找一个最短路径
// source, target 起点 和 终点
func (m *AStarMap) FindPath(source, target Point) *PathPoint {
	m.openList = &sync.Map{}
	m.closeList = &sync.Map{}

	// 算起点的h(s)
	sourcePathPoint := &PathPoint{Point: source, Parent: nil, GScore: 0}
	sourcePathPoint.HScore = m.getHScore(source, target)

	// 将起点放入OPEN表
	m.openList.Store(source, sourcePathPoint)

	for {
		// 从OPEN表中取f(n)最小的节点n;
		n := getMinFScore(m.openList)
		if n == nil { // open 列表没有数据， 则退出
			break
		}

		m.currN = n // 实时显示用到， 跟计算没关系
		log.Printf("curr: %d,%d", n.Col, n.Row)
		time.Sleep(250 * time.Millisecond)

		if n.Point == target { // 找到目标节点了
			return n
		}

		// 遍历 n 节点的每个临近节点
		for _, x := range m.getNeighbors(n.Point) {
			// 计算f(X);
			x.HScore = m.getHScore(x.Point, target)
			x.GScore = n.GScore + 1

			ox, oexist := m.openList.Load(x.Point)
			if oexist { // X in OPEN
				if x.GetFScore() < ox.(*PathPoint).GetFScore() {
					// Open 中只存最小 F 值的信息。
					// 如果有多条路都可以到达 x 节点， 只存最小的那条
					x.Parent = n
					m.openList.Store(x.Point, x)
				}
			}

			_, cexist := m.closeList.Load(x.Point)
			if cexist { // X in CLOSE
				// 已经不用处理了， 继续下一个
				continue
			}

			if !oexist && !cexist { // X not in both
				// 把n设置为X的父亲
				x.Parent = n
				// 求f(X); 循环进入时已经处理了，这里不用处理
				// 并将X插入OPEN表中;//还没有排序
				m.openList.Store(x.Point, x)

			}

		}
		//  将n节点插入CLOSE表中;
		m.closeList.Store(n.Point, n)

		//  按照f(n)将OPEN表中的节点排序; //实际上是比较OPEN表内节点f的大小，从最小路径的节点向下进行。
		// 这里每次提取最小节点是循环找的，不用排序

	}

	return nil
}

func (m *AStarMap) getHScore(n, target Point) int {
	x := AbsInt(n.Col - target.Col)
	y := AbsInt(n.Row - target.Row)
	//// return x + y
	return 2 * (x + y)
	//return x*x + y*y
}

// AbsInt 取正数的绝对值
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 找 map 中 F 值最小的。
// 如果有两个同样小的，随机选择 map 是没有顺序的集合
func getMinFScore(openList *sync.Map) *PathPoint {
	var min *PathPoint = nil
	var minF int = 0

	openList.Range(func(k, v interface{}) bool {
		pp := v.(*PathPoint)
		f := pp.GetFScore()

		if min == nil || f < minF {
			min = pp
			minF = f
		}
		return true
	})
	// 删除
	if min != nil {
		openList.Delete(min.Point)
	}

	return min
}

// getNeighbors 获得 n 节点可以到达的邻居路径点
func (m *AStarMap) getNeighbors(n Point) []*PathPoint {
	// 这个地图只能往 上下左右四个方向移动。
	arr := []*PathPoint{}
	// 左
	ny, nx := n.Row-1, n.Col
	if ny >= 0 && ny < MapRows && nx >= 0 && nx < MapCols && m.nodeMap[ny][nx] == 0 {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}
	// 右
	ny, nx = n.Row+1, n.Col
	if ny >= 0 && ny < MapRows && nx >= 0 && nx < MapCols && m.nodeMap[ny][nx] == 0 {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}

	// 上
	ny, nx = n.Row, n.Col-1
	if ny >= 0 && ny < MapRows && nx >= 0 && nx < MapCols && m.nodeMap[ny][nx] == 0 {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}

	// 下
	ny, nx = n.Row, n.Col+1
	if ny >= 0 && ny < MapRows && nx >= 0 && nx < MapCols && m.nodeMap[ny][nx] == 0 {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}

	return arr
}
