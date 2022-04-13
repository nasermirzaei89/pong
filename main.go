package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
)

const (
	screenWidth     = 192
	screenHeight    = 108
	fps             = 60
	movementSpeed   = 1
	offScreenLength = 32
)

var (
	player1 *Player
	player2 *Player
	ball    *Ball
	hud     *HUD

	score1 int
	score2 int

	gameColor = color.White

	lastUpdatedAt time.Time
)

func main() {
	game1 := Game{}

	batImg := imageFromData(dataBat)

	player1 = &Player{
		positionX: 1,
		positionY: screenHeight / 2,
		img:       batImg,
		up:        ebiten.KeyW,
		down:      ebiten.KeyS,
	}

	player2 = &Player{
		positionX: screenWidth - 2,
		positionY: screenHeight / 2,
		img:       batImg,
		up:        ebiten.KeyUp,
		down:      ebiten.KeyDown,
	}

	ball = &Ball{
		positionX: screenWidth / 2,
		positionY: screenHeight / 2,
		hSpeed:    movementSpeed,
		vSpeed:    -movementSpeed,
		img:       imageFromData(dataBall),
	}

	hud = &HUD{nums: [10]*ebiten.Image{
		imageFromData(dataNumber0),
		imageFromData(dataNumber1),
		imageFromData(dataNumber2),
		imageFromData(dataNumber3),
		imageFromData(dataNumber4),
		imageFromData(dataNumber5),
		imageFromData(dataNumber6),
		imageFromData(dataNumber7),
		imageFromData(dataNumber8),
		imageFromData(dataNumber9),
	}}

	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Pong")

	lastUpdatedAt = time.Now()

	if err := ebiten.RunGame(&game1); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
