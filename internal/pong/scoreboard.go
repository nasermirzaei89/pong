package pong

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type hud struct {
	nums [10]*ebiten.Image
}

func (s *hud) Draw(g *Game, screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, 0, 0, screenWidth, 0, g.gameColor)
	ebitenutil.DrawLine(screen, 0, screenHeight-1, screenWidth, screenHeight-1, g.gameColor)

	s.drawScore(screen, g.score1, screenWidth*.25, 3)
	s.drawScore(screen, g.score2, screenWidth*.75, 3)
}

func (s *hud) drawScore(screen *ebiten.Image, score, positionX, positionY int) {
	for {
		opts := ebiten.DrawImageOptions{
			GeoM:          ebiten.GeoM{},
			ColorM:        ebiten.ColorM{},
			CompositeMode: 0,
			Filter:        0,
		}

		w, _ := s.nums[score%10].Size()
		positionX -= w + 1

		opts.GeoM.Translate(float64(positionX), float64(positionY))

		screen.DrawImage(s.nums[score%10], &opts)

		score /= 10
		if score == 0 {
			break
		}
	}
}
