package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nasermirzaei89/pong/internal/pong"
	"github.com/pkg/errors"
)

func main() {
	if err := ebiten.RunGame(pong.New()); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
