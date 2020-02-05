package main

import (
	_ "image/png"
	"log"

	"github.com/ghj1976/games/nerveincats"
	"github.com/hajimehoshi/ebiten"
)

var (
	game *nerveincats.Game
)

func update(screen *ebiten.Image) error {
	game.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	game.Draw(screen)

	return nil
}

func main() {

	game, _ = nerveincats.NewGame()

	if err := ebiten.Run(update, nerveincats.ScreenWidth, nerveincats.ScreenHeight, 1, "抓住神经猫"); err != nil {
		log.Fatal(err)
	}
}
