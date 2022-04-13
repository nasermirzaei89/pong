package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/nasermirzaei89/pong/internal/pong"
)

//nolint:gochecknoinits
func init() {
	mobile.SetGame(pong.New())
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
