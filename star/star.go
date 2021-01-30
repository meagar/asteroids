package star

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Star struct {
	X float64
	Y float64
	Z float64

	Img *ebiten.Image
	Op  *ebiten.DrawImageOptions
}

func New(screenWidth, screenHeight float64) *Star {
	depth := rand.Float64()
	img := ebiten.NewImage(1, 1)
	img.Fill(&color.RGBA{
		R: uint8(depth * 255),
		G: uint8(depth * 255),
		B: uint8(depth * 255),
		A: 255,
	})
	return &Star{
		X:   rand.Float64() * screenWidth,
		Y:   rand.Float64() * screenHeight,
		Z:   depth * 255,
		Op:  &ebiten.DrawImageOptions{},
		Img: img,
	}

}
