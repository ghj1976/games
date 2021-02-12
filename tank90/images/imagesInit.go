package images

import (
	"bytes"
	"image"
	"log"

	"github.com/ghj1976/games/tank90/tools"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// TileSize 地图贴图的最小单位
	TileSize = 8
	// TankSize 坦克的尺寸
	TankSize = 16
)

var (
	tilesImage *ebiten.Image // 地砖图片
	tanksImage *ebiten.Image // take的图片
	itemsImage *ebiten.Image // 游戏元素图片
)

// InitGameIMG 加载资源文件
// 考虑到对象会多次创建，这个方案不是类的方案
func InitGameIMG() {
	// 地图地形贴图
	img, _, err := image.Decode(bytes.NewReader(Tiles_png))
	if err != nil {
		log.Println(" Tiles_png 资源文件解析错误。")
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	// 坦克 贴图
	img, _, err = image.Decode(bytes.NewReader(Tanks_png))
	if err != nil {
		log.Println(" Bird_png 资源文件解析错误。")
		log.Fatal(err)
	}
	tanksImage = ebiten.NewImageFromImage(img)

	// 游戏元素图片
	img, _, err = image.Decode(bytes.NewReader(Items_png))
	if err != nil {
		log.Println(" Enemy_1_png 资源文件解析错误。")
		log.Fatal(err)
	}
	itemsImage = ebiten.NewImageFromImage(img)

}

// GetMapTileImage 获得游戏地砖图像
// typeid 地图上的类型ID，  0 道路  1 砖墙 2 森林 3 水域 4 沙漠  5 石墙
// next 对水域才有意义， 如果是 true 下一张图  false 默认第一张图
// area 位运算的数字，标示取图片的区域
// 四位的二进制数字 依次是 top bottom left right 上下左右。 0 标示没有， 1标示有。
// 1111 完整的图，
// topleft     topright
// bottomleft  bottomright
func GetMapTileImage(typeid, area int, next bool) *ebiten.Image {
	if typeid < 0 && typeid > 5 {
		// 不支持的地图类型
		log.Fatalf("不支持的地图类型 %d", typeid)
	}

	offset := 0
	if next { // 第二幅图
		offset = TileSize
	}

	if area == 0 {
		return nil // 全部被销毁了，不应该有图。
	}
	top, bottom, left, right := tools.GetBrickArea(area)

	x0 := offset + TileSize/2*(left-1)
	y0 := TileSize*typeid + TileSize/2*(top-1)
	x1 := offset + TileSize/2*right
	y1 := TileSize*typeid + TileSize/2*bottom

	return tilesImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}

// GetBirdImage 获得鹰巢的图
// failure 是否失败
func GetBirdImage(failure bool) *ebiten.Image {
	offset := 0
	if failure { // 失败的图
		offset = TileSize * 2
	}
	x0 := offset
	y0 := TileSize * 2 * 9
	x1 := offset + TileSize*2
	y1 := TileSize * 2 * 10
	return tilesImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}

// GetTankImage 获取坦克的图片
// num 编号， 从1到8
// color 颜色   red,green,golden,silver
// toward 朝向 top left bottom right
// move 是否移动，移动是第二幅图
func GetTankImage(num int, color, toward string, move bool) *ebiten.Image {

	if num < 0 || num > 8 {
		log.Fatalf("没有这么多种坦克 %d", num)
	}

	// 是否移动的偏移
	moveOffset := 0
	if move {
		moveOffset = TankSize
	}

	// 朝向的偏移
	towardOffset := 0
	switch toward {
	case "top":
		towardOffset = 0
	case "left":
		towardOffset = TankSize * 2
	case "bottom":
		towardOffset = TankSize * 4
	case "right":
		towardOffset = TankSize * 6
	}

	// 颜色偏移
	colorOffsetX, colorOffsetY := 0, 0
	switch color {
	case "golden":
		colorOffsetX, colorOffsetY = 0, 0
	case "silver":
		colorOffsetX, colorOffsetY = TankSize*8, 0
	case "green":
		colorOffsetX, colorOffsetY = 0, TankSize*8
	case "red":
		colorOffsetX, colorOffsetY = TankSize*8, TankSize*8
	}

	x0 := 0 + moveOffset + towardOffset + colorOffsetX
	y0 := TankSize*(num-1) + colorOffsetY
	x1 := TankSize + moveOffset + towardOffset + colorOffsetX
	y1 := TankSize*(num) + colorOffsetY
	return tanksImage.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}
