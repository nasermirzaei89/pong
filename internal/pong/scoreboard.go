package pong

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type HUD struct {
	nums [10]*ebiten.Image
}

func (s *HUD) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, 0, 0, screenWidth, 0, gameColor)
	ebitenutil.DrawLine(screen, 0, screenHeight-1, screenWidth, screenHeight-1, gameColor)

	s.drawScore(screen, score1, screenWidth*.25, 3)
	s.drawScore(screen, score2, screenWidth*.75, 3)
}

func (s *HUD) drawScore(screen *ebiten.Image, score, x, y int) {
	for {
		opts := ebiten.DrawImageOptions{
			GeoM:          ebiten.GeoM{},
			ColorM:        ebiten.ColorM{},
			CompositeMode: 0,
			Filter:        0,
		}

		w, _ := s.nums[score%10].Size()
		x -= w + 1

		opts.GeoM.Translate(float64(x), float64(y))

		screen.DrawImage(s.nums[score%10], &opts)

		score /= 10
		if score == 0 {
			break
		}
	}
}
