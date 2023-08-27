package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	circle  *Circle
	bullets *Bullets
}

func NewGame() *Game {
	circle := NewCircle()

	return &Game{
		circle:  circle,
		bullets: &Bullets{},
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.circle.Update(g.bullets)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.circle.Draw(screen, g.bullets)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
