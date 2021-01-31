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

var imgCache map[byte]*ebiten.Image
var colorCache map[byte]*color.RGBA

func init() {
	imgCache = make(map[byte]*ebiten.Image)
	colorCache = make(map[byte]*color.RGBA)

	for i := uint8(0); i < 255; i++ {
		loadColor(i)
		loadImage(i)
	}
}

func loadImage(depth uint8) {
	img := ebiten.NewImage(1, 1)
	img.Fill(colorCache[depth])
	imgCache[depth] = img
}

func loadColor(depth uint8) {
	colorCache[depth] = &color.RGBA{
		R: depth,
		G: depth,
		B: depth,
		A: 255,
	}
}

func Make(numStars int, screenWidth, screenHeight float64) []*Star {
	stars := make([]*Star, numStars)

	for i := range stars {
		depth := 255 * rand.Float64()

		stars[i] = &Star{
			X:   rand.Float64() * screenWidth,
			Y:   rand.Float64() * screenHeight,
			Z:   depth,
			Op:  &ebiten.DrawImageOptions{},
			Img: imgCache[uint8(depth)],
		}
	}

	return stars
}
