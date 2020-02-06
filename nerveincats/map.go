package nerveincats

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Map 游戏地图
type Map struct {
	TileSet map[string]*Tile // 地图
}

var (
	// Sqrt3 根号3的常量值
	Sqrt3     float64     // = math.Sqrt(3.0)
	mapRadius int     = 5 // 新建地图的半径
)

func init() {
	Sqrt3 = math.Sqrt(3.0)
}

// NewMap 新构建一个地图
func NewMap() *Map {
	m := &Map{}
	m.TileSet = make(map[string]*Tile)

	// 按照  https://indienova.com/indie-game-development/grid-implementation-of-hex/ 这里的思路
	// 生成 六边形的地图
	for q := -mapRadius; q <= mapRadius; q++ {
		r1 := MaxInt(-mapRadius, -q-mapRadius)
		r2 := MinInt(mapRadius, -q+mapRadius)
		log.Println(fmt.Sprintf("q:%d,r1:%d,r2:%d", q, r1, r2))
		for r := r1; r <= r2; r++ {
			t := NewTile(q, r, false)
			if r == r1 || r == r2 || q == mapRadius || q == -mapRadius {
				t.IsBorder = true // 地图的边缘节点
				t.Rank = 0        // 边缘节点不考虑 障碍物，权重都是 0
			} else {
				t.Rank = 1 // 默认权重是1
			}
			t.X, t.Y = hexToPixel(t.Column, t.Row)
			m.TileSet[t.GetKey()] = t
		}
	}
	m.RandomObstacle() // 障碍物布置
	return m
}

// 把六角形坐标系转换成像素坐标系
// 地砖的位置
func hexToPixel(q, r int) (x, y float64) {
	x = ScreenWidth/2 + TileSize*Sqrt3*(float64(q)+float64(r)/2) - float64(tileWidth)/2
	y = ScreenHeight/2 + TileSize*3.0/2.0*float64(r)
	return x, y
}

// 把点击的 像素坐标系 变成 六角形坐标系
// 专门用于地砖的定位
func pixelToHex(x, y int) (q, r int) {
	// 计算出地图中心点的位置，圆心位置
	fcx, fcy := hexToPixel(0, 0)
	cx := int(fcx) + tileWidth/2
	cy := int(fcy) + tileHeight/2

	// 从屏幕坐标，变成中心点坐标，
	x1 := x - cx
	y1 := y - cy

	r = int(math.Round(float64(y1) * 2.0 / 3.0 / float64(TileSize)))
	q = int(math.Round((float64(x1)/Sqrt3 - float64(y1)/3.0) / float64(TileSize)))
	return q, r
}

// CalculateTileRank 重新计算每个位置的权重
func (m *Map) CalculateTileRank() {
	// 对于没有围住的点，计算 最小路径
	// 对于围住的点，计算 最大通路

	preRankTileArr := []string{}

	// rank 值 是 0 和 500 的值， 在计算时，是不用考虑变化的。 只要计算其他的值。
	// 找到所有的权重是0的点，做为开始计算的初始点。
	// 同时需要计算权重的点，全部初始化权重成 -1， 方便后面判断要计算。
	for k, t := range m.TileSet {
		if t.Rank == 0 {
			preRankTileArr = append(preRankTileArr, k)
		} else if t.Rank == 500 {
			continue
		} else {
			// 把权重初始化
			m.TileSet[k].Rank = -1
		}
	}

	currRank := 0
	// 找跟它临近的，不是障碍物，权重待计算的，权重加一  最小路径算法
	for len(preRankTileArr) > 0 {

		currRank++
		currRankTileArr := []string{}

		for _, k := range preRankTileArr {
			q := m.TileSet[k].Column
			r := m.TileSet[k].Row
			if m.getTileRank(q+1, r) == -1 {
				key := getKey(q+1, r)
				m.TileSet[key].Rank = currRank
				currRankTileArr = append(currRankTileArr, key)
			}
			if m.getTileRank(q, r+1) == -1 {
				key := getKey(q, r+1)
				m.TileSet[key].Rank = currRank
				currRankTileArr = append(currRankTileArr, key)
			}
			if m.getTileRank(q-1, r+1) == -1 {
				key := getKey(q-1, r+1)
				m.TileSet[key].Rank = currRank
				currRankTileArr = append(currRankTileArr, key)
			}
			if m.getTileRank(q-1, r) == -1 {
				key := getKey(q-1, r)
				m.TileSet[key].Rank = currRank
				currRankTileArr = append(currRankTileArr, key)
			}
			if m.getTileRank(q, r-1) == -1 {
				key := getKey(q, r-1)
				m.TileSet[key].Rank = currRank
				currRankTileArr = append(currRankTileArr, key)
			}
			if m.getTileRank(q+1, r-1) == -1 {
				key := getKey(q+1, r-1)
				m.TileSet[key].Rank = currRank
				currRankTileArr = append(currRankTileArr, key)
			}
		}

		preRankTileArr = currRankTileArr
	}

	// 最大路径算法
	nullRankTileArr := []string{}
	for k, t := range m.TileSet {
		if t.Rank == -1 {
			nullRankTileArr = append(nullRankTileArr, k)
		}
	}
	for _, tk := range nullRankTileArr {
		// 被包围的权重是负数， 这样永远往最小值方向移动，就是最佳算法。
		m.TileSet[tk].Rank = 0 - m.getNeighborNum(m.TileSet[tk].Column, m.TileSet[tk].Row)
	}

}

