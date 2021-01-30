package main

import (
	"fmt"
	"log"
	"math"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/meagar/asteroids/ship"
	"github.com/meagar/asteroids/star"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	Ship  *ship.Ship
	Stars []*star.Star
}

func NewGame() *Game {
	g := Game{
		Ship:  ship.New("assets/ship.png"),
		Stars: make([]*star.Star, 250),
	}

	for i := 0; i < len(g.Stars); i++ {
		g.Stars[i] = star.New(screenWidth, screenHeight)
	}

	return &g
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		r -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		r += 0.1
	}

	g.Ship.Dy *= 0.001
	g.Ship.Dx *= 0.001

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Ship.Dy = math.Cos(r)
		g.Ship.Dx = -math.Sin(r)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Ship.Dy = -math.Cos(r)
		g.Ship.Dx = math.Sin(r)
	}

	// move stars by delta
	for _, s := range g.Stars {
		s.X += g.Ship.Dx * s.Z * 0.01
		s.Y += g.Ship.Dy * s.Z * 0.01

		if s.X < 0 {
			s.X += screenWidth
		} else if s.X > screenWidth {
			s.X -= screenWidth
		}

		if s.Y < 0 {
			s.Y += screenHeight
		} else if s.Y > screenHeight {
			s.Y -= screenHeight
		}

		s.Op.GeoM.Reset()
		s.Op.GeoM.Translate(s.X, s.Y)
	}
	return nil
}

var r float64

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	// Draw stars
	for _, s := range g.Stars {
		screen.DrawImage(s.Img, s.Op)
		// screen.Set(int(stars[i].X), int(stars[i].Y), stars[i].C)

	}

	g.Ship.Op.GeoM.Reset()
	g.Ship.Op.GeoM.Translate(-g.Ship.Width/2, -g.Ship.Height/1.6)
	g.Ship.Op.GeoM.Rotate(r)
	g.Ship.Op.GeoM.Translate(g.Ship.Width/2, g.Ship.Height/1.6)

	g.Ship.Op.GeoM.Scale(0.2, 0.2)
	var x, y float64
	x = (screenWidth - (g.Ship.Width * 0.2)) / 2.0
	y = (screenHeight - (g.Ship.Height * 0.2)) / 2.0
	g.Ship.Op.GeoM.Translate(x, y)

	screen.DrawImage(g.Ship.Img, g.Ship.Op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := NewGame()

	// Sepcify the window size as you like. Here, a doulbed size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Your game's title")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
