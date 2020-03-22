package main

import (
	_ "image/png"
	"log"

	"github.com/ghj1976/games/tank90"
	"github.com/hajimehoshi/ebiten"
)

var (
	game *tank90.Game
)

func main() {

	game, _ = tank90.NewGame()

	if err := ebiten.Run(update, tank90.ScreenWidth, tank90.ScreenHeight, 2, "坦克大战"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	game.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	game.Draw(screen)

	return nil
}
