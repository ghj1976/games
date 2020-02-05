package nerveincats

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Tile 瓷砖 ，地图上每一个节点
//
type Tile struct {
	Column     int     // 列，对应的x轴
	Row        int     // 行， 对应的z轴
	X          float64 // 实际绘图的基准位置，保存起来，避免每次计算
	Y          float64 // 实际绘图的基准位置，保存起来，避免每次计算
	IsBorder   bool    // 这个节点是不是地图的边缘
	IsObstacle bool    // 是否是障碍物，默认false，不是
	Rank       int     // 用于寻路的权重，当没有被围住时，算的时最短路径，围住时算的是最大通路。
}

// NewTile 初始化一个地砖
func NewTile(q, r int, obstacle bool) *Tile {
	t := &Tile{}
	t.Column = q
	t.Row = r
	t.IsObstacle = obstacle
	t.Rank = 1
	return t
}

var (
	tile0Image *ebiten.Image // 默认的地砖图片
	tile1Image *ebiten.Image // 已经有障碍物的地砖图片
	tileWidth  = 52
	tileHeight = 60

	// TileSize 每个地图区块的半径，六边形每边的边长。
	TileSize = 30.0
)

// InitTileIMG 初始化地砖图片
// 考虑到地砖对象会多次创建，这个方案不是类的方案
func InitTileIMG(catImage *ebiten.Image) {
	// 地砖图像是是一个  长 52  高 60 的图像， 根号3，2 的长宽比。
	tile0Image = catImage.SubImage(image.Rect(204, 30, 256, 90)).(*ebiten.Image)
	tile1Image = catImage.SubImage(image.Rect(274, 30, 326, 90)).(*ebiten.Image)
}

func getKey(q, r int) string {
	return fmt.Sprintf("q%d,r%d", q, r)
}

// GetKey 得到唯一识别码
func (t *Tile) GetKey() string {
	return getKey(t.Column, t.Row)
}

// Draw 绘图
func (t *Tile) Draw(mapImage *ebiten.Image) {
	if ebiten.IsDrawingSkipped() {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(t.X, t.Y)

	if t.IsObstacle {
		mapImage.DrawImage(tile1Image, op)
	} else {
		mapImage.DrawImage(tile0Image, op)
	}
}

// Reset 复原到原始状态，重新玩一盘
func (t *Tile) Reset() {
	t.IsObstacle = false
	t.Rank = 1
}

// Obstacle 把指定瓷砖设置成障碍物
func (t *Tile) Obstacle() {
	t.IsObstacle = true
	t.Rank = 500
}
