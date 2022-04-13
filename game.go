package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type game struct{}

func (g *game) Update() error {
	delta := time.Since(lastUpdatedAt).Seconds() * fps
	lastUpdatedAt = time.Now()

	if !ebiten.IsFocused() {
		return nil
	}

	player1.Update(delta)
	player2.Update(delta)
	ball1.Update(delta)

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, 0, 0, screenWidth, 0, gameColor)
	ebitenutil.DrawLine(screen, 0, screenHeight-1, screenWidth, screenHeight-1, gameColor)
	player1.Draw(screen)
	player2.Draw(screen)
	ball1.Draw(screen)
	scoreboard1.Draw(screen)
}

func (g *game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}
