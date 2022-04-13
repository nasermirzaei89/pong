package pong

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth    = 192
	screenHeight   = 108
	movementSpeed  = 1
	idleDuration   = time.Second
	warmUpDuration = time.Second * 2
)

var (
	player1 *Player
	player2 *Player
	ball    *Ball
	hud     *HUD

	score1 int
	score2 int

	gameColor = color.White
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

func New() *Game {
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
		waitUntil: time.Now().Add(warmUpDuration + idleDuration),
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
	ebiten.SetRunnableOnUnfocused(false)

	return &game1
}
