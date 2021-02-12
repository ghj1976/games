package tank90

import (
	"github.com/ghj1976/games/tank90/images"
	"github.com/hajimehoshi/ebiten/v2"
)

// Bird 总部
type Bird struct {
	X0     int // 总部 左上角所在的坐标 X （相对地图数组的坐标）
	Y0     int // 总部 左上角所在的坐标 Y（相对地图数组的坐标）
	cx     int // 图像的中心点，屏幕坐标点，用于碰撞检测用
	cy     int // 图像的中心点，屏幕坐标点，用于碰撞检测用
	width  int // 宽度，用于碰撞检测用
	height int // 高度，用于碰撞检测用
}

// NewBird 新建总部对象
func NewBird(x0, y0 int) *Bird {
	b := &Bird{}
	// 绘图需要的数据
	b.X0 = x0
	b.Y0 = y0
	// log.Println(fmt.Sprintf("bird x0:%d y0:%d", x0, y0))

	// 碰撞判断需要的数据
	b.cx = images.TileSize * (x0 + 1)
	b.cy = images.TileSize * (y0 + 1)
	b.width = images.TankSize
	b.height = images.TankSize

	return b
}

// GetCentorPositionAndSize 获得中心点位置及长宽, 碰撞判断用
func (b *Bird) GetCentorPositionAndSize() (x, y, w, h int) {
	return b.cx, b.cy, b.width, b.height
}

// Draw 画出墙
func (b *Bird) Draw(mapImage *ebiten.Image) {

	opts2 := &ebiten.DrawImageOptions{}
	opts2.GeoM.Translate(float64(b.cx-b.width/2), float64(b.cy-b.height/2))
	mapImage.DrawImage(images.GetBirdImage(false), opts2) // 游戏地图区域绘制
}
