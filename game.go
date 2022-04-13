package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	delta := time.Since(lastUpdatedAt).Seconds()
	lastUpdatedAt = time.Now()

	if !ebiten.IsFocused() {
		return nil
	}

	player1.Update(delta)
	player2.Update(delta)
	ball.Update(delta)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	player1.Draw(screen)
	player2.Draw(screen)
	ball.Draw(screen)
	hud.Draw(screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}