// getTileRank 获得指定位置的 rank 值
func (m *Map) getTileRank(q, r int) int {
	t, b := m.TileSet[getKey(q, r)]
	if !b {
		// 如果在地图之外，返回 -100
		return -100
	}
	return t.Rank
}

// 获得周围可以移动的邻居数量
// 最大路径算法用
func (m *Map) getNeighborNum(q, r int) int {
	count := 0
	if m.GetTileStatus(q+1, r) == Normal {
		count++
	}
	if m.GetTileStatus(q, r+1) == Normal {
		count++
	}
	if m.GetTileStatus(q-1, r+1) == Normal {
		count++
	}
	if m.GetTileStatus(q-1, r) == Normal {
		count++
	}
	if m.GetTileStatus(q, r-1) == Normal {
		count++
	}
	if m.GetTileStatus(q+1, r-1) == Normal {
		count++
	}
	return count
}

// Draw 在地图上绘制猫的图像
func (m *Map) Draw(screen *ebiten.Image) {
	if ebiten.IsDrawingSkipped() {
		return
	}
	screen.Fill(color.RGBA{102, 102, 102, 255})

	for _, t := range m.TileSet {
		t.Draw(screen)
	}

}

// RandomObstacle 随机给地图障碍物
func (m *Map) RandomObstacle() {

	// key 字符串数组
	keyArr := make([]string, len(m.TileSet))
	i := 0
	for k := range m.TileSet {
		keyArr[i] = k
		i++
	}

	// 最少给 10个， 最多给 16个。
	for i := 0; i < 10+rand.Intn(6); i++ {
		rr := rand.Intn(len(keyArr))
		m.TileSet[keyArr[rr]].Obstacle()
	}

	// 中心节点一定不能障碍物，这是游戏的起点。
	m.TileSet[m.GetCenterKey()].Reset()
}

// GetCenterKey 获得中心点位置的key
func (m *Map) GetCenterKey() string {
	return getKey(0, 0)
}

// CatRandomMove 猫随机移动
// 返回的是true，标示已经到边了
func (m *Map) CatRandomMove(c *Cat) GameStatus {

	if m.TileSet[c.GetKey()].IsBorder {
		// 游戏该结束了，到了地图边缘，猫成功
		return Failure
	}

	rankMap := map[string]int{}
	r := m.getTileRank(c.Q+1, c.R)
	if r > -100 && r < 500 {
		rankMap[getKey(c.Q+1, c.R)] = r
	}
	r = m.getTileRank(c.Q, c.R+1)
	if r > -100 && r < 500 {
		rankMap[getKey(c.Q, c.R+1)] = r
	}
	r = m.getTileRank(c.Q-1, c.R+1)
	if r > -100 && r < 500 {
		rankMap[getKey(c.Q-1, c.R+1)] = r
	}
	r = m.getTileRank(c.Q-1, c.R)
	if r > -100 && r < 500 {
		rankMap[getKey(c.Q-1, c.R)] = r
	}
	r = m.getTileRank(c.Q, c.R-1)
	if r > -100 && r < 500 {
		rankMap[getKey(c.Q, c.R-1)] = r
	}
	r = m.getTileRank(c.Q+1, c.R-1)
	if r > -100 && r < 500 {
		rankMap[getKey(c.Q+1, c.R-1)] = r
	}
	if len(rankMap) == 0 {
		// 被完全包围了，游戏结束
		return Success
	}

	mink := ""
	minr := 500
	for k, r := range rankMap {
		if minr > r {
			mink = k
			minr = r
		}
	}

	// 选择最小移动方向
	c.Q = m.TileSet[mink].Column
	c.R = m.TileSet[mink].Row

	if m.TileSet[mink].Rank < 0 {
		c.Surrounded()
		return Surrounded
	}
	return Processing
}

// TileStatus 地砖的三种状态
type TileStatus int

const (
	// Normal 0 正常
	Normal TileStatus = iota
	// Obstacle 1 障碍物
	Obstacle
	// Outside 2 地图之外的位置
	Outside
)

// GetTileStatus 获得指定位置的状态
// 一共有三种状态， 障碍物、地图之外、正常
func (m *Map) GetTileStatus(q, r int) TileStatus {

	t, b := m.TileSet[getKey(q, r)]
	if !b {
		return Outside
	}
	if t.IsObstacle {
		return Obstacle
	}
	return Normal
}

// Reset 复原到原始状态，重新玩一盘
func (m *Map) Reset() {
	for k := range m.TileSet {
		m.TileSet[k].Reset()
	}

	// 六边形的地图 的 rank 值重新赋值
	for q := -mapRadius; q <= mapRadius; q++ {
		r1 := MaxInt(-mapRadius, -q-mapRadius)
		r2 := MinInt(mapRadius, -q+mapRadius)
		for r := r1; r <= r2; r++ {
			key := getKey(q, r)
			if r == r1 || r == r2 {
				m.TileSet[key].IsBorder = true // 地图的边缘节点
				m.TileSet[key].Rank = 0        // 边缘节点不考虑 障碍物，权重都是 0
			} else {
				m.TileSet[key].Rank = 1 // 默认权重是1
			}
		}
	}

	m.RandomObstacle() // 障碍物布置
}
