package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	game := NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("2D Game in Golang with Ebiten")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
