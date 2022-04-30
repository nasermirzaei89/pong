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

type Game struct {
	player1 *player
	player2 *player
	ball1   *ball
	hud1    *hud

	score1 int
	score2 int

	gameColor color.Color
}

func (g *Game) Update() error {
	g.player1.Update()
	g.player2.Update()
	g.ball1.Update(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player1.Draw(screen)
	g.player2.Draw(screen)
	g.ball1.Draw(screen)
	g.hud1.Draw(g, screen)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func New() *Game {
	game1 := new(Game)

	game1.gameColor = color.White

	batImg := game1.imageFromData(dataBat)

	game1.player1 = &player{
		positionX: 1,
		positionY: screenHeight / 2,
		img:       batImg,
		up:        ebiten.KeyW,
		down:      ebiten.KeyS,
	}

	game1.player2 = &player{
		positionX: screenWidth - 2,
		positionY: screenHeight / 2,
		img:       batImg,
		up:        ebiten.KeyUp,
		down:      ebiten.KeyDown,
	}

	game1.ball1 = new(ball)
	game1.ball1.img = game1.imageFromData(dataBall)
	game1.ball1.resetPosition(time.Now().Add(warmUpDuration + idleDuration))

	game1.hud1 = &hud{nums: [10]*ebiten.Image{
		game1.imageFromData(dataNumber0),
		game1.imageFromData(dataNumber1),
		game1.imageFromData(dataNumber2),
		game1.imageFromData(dataNumber3),
		game1.imageFromData(dataNumber4),
		game1.imageFromData(dataNumber5),
		game1.imageFromData(dataNumber6),
		game1.imageFromData(dataNumber7),
		game1.imageFromData(dataNumber8),
		game1.imageFromData(dataNumber9),
	}}

	return game1
}
