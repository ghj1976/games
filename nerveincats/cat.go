package nerveincats

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Cat 神经猫
type Cat struct {
	IsSurrounded bool // 是不是已经被围住了，显示的图将不一样
	Q            int  // 所在位置 q
	R            int  // 所在位置 R

	count int

	// 普通状态的图
	cat11Image *ebiten.Image
	cat12Image *ebiten.Image
	cat13Image *ebiten.Image

	// 被围住状态的
	cat21Image *ebiten.Image
	cat22Image *ebiten.Image
	cat23Image *ebiten.Image
	cat24Image *ebiten.Image
	cat25Image *ebiten.Image

	frameWidth  int
	frameHeight int
}

// NewCat 初始化猫对象
// 包括图片的提前加载
func NewCat(catImage *ebiten.Image) *Cat {
	c := &Cat{IsSurrounded: false}
	c.Q = 0
	c.R = 0

	c.count = 0
	c.frameWidth = 60
	c.frameHeight = 100

	c.cat11Image = catImage.SubImage(image.Rect(0, 0, 60, 100)).(*ebiten.Image)
	c.cat12Image = catImage.SubImage(image.Rect(60, 0, 120, 100)).(*ebiten.Image)
	c.cat13Image = catImage.SubImage(image.Rect(120, 0, 180, 100)).(*ebiten.Image)

	c.cat21Image = catImage.SubImage(image.Rect(0, 100, 70, 200)).(*ebiten.Image)
	c.cat22Image = catImage.SubImage(image.Rect(70, 100, 140, 200)).(*ebiten.Image)
	c.cat23Image = catImage.SubImage(image.Rect(140, 100, 210, 200)).(*ebiten.Image)
	c.cat24Image = catImage.SubImage(image.Rect(210, 100, 280, 200)).(*ebiten.Image)
	c.cat25Image = catImage.SubImage(image.Rect(280, 100, 350, 200)).(*ebiten.Image)

	return c
}

// Surrounded 设置猫被围住了
func (c *Cat) Surrounded() {
	c.IsSurrounded = true
	c.frameWidth = 70
}

// GetKey 获得猫目前所在位置对应的地砖key
func (c *Cat) GetKey() string {
	return getKey(c.Q, c.R)
}

// Draw 在地图上绘制猫的图像
func (c *Cat) Draw(mapImage *ebiten.Image) {
	if ebiten.IsDrawingSkipped() {
		return
	}

	c.count++
	if c.count > 2147483640 {
		c.count = 0
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(c.frameWidth)/2+float64(tileWidth)/2, -float64(c.frameHeight)/2)

	x, y := hexToPixel(c.Q, c.R)
	op.GeoM.Translate(x, y)

	if c.IsSurrounded {
		// 被围住
		switch (c.count / 10) % 5 {
		case 0:
			mapImage.DrawImage(c.cat21Image, op)
		case 1:
			mapImage.DrawImage(c.cat22Image, op)
		case 2:
			mapImage.DrawImage(c.cat23Image, op)
		case 3:
			mapImage.DrawImage(c.cat24Image, op)
		case 4:
			mapImage.DrawImage(c.cat25Image, op)
		default:
			mapImage.DrawImage(c.cat21Image, op)
		}
	} else { // 还未被围住
		i := ((c.count / 6) % 8)
		if i > 0 && i <= 3 {
			mapImage.DrawImage(c.cat11Image, op)
		} else if i > 4 {
			mapImage.DrawImage(c.cat13Image, op)
		} else {
			mapImage.DrawImage(c.cat12Image, op)
		}
	}

}

// Reset 复原到原始状态，重新玩一盘
func (c *Cat) Reset() {
	c.count = 0
	c.Q = 0
	c.R = 0
	c.frameWidth = 60
	c.IsSurrounded = false
}
