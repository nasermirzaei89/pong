package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type ball struct {
	positionX float64
	positionY float64
	hSpeed    float64
	vSpeed    float64
	img       *ebiten.Image
}

func (b *ball) Update() error {
	b.positionX += b.hSpeed * ebiten.DefaultTPS / ebiten.CurrentTPS()
	if b.positionX <= -600 {
		score2++
		b.positionX = screenWidth / 2
		b.positionY = screenHeight / 2
	}
	if b.positionX >= screenWidth+500 {
		score1++
		b.positionX = screenWidth / 2
		b.positionY = screenHeight / 2
	}

	for i := range objects {
		if p, ok := objects[i].(*player); ok {
			if b.positionX >= p.positionX-100 && b.positionX <= p.positionX+100 && b.positionY >= p.positionY && b.positionY <= p.positionY+540 {
				b.hSpeed = -b.hSpeed
			}
		}
	}

	b.positionY = math.Min(math.Max(b.positionY+b.vSpeed*ebiten.DefaultTPS/ebiten.CurrentTPS(), 0), screenHeight-100)
	if b.positionY <= 0 || b.positionY >= screenHeight-100 {
		b.vSpeed = -b.vSpeed
	}

	return nil
}

func (b *ball) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{
		GeoM:          ebiten.GeoM{},
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
	}

	opts.GeoM.Translate(b.positionX, b.positionY)

	screen.DrawImage(b.img, &opts)
}
