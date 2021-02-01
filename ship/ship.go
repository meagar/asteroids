package ship

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	Width  float64
	Height float64
	Img    *ebiten.Image

	Dx float64
	Dy float64
	X  float64
	Y  float64
	Op *ebiten.DrawImageOptions
}

func New(path string) *Ship {
	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		panic(err)
	}

	w, h := img.Size()

	return &Ship{
		Img: img,
		// ScaleX: 0.5,
		// ScaleY: 0.5,
		Width:  float64(w),
		Height: float64(h),
		Op:     &ebiten.DrawImageOptions{},
	}
}
