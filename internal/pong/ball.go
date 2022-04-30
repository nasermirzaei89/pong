package pong

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type ball struct {
	positionX float64
	positionY float64
	hSpeed    float64
	vSpeed    float64
	img       *ebiten.Image
	waitUntil time.Time
}

func (obj *ball) Width() float64 {
	w, _ := obj.img.Size()

	return float64(w)
}

func (obj *ball) Height() float64 {
	_, h := obj.img.Size()

	return float64(h)
}

func (obj *ball) Update(game *Game) {
	if time.Now().Before(obj.waitUntil) {
		return
	}

	obj.translate()
	obj.checkBounce(game)
	obj.checkStatus(game)
}

func (obj *ball) checkStatus(game *Game) {
	if obj.positionX <= -obj.Width() {
		game.score2++

		obj.resetPosition(time.Now().Add(idleDuration))
	}

	if obj.positionX >= screenWidth+obj.Width() {
		game.score1++

		obj.resetPosition(time.Now().Add(idleDuration))
	}
}

func (obj *ball) resetPosition(waitUntil time.Time) {
	obj.positionX = screenWidth / 2
	obj.positionY = screenHeight / 2

	obj.hSpeed = float64(movementSpeed * randomize())
	obj.vSpeed = float64(movementSpeed * randomize())

	obj.waitUntil = waitUntil
}

func (obj *ball) checkBounce(game *Game) {
	if collides(
		obj.positionX, obj.positionY, obj.Width(), obj.Height(),
		game.player1.positionX, game.player1.positionY, game.player1.Width(), game.player1.Height(),
	) {
		obj.hSpeed = math.Abs(obj.hSpeed)
	}

	if collides(
		obj.positionX, obj.positionY, obj.Width(), obj.Height(),
		game.player2.positionX, game.player2.positionY, game.player2.Width(), game.player2.Height(),
	) {
		obj.hSpeed = -math.Abs(obj.hSpeed)
	}

	if obj.positionY <= 0 {
		obj.vSpeed = math.Abs(obj.vSpeed)
	}

	if obj.positionY >= screenHeight-obj.Height() {
		obj.vSpeed = -math.Abs(obj.vSpeed)
	}
}

func collides(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
}

func (obj *ball) translate() {
	obj.positionX += obj.hSpeed

	obj.positionY = math.Min(math.Max(obj.positionY+obj.vSpeed, 0), screenHeight-obj.Height())
}

func (obj *ball) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{
		GeoM:          ebiten.GeoM{},
		ColorM:        ebiten.ColorM{},
		CompositeMode: 0,
		Filter:        0,
	}

	opts.GeoM.Translate(obj.positionX, obj.positionY)

	screen.DrawImage(obj.img, &opts)
}

func randomize() int {
	if rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2) == 0 { //nolint:gosec
		return -1
	}

	return 1
}
