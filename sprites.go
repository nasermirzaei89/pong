package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	dataNumber0 = [][]bool{
		{true, true, true, true},
		{true, false, false, true},
		{true, false, false, true},
		{true, false, false, true},
		{true, true, true, true},
	}
	dataNumber1 = [][]bool{
		{true},
		{true},
		{true},
		{true},
		{true},
	}
	dataNumber2 = [][]bool{
		{true, true, true, true},
		{false, false, false, true},
		{true, true, true, true},
		{true, false, false, false},
		{true, true, true, true},
	}
	dataNumber3 = [][]bool{
		{true, true, true, true},
		{false, false, false, true},
		{true, true, true, true},
		{false, false, false, true},
		{true, true, true, true},
	}
	dataNumber4 = [][]bool{
		{true, false, false, true},
		{true, false, false, true},
		{true, true, true, true},
		{false, false, false, true},
		{false, false, false, true},
	}
	dataNumber5 = [][]bool{
		{true, true, true, true},
		{true, false, false, false},
		{true, true, true, true},
		{false, false, false, true},
		{true, true, true, true},
	}
	dataNumber6 = [][]bool{
		{true, true, true, true},
		{true, false, false, false},
		{true, true, true, true},
		{true, false, false, true},
		{true, true, true, true},
	}
	dataNumber7 = [][]bool{
		{true, true, true, true},
		{false, false, false, true},
		{false, false, false, true},
		{false, false, false, true},
		{false, false, false, true},
	}
	dataNumber8 = [][]bool{
		{true, true, true, true},
		{true, false, false, true},
		{true, true, true, true},
		{true, false, false, true},
		{true, true, true, true},
	}
	dataNumber9 = [][]bool{
		{true, true, true, true},
		{true, false, false, true},
		{true, true, true, true},
		{false, false, false, true},
		{true, true, true, true},
	}
	dataBat = [][]bool{
		{true},
		{true},
		{true},
		{true},
		{true},
		{true},
		{true},
		{true},
	}
	dataBall = [][]bool{
		{true},
	}
)

func imageFromData(data [][]bool) *ebiten.Image {
	h := len(data)
	w := len(data[0])
	img := ebiten.NewImage(w, h)

	for y := range data {
		for x := range data[y] {
			if data[y][x] {
				img.Set(x, y, gameColor)
			}
		}
	}

	return img
}
