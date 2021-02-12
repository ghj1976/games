package tools

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ICollision 检查碰撞的接口
type ICollision interface {
	GetCentorPositionAndSize() (x, y, w, h int)
}

// CheckCollision 检查是否碰撞
func CheckCollision(co1, co2 ICollision) bool {
	x1, y1, w1, h1 := co1.GetCentorPositionAndSize()
	x2, y2, w2, h2 := co2.GetCentorPositionAndSize()
	if AbsInt(x1-x2) < w1/2+w2/2 && AbsInt(y1-y2) < h1/2+h2/2 {
		return true
	}
	return false
}

// RectangleCollisionObject 矩形碰撞判断及绘图的对象
type RectangleCollisionObject struct {
	cx     int // 图像的中心点，屏幕坐标点，用于碰撞检测用
	cy     int // 图像的中心点，屏幕坐标点，用于碰撞检测用
	width  int // 宽度，用于碰撞检测用
	height int // 高度，用于碰撞检测用

}

// NewRectangleCollisionObject 初始化一个 矩形绘图及碰撞对象
func NewRectangleCollisionObject(x, y, width, height int) RectangleCollisionObject {
	rco := RectangleCollisionObject{}
	rco.cx = x
	rco.cy = y
	rco.width = width
	rco.height = height
	return rco
}

// GetCentorPosition 获得中心点位置
func (co *RectangleCollisionObject) GetCentorPosition() (x, y int) {
	return co.cx, co.cy
}

// GetCentorPositionAndSize 获得中心点位置及长宽
func (co *RectangleCollisionObject) GetCentorPositionAndSize() (x, y, w, h int) {
	return co.cx, co.cy, co.width, co.height
}

// GetLeftTopPoint 获得左上角点的位置
func (co *RectangleCollisionObject) GetLeftTopPoint() (x, y int) {
	x = co.cx - co.width/2
	y = co.cx - co.height/2
	return x, y
}

// GetLeftTopPointFloat64 获得左上角点的位置 浮点值
func (co *RectangleCollisionObject) GetLeftTopPointFloat64() (x, y float64) {
	ix, iy := co.GetLeftTopPoint()
	return float64(ix), float64(iy)
}

// GetDrawOpts 获得绘图位置
func (co *RectangleCollisionObject) GetDrawOpts() *ebiten.DrawImageOptions {
	opts2 := &ebiten.DrawImageOptions{}
	// opts2.GeoM.Translate(float64(co.CX-co.Width/2), float64(co.CY-co.Height/2))
	opts2.GeoM.Translate(co.GetLeftTopPointFloat64())
	return opts2
}

// CheckCollision 是否发生碰撞判断
func (co *RectangleCollisionObject) CheckCollision(next *RectangleCollisionObject) bool {
	if AbsInt(co.cx-next.cx) < co.width/2+next.width/2 && AbsInt(co.cy-next.cy) < co.height/2+next.height/2 {
		// log.Printf("%d,%d,%d,%d -- %d,%d,%d,%d", co.CX, co.CY, co.Width, co.Height, next.CX, next.CY, next.Width, next.Height)
		return true
	}
	return false
}
