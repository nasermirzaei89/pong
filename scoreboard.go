package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type scoreboard struct {
	nums [10]*ebiten.Image
}

func (s *scoreboard) Update(float64) error {
	return nil
}

func (s *scoreboard) Draw(screen *ebiten.Image) {
	s.drawScore(screen, score1, screenWidth/2-400, 100)
	s.drawScore(screen, score2, screenWidth/2+400, 100)

}

func (s *scoreboard) drawScore(screen *ebiten.Image, score, x, y int) {
	for {
		opts := ebiten.DrawImageOptions{
			GeoM:          ebiten.GeoM{},
			ColorM:        ebiten.ColorM{},
			CompositeMode: 0,
			Filter:        0,
		}

		opts.GeoM.Translate(float64(x), float64(y))

		screen.DrawImage(s.nums[score%10], &opts)
		x -= 100
		score /= 10
		if score == 0 {
			break
		}
	}
}
