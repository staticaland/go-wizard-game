package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	playerRadius = 50
	playerSpeed  = 2
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

type Player struct {
	x, y     float64
	lastShot int64
}

func NewPlayer() *Player {
	return &Player{
		x:        400,
		y:        300,
		lastShot: time.Now().UnixNano(),
	}
}

func (p *Player) Update(bullets *Bullets) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.y -= playerSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.y += playerSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= playerSpeed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += playerSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && time.Now().UnixNano()-p.lastShot >= 1000000000 {
		*bullets = append(*bullets, NewBullet(p.x, p.y, 0, -2))
		p.lastShot = time.Now().UnixNano()
	}

	bullets.Update()
}

func (p *Player) Draw(screen *ebiten.Image, bullets *Bullets) {
	ebitenutil.DrawRect(screen, p.x-playerRadius, p.y-playerRadius, playerRadius*2, playerRadius*2, color.RGBA{255, 0, 0, 255})
	bullets.Draw(screen)
}
