package main

import (
	_ "image/png"
	"log"

	"github.com/ghj1976/games/tank90"
	"github.com/ghj1976/games/tank90/navigation"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(tank90.ScreenWidth*2, tank90.ScreenHeight*2)
	ebiten.SetWindowTitle("坦克大战")

	gc := navigation.GetGameController()
	if err := ebiten.RunGame(gc); err != nil {
		log.Fatal(err)
	}
}
