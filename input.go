package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Input struct {
	up, down, left, right bool
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() {
	i.up = ebiten.IsKeyPressed(ebiten.KeyUp)
	i.down = ebiten.IsKeyPressed(ebiten.KeyDown)
	i.left = ebiten.IsKeyPressed(ebiten.KeyLeft)
	i.right = ebiten.IsKeyPressed(ebiten.KeyRight)
}

func (i *Input) Up() bool {
	return i.up
}

func (i *Input) Down() bool {
	return i.down
}

func (i *Input) Left() bool {
	return i.left
}

func (i *Input) Right() bool {
	return i.right
}
