package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	aMap *AStarMap
)

const (
	ScreenWidth  = 380
	ScreenHeight = 380
	MapSize      = 352
)

func main() {
	prepareImage()
	aMap = prepareData()

	go aMap.FindPath(Point{Row: 0, Col: 0}, Point{Row: 49, Col: 49})
	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 1, "A*寻路算法演示"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {

	aMap.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Fill(color.RGBA{132, 132, 132, 255})
	square, _ := ebiten.NewImage(MapSize, MapSize, ebiten.FilterNearest)
	square.Fill(color.Black)

	aMap.Draw(square)

	// 把map区域画在地图上。
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(ScreenWidth/2-MapSize/2, ScreenHeight/2-MapSize/2)
	screen.DrawImage(square, opts) // 游戏地图区域绘制

	return nil
}
