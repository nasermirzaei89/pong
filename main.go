package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/pong/internal/pong"
	"github.com/pkg/errors"
)

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Pong")
	ebiten.SetRunnableOnUnfocused(false)

	if err := ebiten.RunGame(pong.New()); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
