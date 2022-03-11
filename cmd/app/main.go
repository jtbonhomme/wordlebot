package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jtbonhomme/wordlebot/internal/app"
)

func main() {
	game, err := app.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(app.ScreenWidth, app.ScreenHeight)
	ebiten.SetWindowTitle("WordleBot")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
