package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type player struct {
	positionX float64
	positionY float64
	img       *ebiten.Image
	up        ebiten.Key
	down      ebiten.Key
}

func (p *player) Height() float64 {
	_, h := p.img.Size()

	return float64(h)
}

func (p *player) Width() float64 {
	w, _ := p.img.Size()

	return float64(w)
}

func (p *player) Update(delta float64) {
	if ebiten.IsKeyPressed(p.up) {
		p.positionY = math.Max(p.positionY-movementSpeed*delta, 0)
	}

	if ebiten.IsKeyPressed(p.down) {
		p.positionY = math.Min(p.positionY+movementSpeed*delta, screenHeight-p.Height())
	}
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
