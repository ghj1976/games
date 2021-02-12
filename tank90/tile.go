package tank90

import (
	"image/color"

	"github.com/ghj1976/games/tank90/images"
	"github.com/ghj1976/games/tank90/tools"
	"github.com/hajimehoshi/ebiten/v2"
)

// Tile 地砖，地图的每一块
type Tile struct {
	TileTypeID int  // 瓷砖类型， 地图上的标示数字
	MapX       int  // 地图上的位置
	MapY       int  // 地图上的位置
	Area       int  // 墙剩余的部分信息，用二进制数据的后四位来保存
	NextIMG    bool // 是否要使用第二张图
	cx         int  // 图像的中心点，屏幕坐标点，用于碰撞检测用
	cy         int  // 图像的中心点，屏幕坐标点，用于碰撞检测用
	width      int  // 宽度，用于碰撞检测用
	height     int  // 高度，用于碰撞检测用
}

// NewTile 构造
// 默认地图元素全显示
func NewTile(ttid, mapx, mapy int) *Tile {
	t := &Tile{}
	t.TileTypeID = ttid
	t.NextIMG = false
	t.MapX = mapx
	t.MapY = mapy
	t.Area = 0b1111 // 默认一定全有

	// 碰撞判断需要的数据
	t.cx = images.TileSize*mapx + images.TileSize/2
	t.cy = images.TileSize*mapy + images.TileSize/2
	t.width = images.TileSize
	t.height = images.TileSize
	return t
}

// GetCentorPositionAndSize 获得中心点位置及长宽, 碰撞判断用
func (t *Tile) GetCentorPositionAndSize() (x, y, w, h int) {
	return t.cx, t.cy, t.width, t.height
}

// Draw 画出墙
func (t *Tile) Draw(mapImage *ebiten.Image) {

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(t.cx-t.width/2), float64(t.cy-t.height/2))
	mapImage.DrawImage(images.GetMapTileImage(t.TileTypeID, t.Area, t.NextIMG), opts2) // 游戏地图区域绘制

	// txt := fmt.Sprintf("%d,%d", t.MapX, t.MapY)
	// text.Draw(mapImage, txt, settingFace, t.CX-t.Width/2+1, t.CY-t.Height/2+3, color.RGBA{0, 0, 205, 255})
}

// DrawDebugInfo 画出调试信息
func DrawDebugInfo(mapImage *ebiten.Image, mapx, mapy int, colour color.Color) {

	mazeImage := ebiten.NewImage(images.TileSize-2, images.TileSize-2)
	mazeImage.Fill(colour)

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(mapx*images.TileSize+1.0), float64(mapy*images.TileSize+1.0))

	mapImage.DrawImage(mazeImage, opts2)
}

// Attacked 墙被攻击的数据处理
// 这个的前提是发生了碰撞
// 要确保只调用一次
// 返回值 bool， 这堵墙是不是要完全销毁，  true 完全销毁， false 部分摧毁。 完全销毁要在外部map中删除它。
func (t *Tile) Attacked(b *Bullet) bool {

	from := ""
	switch b.Towards {
	case "top": // 子弹往上飞，攻击了bottom
		from = "bottom"
	case "bottom":
		from = "top"
	case "left":
		from = "right"
	case "right":
		from = "left"
	default:
		// 参数异常
	}
	t.Area = tools.BrickAttacked(t.Area, b.Power, from)

	// TODO 重算 墙的 碰撞属性  中心点， 长宽

	b.IsFinish = true
	if t.Area == 0 {
		// 被完全销毁了
		return true
	}

	// 还有残留
	return false
}
