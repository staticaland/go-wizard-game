package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	player  *Player
	bullets *Bullets
}

func NewGame() *Game {
	player := NewPlayer()

	return &Game{
		player:  player,
		bullets: &Bullets{},
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.player.Update(g.bullets)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen, g.bullets)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
