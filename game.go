package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	player1.Update()
	player2.Update()
	ball.Update()

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
