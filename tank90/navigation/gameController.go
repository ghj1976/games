package navigation

import (
	"sync"

	"github.com/ghj1976/games/tank90"
	"github.com/ghj1976/games/tank90/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

// GameController 全局游戏控制，页面切换
type GameController struct {
	currPage string
	game     *tank90.Game
	nav      *Navigation
}

var instance *GameController
var once sync.Once

// GetGameController 得到邮件控制单件类，这个类负责各个页面切换
// 单件模式参考： https://www.jianshu.com/p/d2fc1c998fc9
func GetGameController() *GameController {
	once.Do(func() {
		instance = &GameController{}
	})

	resources.InitAudio()
	resources.InitFontFace()

	instance.currPage = "game"
	instance.game, _ = tank90.NewGame()
	instance.nav = NewNavigation()
	return instance
}

// Update updates the current game state.
func (gc *GameController) Update() error {
	if gc.currPage == "nav" {
		gc.nav.Update()
	} else if gc.currPage == "game" {
		gc.game.Update()
	}
	return nil
}

// Draw 绘图
func (gc *GameController) Draw(screen *ebiten.Image) {
	if gc.currPage == "nav" {
		gc.nav.Draw(screen)
	} else if gc.currPage == "game" {
		gc.game.Draw(screen)
	}

}

func (g *GameController) Layout(outsideWidth, outsideHeight int) (int, int) {
	return tank90.ScreenWidth, tank90.ScreenHeight
}
