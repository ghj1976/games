package nerveincats

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/ghj1976/games/nerveincats/resources"
	"github.com/hajimehoshi/ebiten"
)

const (
	// ScreenWidth 屏幕大小，宽度
	ScreenWidth = 640
	// ScreenHeight 屏幕大小，高度
	ScreenHeight = 640
)

// GameStatus 游戏状态
type GameStatus int

const (
	// Processing  进行中， 还没被包围
	Processing GameStatus = iota

	// Surrounded  猫被包围了，但是还可以动
	Surrounded

	// Success  游戏过关了，猫被完全限制行动了。
	Success

	// Failure  失败，猫跑掉了
	Failure
)

// Game 游戏逻辑
type Game struct {
	NCat   *Cat
	NMap   *Map
	Status GameStatus // 游戏当前所处的状态
	Step   int        // 一共几步完成。

	replayBtn          *ReplayBtn
	result             *Result
	currCursorPosition string // 当前处理的点击位置点
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{}
	g.Step = 0

	// 加载资源文件
	img, _, err := image.Decode(bytes.NewReader(resources.Cat_png))
	if err != nil {
		log.Println(" Cat_png 资源文件解析错误。")
		log.Fatal(err)
	}
	catImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	// 猫
	g.NCat = NewCat(catImage)

	// 地图
	InitTileIMG(catImage)
	g.NMap = NewMap()

	// 游戏结果提示及按钮部分
	g.result = NewResult()
	g.replayBtn = NewReplayBtn()

	return g, nil
}

// Draw 绘图
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{102, 102, 102, 255})

	g.NMap.Draw(screen)
	g.NCat.Draw(screen)
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("x:%d y:%d", x, y))

	if g.Status == Success || g.Status == Failure {
		txt := ""
		if g.Status == Success {
			txt += fmt.Sprintf("您用%d步抓住了神经猫。", g.Step)
		} else {
			txt += "你没有抓住神经猫！！"
		}
		g.result.Draw(screen, g.Status, txt)
		g.replayBtn.Draw(screen)
	}
}

// Update updates the current game state.
func (g *Game) Update() error {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

		// 按下鼠标左键，判断那个位置要变成障碍物
		x, y := ebiten.CursorPosition()
		ccp := fmt.Sprintf("x:%d y:%d", x, y)
		if g.currCursorPosition == ccp {
			// 当前位置已经处理过一边了，不用再次处理了。
			return nil
		}

		log.Println(fmt.Sprintf("x:%d y:%d", x, y))

		if g.Status == Success || g.Status == Failure {
			// 这里只要处理 重新玩的按钮逻辑
			if g.replayBtn.In(x, y) { // 重玩按钮被点击
				g.Reset()
			}
			g.currCursorPosition = ccp
			return nil
		}

		log.Println("update IsMouseButtonPressed")
		// 正常游戏的逻辑
		q, r := pixelToHex(x, y)
		key := getKey(q, r)
		log.Println(key)

		// 算出来的位置超过地图的大小
		if math.Abs(float64(q)) > float64(mapRadius) {
			g.currCursorPosition = ccp
			return nil
		}
		if math.Abs(float64(r)) > float64(mapRadius) {
			g.currCursorPosition = ccp
			return nil
		}

		ts, b := g.NMap.TileSet[key]
		if !b {
			g.currCursorPosition = ccp
			return nil
		}
		if !ts.IsObstacle {
			// 设置点的位置为障碍物
			g.NMap.TileSet[key].Obstacle()

			// 重新计算每个位置的权重
			g.NMap.CalculateTileRank()
			g.Step++
			// 猫随机移动一个位置
			g.Status = g.NMap.CatRandomMove(g.NCat)

		}

		g.currCursorPosition = ccp
	}

	return nil
}

// Reset 重新启动一盘游戏
func (g *Game) Reset() {
	g.Step = 0
	g.NMap.Reset()
	g.NCat.Reset()
	g.Status = Processing
	log.Println("-----reset-----")
}
