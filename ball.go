package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type ball struct {
	positionX float64
	positionY float64
	hSpeed    float64
	vSpeed    float64
	img       *ebiten.Image
}

func (b *ball) Width() float64 {
	w, _ := b.img.Size()

	return float64(w)
}

func (b *ball) Height() float64 {
	_, h := b.img.Size()

	return float64(h)
}

func (b *ball) Update(delta float64) {
	b.translate(delta)
	b.checkBounce()
	b.checkStatus()
}

func (b *ball) checkStatus() {
	if b.positionX <= -b.Width()-offScreenLength {
		score2++

		b.positionX = screenWidth / 2
		b.positionY = screenHeight / 2
	}

	if b.positionX >= screenWidth+b.Width()+offScreenLength {
		score1++

		b.positionX = screenWidth / 2
		b.positionY = screenHeight / 2
	}
}

func (b *ball) checkBounce() {
	if collides(b.positionX, b.positionY, b.Width(), b.Height(), player1.positionX, player1.positionY, player1.Width(), player1.Height()) {
		b.hSpeed = math.Abs(b.hSpeed)
	}

	if collides(b.positionX, b.positionY, b.Width(), b.Height(), player2.positionX, player2.positionY, player2.Width(), player2.Height()) {
		b.hSpeed = -math.Abs(b.hSpeed)
	}

	if b.positionY <= 0 || b.positionY >= screenHeight-b.Height() {
		b.vSpeed = -b.vSpeed
	}
}

func collides(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
}

func (b *ball) translate(delta float64) {
	b.positionX += b.hSpeed * delta

	b.positionY = math.Min(math.Max(b.positionY+b.vSpeed*delta, 0), screenHeight-b.Height())
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
