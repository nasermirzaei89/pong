package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type object interface {
	Update() error
	Draw(screen *ebiten.Image)
}
