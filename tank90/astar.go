package tank90

import (
	"github.com/ghj1976/games/tank90/tools"
)

/*
A* 寻路相关代码
*/

// Point 节点
type Point struct {
	Row int // 行 Y轴的值
	Col int // 列 X轴的值
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

// getHScore 估算值 H
func getHScore(n, target Point) int {
	x := tools.AbsInt(n.Col - target.Col)
	y := tools.AbsInt(n.Row - target.Row)
	return 2 * (x + y)
}

// 找 map 中 F 值最小的。
// 如果有两个同样小的，随机选择 map 是没有顺序的集合
func popMinFScore(openList map[Point]*PathPoint) *PathPoint {
	var min *PathPoint = nil
	var minF int = 0

	for _, pp := range openList {
		f := pp.GetFScore()
		if min == nil || f < minF {
			min = pp
			minF = f
		}
	}

	// 删除
	if min != nil {
		delete(openList, min.Point)
	}

	return min
}

// getNeighbors 获得 n 节点可以到达的邻居路径点
func getNeighbors(mapArr [26][26]int, n Point, shipable bool) []*PathPoint {

	// 这个地图只能往 上下左右四个方向移动。
	arr := []*PathPoint{}
	// 左
	ny, nx := n.Row-1, n.Col
	if CheckTankStayable(mapArr, nx, ny, shipable) {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}
	// 右
	ny, nx = n.Row+1, n.Col
	if CheckTankStayable(mapArr, nx, ny, shipable) {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}

	// 上
	ny, nx = n.Row, n.Col-1
	if CheckTankStayable(mapArr, nx, ny, shipable) {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}

	// 下
	ny, nx = n.Row, n.Col+1
	if CheckTankStayable(mapArr, nx, ny, shipable) {
		arr = append(arr, &PathPoint{Point: Point{Col: nx, Row: ny}, Parent: nil, GScore: -1, HScore: -1})
	}

	return arr
}

// FindPath 通过A*算法寻找一个最短路径
// source, target 起点 和 终点
func FindPath(mapArr [26][26]int, source, target Point, shipable bool) *PathPoint {
	openList := make(map[Point]*PathPoint)
	closeList := make(map[Point]*PathPoint)

	// 算起点的h(s)
	sourcePathPoint := &PathPoint{Point: source, Parent: nil, GScore: 0}
	sourcePathPoint.HScore = getHScore(source, target)

	// 将起点放入OPEN表
	openList[source] = sourcePathPoint

	for {
		// 从OPEN表中取f(n)最小的节点n;
		n := popMinFScore(openList)
		if n == nil { // open 列表没有数据， 则退出
			break
		}

		if n.Point == target { // 找到目标节点了
			return n
		}

		// 遍历 n 节点的每个临近节点
		for _, x := range getNeighbors(mapArr, n.Point, shipable) {
			// 计算f(X);
			x.HScore = getHScore(x.Point, target)
			x.GScore = n.GScore + 1

			ox, oexist := openList[x.Point]
			if oexist { // X in OPEN
				if x.GetFScore() < ox.GetFScore() {
					// Open 中只存最小 F 值的信息。
					// 如果有多条路都可以到达 x 节点， 只存最小的那条
					x.Parent = n
					openList[x.Point] = x
				}
			}

			_, cexist := closeList[x.Point]
			if cexist { // X in CLOSE
				// 已经不用处理了， 继续下一个
				continue
			}

			if !oexist && !cexist { // X not in both
				// 把n设置为X的父亲
				x.Parent = n
				// 求f(X); 循环进入时已经处理了，这里不用处理
				// 并将X插入OPEN表中;//还没有排序
				openList[x.Point] = x

			}

		}

		//  将n节点插入CLOSE表中;
		closeList[n.Point] = n
		//  按照f(n)将OPEN表中的节点排序; //实际上是比较OPEN表内节点f的大小，从最小路径的节点向下进行。
		// 这里每次提取最小节点是循环找的，不用排序
	}
	return nil
}
