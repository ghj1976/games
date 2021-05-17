package main

import (
	"image/color"
	"log"

	"github.com/blizzy78/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ScreenWidth, ScreenHeight       int
	CurrDemoMapType, CurrHScoreType int
	CurrGame                        *Game
)

type Game struct {
	aMap      *AStarMap    // 游戏区域，地图区域
	controlUI *ebitenui.UI // 控制菜单区域
}

// Update 逻辑刷新
func (g *Game) Update() error {
	g.aMap.Update()
	g.controlUI.Update()
	return nil
}

// Draw 绘图
func (g *Game) Draw(screen *ebiten.Image) {

	// if ebiten.IsDrawingSkipped() {
	// 	return nil
	// }

	// 整体背景颜色
	screen.Fill(color.RGBA{132, 132, 132, 255})

	// map背景颜色
	square := ebiten.NewImage(g.aMap.MapWidth, g.aMap.MapHeight)
	square.Fill(color.Black)
	g.aMap.Draw(square)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(ScreenWidth)-float64(g.aMap.MapWidth)-10.0, float64(ScreenHeight)/2.0-float64(g.aMap.MapHeight)/2.0)
	// 把map区域画在地图上。
	screen.DrawImage(square, opts) // 游戏地图区域绘制

	g.controlUI.Draw(screen)
}

// Layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// return ScreenWidth, ScreenHeight
	return outsideWidth, outsideHeight
}

func main() {
	CurrDemoMapType = 1
	CurrHScoreType = 2
	prepareGameImage()
	CurrGame = &Game{}
	CurrGame.aMap = Prepare(CurrDemoMapType, CurrHScoreType)

	// 控制界面
	ui, closeUI, err := createUI()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("100")
	CurrGame.controlUI = ui
	defer closeUI()

	// 开始游戏
	CurrGame.aMap.Reset(CurrDemoMapType, CurrHScoreType)

	// 窗口大小
	// ScreenWidth = (g.aMap.MapWidth/100+1)*100 + 300
	// ScreenHeight = (g.aMap.MapHeight/100+1)*100 + 30
	ScreenWidth = 600
	ScreenHeight = 400

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("A*寻路算法演示")
	ebiten.SetScreenClearedEveryFrame(false)

	if err := ebiten.RunGame(CurrGame); err != nil {
		log.Fatal(err)
	}
}
