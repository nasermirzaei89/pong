package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const playerSpeed = 32

type player struct {
	positionX float64
	positionY float64
	img       *ebiten.Image
	up        ebiten.Key
	down      ebiten.Key
}

func (p *player) Update(delta float64) error {
	if ebiten.IsKeyPressed(p.up) {
		p.positionY = math.Max(p.positionY-playerSpeed*delta, 0)
	}

	if ebiten.IsKeyPressed(p.down) {
		p.positionY = math.Min(p.positionY+playerSpeed*delta, screenHeight-540)
	}

	return nil
}

func (p *player) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{
		GeoM:          ebiten.GeoM{},
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
	}

	opts.GeoM.Translate(p.positionX, p.positionY)

	screen.DrawImage(p.img, &opts)
}
