package main

import (
	_ "image/png"
	"log"

	"github.com/ghj1976/games/tank90"
	"github.com/ghj1976/games/tank90/navigation"
	"github.com/hajimehoshi/ebiten"
)

var (
	gc *navigation.GameController
)

func main() {

	gc = navigation.GetGameController()

	if err := ebiten.Run(update, tank90.ScreenWidth, tank90.ScreenHeight, 2, "坦克大战"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	gc.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}
	gc.Draw(screen)

	return nil
}
