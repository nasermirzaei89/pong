package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type object interface {
	Update(delta float64) error
	Draw(screen *ebiten.Image)
}
