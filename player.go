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

func (p *player) Update() error {
	if ebiten.IsKeyPressed(p.up) {
		p.positionY = math.Max(p.positionY-playerSpeed*ebiten.DefaultTPS/ebiten.CurrentTPS(), 0)
	}

	if ebiten.IsKeyPressed(p.down) {
		p.positionY = math.Min(p.positionY+playerSpeed*ebiten.DefaultTPS/ebiten.CurrentTPS(), screenHeight-540)
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
