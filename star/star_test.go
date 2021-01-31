package star_test

import (
	"testing"

	"github.com/meagar/asteroids/star"
)

func Benchmark10kStars(b *testing.B) {
	for n := 0; n < b.N; n++ {
		star.Make(10_000, 640, 480)
	}
}
