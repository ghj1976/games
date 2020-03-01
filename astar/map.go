package main

/*
这个文件放个绘图有关
*/

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	// Size 地图的尺寸
	Size = 50
)

var (
	whiteSquare *ebiten.Image
	redSquare   *ebiten.Image
	greenSquare *ebiten.Image
	goldSquare  *ebiten.Image
)

func prepareImage() {

	whiteSquare, _ = ebiten.NewImage(5, 5, ebiten.FilterDefault)
	whiteSquare.Fill(color.White)
	redSquare, _ = ebiten.NewImage(5, 5, ebiten.FilterDefault)
	redSquare.Fill(color.RGBA{255, 0, 255, 255})
	greenSquare, _ = ebiten.NewImage(5, 5, ebiten.FilterDefault)
	greenSquare.Fill(color.RGBA{60, 179, 113, 255})
	goldSquare, _ = ebiten.NewImage(5, 5, ebiten.FilterDefault)
	goldSquare.Fill(color.RGBA{255, 215, 0, 255})
}

// Draw 绘图
func (m *AStarMap) Draw(screen *ebiten.Image) {
	// 地图布局绘制
	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			if m.nodeMap[y][x] == 0 {
				drawNode(screen, x, y, "white")
			}

		}
	}
	// open 列表的数据绘制
	m.openList.Range(func(k, _ interface{}) bool {
		p := k.(Point)
		drawNode(screen, p.Col, p.Row, "green")
		return true
	})

	// close 列表的数据绘制
	m.closeList.Range(func(k, _ interface{}) bool {
		p := k.(Point)
		drawNode(screen, p.Col, p.Row, "red")
		return true
	})

	// 目前探索出来的最佳路径绘制
	n := m.currN
	for n != nil {
		drawNode(screen, n.Col, n.Row, "gold")
		n = n.Parent
	}

}

func drawNode(screen *ebiten.Image, x, y int, color string) {
	op := &ebiten.DrawImageOptions{}
	x0 := 2 + x*7
	y0 := 2 + y*7
	op.GeoM.Translate(float64(x0), float64(y0))

	switch color {
	case "white":
		screen.DrawImage(whiteSquare, op)
	case "red":
		screen.DrawImage(redSquare, op)
	case "green":
		screen.DrawImage(greenSquare, op)
	case "gold":
		screen.DrawImage(goldSquare, op)
	}

}

// Update 绘图前更新状态
func (m *AStarMap) Update() {

}
