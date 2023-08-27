package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	circleRadius = 50
	circleSpeed  = 2
	bulletSize   = 4
)

type Bullet struct {
	x, y   float64
	vx, vy float64
}

func NewBullet(x, y, vx, vy float64) *Bullet {
	return &Bullet{
		x:  x,
		y:  y,
		vx: vx,
		vy: vy,
	}
}

func (b *Bullet) Update() {
	b.x += b.vx
	b.y += b.vy
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, b.x-bulletSize/2, b.y-bulletSize/2, bulletSize, bulletSize, color.RGBA{255, 255, 255, 255})
}

type Bullets []*Bullet

func (bs *Bullets) Update() {
	for _, b := range *bs {
		b.Update()
	}
}

func (bs *Bullets) Draw(screen *ebiten.Image) {
	for _, b := range *bs {
		b.Draw(screen)
	}
}

type Circle struct {
	x, y     float64
	lastShot int64
}

func NewCircle() *Circle {
	return &Circle{
		x:        400,
		y:        300,
		lastShot: time.Now().UnixNano(),
	}
}

func (c *Circle) Update(bullets *Bullets) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		c.y -= circleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		c.y += circleSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		c.x -= circleSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		c.x += circleSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && time.Now().UnixNano()-c.lastShot >= 1000000000 {
		*bullets = append(*bullets, NewBullet(c.x, c.y, 0, -2))
		c.lastShot = time.Now().UnixNano()
	}

	bullets.Update()
}

func (c *Circle) Draw(screen *ebiten.Image, bullets *Bullets) {
	ebitenutil.DrawRect(screen, c.x-circleRadius, c.y-circleRadius, circleRadius*2, circleRadius*2, color.RGBA{255, 0, 0, 255})
	bullets.Draw(screen)
}
