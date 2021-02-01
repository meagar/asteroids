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

func Make(numStars int, screenWidth, screenHeight float64) []*Star {
	imgCache := make([]*ebiten.Image, 255)

	for i := uint8(0); i < 255; i++ {
		img := ebiten.NewImage(1, 1)
		img.Set(0, 0, &color.Gray{Y: i})

		imgCache[i] = img
	}

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
		stars[i].Op.GeoM.Reset()
	}

	return stars
}
