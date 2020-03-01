package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	aMap *AStarMap
)

var (
	MapWidth     int
	MapHeight    int
	ScreenWidth  int
	ScreenHeight int
)

func main() {
	prepareImage()

	// 随机地图
	// rows, cols := 50, 50
	// aMap = prepareData1(rows, cols)

	// 坦克大战的地图
	rows, cols := 26, 26
	aMap = prepareDataTank90(rows, cols)

	MapWidth = 2 + (5+2)*cols
	MapHeight = 2 + (5+2)*rows
	ScreenWidth = (MapWidth/100 + 1) * 100
	ScreenHeight = (MapHeight/100 + 1) * 100

	go aMap.FindPath(Point{Row: 0, Col: 0}, Point{Row: rows - 1, Col: cols - 1})
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
	square, _ := ebiten.NewImage(MapWidth, MapHeight, ebiten.FilterNearest)
	square.Fill(color.Black)

	aMap.Draw(square)

	// 把map区域画在地图上。
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(ScreenWidth)/2.0-float64(MapWidth)/2.0, float64(ScreenHeight)/2.0-float64(MapHeight)/2.0)
	screen.DrawImage(square, opts) // 游戏地图区域绘制

	return nil
}
