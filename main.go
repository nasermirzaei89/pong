package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/pkg/errors"
	"image"
	_ "image/png"
	"time"
)

const (
	screenWidth  = 3840
	screenHeight = 2160
	fps          = 60
)

var (
	objects       []object
	score1        int
	score2        int
	lastUpdatedAt time.Time
)

type game struct{}

func (g *game) Update() error {
	delta := time.Since(lastUpdatedAt).Seconds() * fps
	lastUpdatedAt = time.Now()

	if !ebiten.IsFocused() {
		return nil
	}

	for i := range objects {
		err := objects[i].Update(delta)
		if err != nil {
			return errors.Wrap(err, "error on update object")
		}
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for i := range objects {
		objects[i].Draw(screen)
	}
}

func (g *game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game1 := game{}

	img, _, err := ebitenutil.NewImageFromFile("sprites.png")
	if err != nil {
		panic(errors.Wrap(err, "error on new image from file"))
	}

	objects = append(objects, &player{
		positionX: 16,
		positionY: screenHeight / 2,
		img:       img.SubImage(image.Rect(340, 0, 440, 540)).(*ebiten.Image),
		up:        ebiten.KeyW,
		down:      ebiten.KeyS,
	})

	objects = append(objects, &player{
		positionX: screenWidth - 116,
		positionY: screenHeight / 2,
		img:       img.SubImage(image.Rect(440, 0, 540, 540)).(*ebiten.Image),
		up:        ebiten.KeyUp,
		down:      ebiten.KeyDown,
	})

	objects = append(objects, &ball{
		positionX: screenWidth / 2,
		positionY: screenHeight / 2,
		hSpeed:    32,
		vSpeed:    -32,
		img:       img.SubImage(image.Rect(400, 560, 500, 660)).(*ebiten.Image),
	})

	objects = append(objects, &scoreboard{nums: [10]*ebiten.Image{
		img.SubImage(image.Rect(192, 628, 259, 727)).(*ebiten.Image),
		img.SubImage(image.Rect(39, 20, 104, 119)).(*ebiten.Image),
		img.SubImage(image.Rect(24, 140, 115, 255)).(*ebiten.Image),
		img.SubImage(image.Rect(44, 271, 112, 400)).(*ebiten.Image),
		img.SubImage(image.Rect(40, 416, 128, 520)).(*ebiten.Image),
		img.SubImage(image.Rect(43, 544, 124, 660)).(*ebiten.Image),
		img.SubImage(image.Rect(156, 24, 232, 147)).(*ebiten.Image),
		img.SubImage(image.Rect(155, 168, 251, 291)).(*ebiten.Image),
		img.SubImage(image.Rect(179, 307, 259, 455)).(*ebiten.Image),
		img.SubImage(image.Rect(183, 480, 259, 603)).(*ebiten.Image),
	}})

	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Pong")

	lastUpdatedAt = time.Now()
	if err := ebiten.RunGame(&game1); err != nil {
		panic(errors.Wrap(err, "error on run game"))
	}
}
