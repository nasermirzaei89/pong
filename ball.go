package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	positionX float64
	positionY float64
	hSpeed    float64
	vSpeed    float64
	img       *ebiten.Image
	waitUntil time.Time
}

func (b *Ball) Width() float64 {
	w, _ := b.img.Size()

	return float64(w)
}

func (b *Ball) Height() float64 {
	_, h := b.img.Size()

	return float64(h)
}

func (b *Ball) Update() {
	if time.Now().Before(b.waitUntil) {
		return
	}

	b.translate()
	b.checkBounce()
	b.checkStatus()
}

func (b *Ball) checkStatus() {
	if b.positionX <= -b.Width() {
		score2++

		b.resetPosition()
	}

	if b.positionX >= screenWidth+b.Width() {
		score1++

		b.resetPosition()
	}
}

func (b *Ball) resetPosition() {
	b.positionX = screenWidth / 2
	b.positionY = screenHeight / 2

	b.waitUntil = time.Now().Add(idleDuration)
}

func (b *Ball) checkBounce() {
	if collides(b.positionX, b.positionY, b.Width(), b.Height(), player1.positionX, player1.positionY, player1.Width(), player1.Height()) {
		b.hSpeed = math.Abs(b.hSpeed)
	}

	if collides(b.positionX, b.positionY, b.Width(), b.Height(), player2.positionX, player2.positionY, player2.Width(), player2.Height()) {
		b.hSpeed = -math.Abs(b.hSpeed)
	}

	if b.positionY <= 0 {
		b.vSpeed = math.Abs(b.vSpeed)
	}

	if b.positionY >= screenHeight-b.Height() {
		b.vSpeed = -math.Abs(b.vSpeed)
	}
}

func collides(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
}

func (b *Ball) translate() {
	b.positionX += b.hSpeed

	b.positionY = math.Min(math.Max(b.positionY+b.vSpeed, 0), screenHeight-b.Height())
}

func (b *Ball) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{
		GeoM:          ebiten.GeoM{},
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
	}

	opts.GeoM.Translate(b.positionX, b.positionY)

	screen.DrawImage(b.img, &opts)
}
