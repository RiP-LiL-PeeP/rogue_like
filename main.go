package main

import (
	"log"
	"sara_quest/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := core.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Side-Scroller")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
