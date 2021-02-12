package main

import (
	_ "image/png"
	"log"

	"github.com/ghj1976/games/nerveincats"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	game *nerveincats.Game
)

func main() {

	game, _ = nerveincats.NewGame()

	ebiten.SetWindowSize(nerveincats.ScreenWidth, nerveincats.ScreenHeight)
	ebiten.SetWindowTitle("抓住神经猫")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
